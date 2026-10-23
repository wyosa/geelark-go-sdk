#!/usr/bin/env python3
"""Verify fixtures offline with `check`; download official API examples with `refresh`."""

import argparse
import concurrent.futures
import datetime
import hashlib
import html
import json
from pathlib import Path
import re
import urllib.parse
import urllib.request

ROOT = Path(__file__).resolve().parent.parent
DATA = ROOT / "testdata" / "official"
API = "https://doc.geelark.com/server/index.php?s=/api/"
URL_PATTERN = r"https?://[^/\s`]+(?:/open/v[12]|/api/v1)/[^`\s<\)\";]+"


def read_json(path):
    return json.loads(path.read_text())


def write_json(path, value):
    path.write_text(json.dumps(value, ensure_ascii=False, indent=2) + "\n")


def blocks(content):
    """Extract fenced and four-space indented examples in source order."""
    matches = list(re.finditer(r"```[^\n]*\n(.*?)```", content, re.S))
    result = [(match.start(), match.group(1)) for match in matches]
    for match in re.finditer(r"(?m)^(?:[ ]{4}[^\n]*\n|[ \t]*\n)+", content):
        if any(fence.start() <= match.start() < fence.end() for fence in matches):
            continue
        body = "\n".join(line[4:] if line.startswith("    ") else ""
                         for line in match.group().splitlines())
        if body.lstrip().startswith(("{", "[")):
            result.append((match.start(), body))
    return sorted(result)


def clean(raw):
    """Remove presentation whitespace, comments, and trailing commas only.

    Preserve keys, types, numbers, and string contents. Remaining invalid JSON
    is recorded in fixture-issues.json rather than silently repaired.
    """
    out = []
    quoted = escaped = False
    i = 0
    while i < len(raw):
        char = raw[i]
        if quoted:
            out.append(char)
            if escaped:
                escaped = False
            elif char == "\\":
                escaped = True
            elif char == '"':
                quoted = False
            i += 1
            continue
        if char == '"':
            quoted = True
        elif raw[i:i + 2] == "//":
            while i < len(raw) and raw[i] != "\n":
                i += 1
            continue
        elif raw[i:i + 2] == "/*":
            end = raw.find("*/", i + 2)
            i = len(raw) if end < 0 else end + 2
            continue
        if char == ",":
            j = i + 1
            while j < len(raw) and raw[j].isspace():
                j += 1
            if j < len(raw) and raw[j] in "]}":
                i += 1
                continue
        out.append(" " if char == "\u00a0" and not quoted else char)
        i += 1
    return "".join(out)


def unique(values):
    return list({json.dumps(value, sort_keys=True): value for value in values}.values())


def derive(snapshot):
    manifest = read_json(ROOT / "testdata" / "endpoints.json")["endpoints"]
    pages = {page["page_id"]: page for page in snapshot["pages"]}
    for page in pages.values():
        digest = hashlib.sha256(page["content"].encode()).hexdigest()
        if digest != page["sha256"]:
            raise ValueError("Source checksum differs: " + page["page_id"])
    documented = {urllib.parse.urlsplit(url).path
                  for page in pages.values()
                  for url in re.findall(URL_PATTERN, page["content"])}
    implemented = {entry["path"] for entry in manifest}
    if documented != implemented:
        raise ValueError(f"Endpoint drift: missing={sorted(documented - implemented)}, "
                         f"extra={sorted(implemented - documented)}")
    fixtures, issues = [], []
    for entry in manifest:
        page = pages[entry["page_id"]]
        urls = re.findall(URL_PATTERN, page["content"])
        if not urls or urllib.parse.urlsplit(urls[0]).path != entry["path"]:
            raise ValueError("Endpoint source differs: " + entry["method"])
        requests, responses, callbacks = [], [], []
        for _, raw in blocks(page["content"]):
            if not raw.strip().startswith(("{", "[")):
                continue
            try:
                value = json.loads(clean(raw))
            except json.JSONDecodeError:
                issues.append({"endpoint": entry["package"] + "." + entry["method"],
                               "problem": "invalid_example", "raw": raw})
                continue
            if not isinstance(value, dict):
                issues.append({"endpoint": entry["package"] + "." + entry["method"],
                               "problem": "top_level_array_example", "raw": raw})
                continue
            if "Content-Type" in value or "Authorization" in value:
                continue
            if "code" in value and any(key in value for key in ("msg", "traceId", "data")):
                responses.append(value)
            elif "event" in value:
                callbacks.append(value)
            else:
                requests.append(value)
        fixtures.append(dict(entry, url=urls[0], slug=page["slug"],
                             requests=unique(requests), responses=unique(responses),
                             callbacks=unique(callbacks)))
    callbacks = []
    for _, raw in blocks(pages["101527975"]["content"]):
        if not raw.strip().startswith("{"):
            continue
        value = json.loads(clean(raw))
        if "type" in value:
            callbacks.append(value)
    return {"fixtures.json": fixtures, "fixture-issues.json": issues,
            "callbacks.json": unique(callbacks)}


def post(method, params):
    request = urllib.request.Request(API + method,
                                     data=urllib.parse.urlencode(params).encode())
    with urllib.request.urlopen(request, timeout=30) as response:
        result = json.load(response)
    if result.get("error_code") != 0:
        raise ValueError("Documentation request failed: " + method)
    return result["data"]


def refresh():
    def walk(node, category=""):
        for page in node.get("pages", []):
            yield page["page_id"], category
        for child in node.get("catalogs", []):
            yield from walk(child, (category + "/" + child["cat_name"]).strip("/"))

    def fetch(entry):
        page_id, category = entry
        page = post("page/info", {"page_id": page_id})
        content = html.unescape(page["page_content"])
        slug = page.get("page_unique_id", "")
        url = "https://open.geelark.com/api/" + slug if slug else (
            "https://doc.geelark.com/web/#/602813392/" + page_id)
        return dict(page_id=page_id, title=page["page_title"], slug=slug,
                    category=category, url=url, content=content,
                    sha256=hashlib.sha256(content.encode()).hexdigest())

    entries = list(walk(post("item/info", {"item_id": "602813392"})["menu"]))
    with concurrent.futures.ThreadPoolExecutor(max_workers=6) as pool:
        pages = list(pool.map(fetch, entries))
    return {"snapshot": datetime.datetime.now(datetime.timezone.utc).date().isoformat(),
            "catalog_url": "https://doc.geelark.com/web/#/602813392", "pages": pages}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("command", choices=["check", "refresh"])
    args = parser.parse_args()
    snapshot = refresh() if args.command == "refresh" else read_json(DATA / "pages.json")
    derived = derive(snapshot)
    if args.command == "refresh":
        # Validate the entire download before replacing any pinned inputs.
        write_json(DATA / "pages.json", snapshot)
        for name, value in derived.items():
            write_json(DATA / name, value)
        print("Refreshed snapshot; review source, fixture, and contract-test differences.")
    else:
        for name, value in derived.items():
            if value != read_json(DATA / name):
                raise SystemExit(f"Derived fixture differs: {name}; refresh and review the snapshot")
        print(f"Verified {len(snapshot['pages'])} source hashes and "
              f"{len(derived['fixtures.json'])} endpoint fixtures offline.")


if __name__ == "__main__":
    main()
