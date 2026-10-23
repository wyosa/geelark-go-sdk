# API coverage

Snapshot: 2026-10-23. The live GeeLark documentation is the source for paths and fields.

The SDK exposes 206 documented endpoint paths across 18 packages. The Browser extension category endpoint uses the remote host; all other Browser endpoints use the local client.

| Package | Methods |
| --- | ---: |
| `adb` | 2 |
| `analytics` | 9 |
| `app` | 19 |
| `automation` | 64 |
| `billing` | 7 |
| `browser` | 18 |
| `browserautomation` | 20 |
| `cloudphone` | 30 |
| `files` | 5 |
| `group` | 4 |
| `library` | 8 |
| `livestream` | 3 |
| `oem` | 1 |
| `proxy` | 5 |
| `proxygroup` | 4 |
| `shell` | 1 |
| `tag` | 4 |
| `webhook` | 2 |

## Endpoint map

| Package | SDK method | POST path | Official page |
| --- | --- | --- | --- |
| `cloudphone` | `Create` | `/open/v1/phone/addNew` | [Source](https://doc.geelark.com/web/#/602813392/101528476) |
| `cloudphone` | `List` | `/open/v1/phone/list` | [Source](https://doc.geelark.com/web/#/602813392/101527838) |
| `cloudphone` | `Status` | `/open/v1/phone/status` | [Source](https://doc.geelark.com/web/#/602813392/101527839) |
| `cloudphone` | `Start` | `/open/v1/phone/start` | [Source](https://doc.geelark.com/web/#/602813392/101527840) |
| `cloudphone` | `Stop` | `/open/v1/phone/stop` | [Source](https://doc.geelark.com/web/#/602813392/101527841) |
| `cloudphone` | `Restart` | `/open/v1/phone/restart` | [Source](https://doc.geelark.com/web/#/602813392/101529774) |
| `cloudphone` | `Delete` | `/open/v1/phone/delete` | [Source](https://doc.geelark.com/web/#/602813392/101527842) |
| `cloudphone` | `GetGPS` | `/open/v1/phone/gps/get` | [Source](https://doc.geelark.com/web/#/602813392/101527933) |
| `cloudphone` | `SetGPS` | `/open/v1/phone/gps/set` | [Source](https://doc.geelark.com/web/#/602813392/101527934) |
| `cloudphone` | `Update` | `/open/v1/phone/detail/update` | [Source](https://doc.geelark.com/web/#/602813392/101527935) |
| `cloudphone` | `Reset` | `/open/v2/phone/newOne` | [Source](https://doc.geelark.com/web/#/602813392/101528107) |
| `cloudphone` | `Clone` | `/open/v1/phone/clone` | [Source](https://doc.geelark.com/web/#/602813392/101529600) |
| `cloudphone` | `Screenshot` | `/open/v1/phone/screenShot` | [Source](https://doc.geelark.com/web/#/602813392/101528205) |
| `cloudphone` | `ScreenshotResult` | `/open/v1/phone/screenShot/result` | [Source](https://doc.geelark.com/web/#/602813392/101528206) |
| `cloudphone` | `SetRoot` | `/open/v1/root/setStatus` | [Source](https://doc.geelark.com/web/#/602813392/101528240) |
| `cloudphone` | `GetRoot` | `/open/v1/root/getStatus` | [Source](https://doc.geelark.com/web/#/602813392/101529825) |
| `cloudphone` | `DeviceID` | `/open/v1/phone/serialNum/get` | [Source](https://doc.geelark.com/web/#/602813392/101528243) |
| `cloudphone` | `SendSMS` | `/open/v1/phone/sendSms` | [Source](https://doc.geelark.com/web/#/602813392/101528255) |
| `cloudphone` | `Brands` | `/open/v1/phone/brand/list` | [Source](https://doc.geelark.com/web/#/602813392/101528264) |
| `cloudphone` | `TeamBrands` | `/open/v1/phone/brand/teamList` | [Source](https://doc.geelark.com/web/#/602813392/101529669) |
| `cloudphone` | `Transfer` | `/open/v1/phone/transfer` | [Source](https://doc.geelark.com/web/#/602813392/101528449) |
| `cloudphone` | `SetNetworkType` | `/open/v1/phone/net/set` | [Source](https://doc.geelark.com/web/#/602813392/101528562) |
| `cloudphone` | `HideAccessibility` | `/open/v1/phone/hideAccessibility` | [Source](https://doc.geelark.com/web/#/602813392/101529420) |
| `cloudphone` | `MoveGroup` | `/open/v1/phone/moveGroup` | [Source](https://doc.geelark.com/web/#/602813392/101529422) |
| `cloudphone` | `ImportContacts` | `/open/v1/phone/importContacts` | [Source](https://doc.geelark.com/web/#/602813392/101529436) |
| `cloudphone` | `ImportContactsResult` | `/open/v1/phone/importContactsResult` | [Source](https://doc.geelark.com/web/#/602813392/101529437) |
| `cloudphone` | `GetNetworkSettings` | `/open/v1/phone/netConfig/get` | [Source](https://doc.geelark.com/web/#/602813392/101529608) |
| `cloudphone` | `SetNetworkSettings` | `/open/v1/phone/netConfig/set` | [Source](https://doc.geelark.com/web/#/602813392/101529609) |
| `cloudphone` | `UIDump` | `/open/v1/phone/uiDump` | [Source](https://doc.geelark.com/web/#/602813392/101529828) |
| `cloudphone` | `UIDumpResult` | `/open/v1/phone/uiDump/result` | [Source](https://doc.geelark.com/web/#/602813392/101529829) |
| `automation` | `QueryTask` | `/open/v1/task/query` | [Source](https://doc.geelark.com/web/#/602813392/101528366) |
| `automation` | `ListTasks` | `/open/v1/task/historyRecords` | [Source](https://doc.geelark.com/web/#/602813392/101528505) |
| `automation` | `CancelTask` | `/open/v1/task/cancel` | [Source](https://doc.geelark.com/web/#/602813392/101528367) |
| `automation` | `RetryTask` | `/open/v1/task/restart` | [Source](https://doc.geelark.com/web/#/602813392/101528368) |
| `automation` | `TaskDetail` | `/open/v1/task/detail` | [Source](https://doc.geelark.com/web/#/602813392/101528416) |
| `automation` | `AddTask` | `/open/v1/task/add` | [Source](https://doc.geelark.com/web/#/602813392/101528369) |
| `automation` | `TikTokLogin` | `/open/v1/rpa/task/tiktokLogin` | [Source](https://doc.geelark.com/web/#/602813392/101528370) |
| `automation` | `TikTokEdit` | `/open/v1/rpa/task/tiktokEdit` | [Source](https://doc.geelark.com/web/#/602813392/101528371) |
| `automation` | `TikTokRandomStar` | `/open/v1/rpa/task/tiktokRandomStar` | [Source](https://doc.geelark.com/web/#/602813392/101528372) |
| `automation` | `TikTokRandomStarAsia` | `/open/v1/rpa/task/tiktokRandomStarAsia` | [Source](https://doc.geelark.com/web/#/602813392/101528499) |
| `automation` | `TikTokRandomComment` | `/open/v1/rpa/task/tiktokRandomComment` | [Source](https://doc.geelark.com/web/#/602813392/101528373) |
| `automation` | `TikTokRandomCommentAsia` | `/open/v1/rpa/task/tiktokRandomCommentAsia` | [Source](https://doc.geelark.com/web/#/602813392/101528639) |
| `automation` | `TikTokMessage` | `/open/v1/rpa/task/tiktokMessage` | [Source](https://doc.geelark.com/web/#/602813392/101528468) |
| `automation` | `TikTokMessageAsia` | `/open/v1/rpa/task/tiktokMessageAsia` | [Source](https://doc.geelark.com/web/#/602813392/101528500) |
| `automation` | `TikTokRandomFollow` | `/open/v1/rpa/task/tiktokRandomFollow` | [Source](https://doc.geelark.com/web/#/602813392/101529519) |
| `automation` | `TikTokRandomFollowAsia` | `/open/v1/rpa/task/tiktokRandomFollowAsia` | [Source](https://doc.geelark.com/web/#/602813392/101529520) |
| `automation` | `TikTokHide` | `/open/v1/rpa/task/tiktokHide` | [Source](https://doc.geelark.com/web/#/602813392/101529594) |
| `automation` | `TikTokHideAsia` | `/open/v1/rpa/task/tiktokHideAsia` | [Source](https://doc.geelark.com/web/#/602813392/101529595) |
| `automation` | `TikTokDel` | `/open/v1/rpa/task/tiktokDel` | [Source](https://doc.geelark.com/web/#/602813392/101529596) |
| `automation` | `TikTokDelAsia` | `/open/v1/rpa/task/tiktokDelAsia` | [Source](https://doc.geelark.com/web/#/602813392/101529597) |
| `automation` | `TikTokCommentInterception` | `/open/v1/rpa/task/tiktokCommentInterception` | [Source](https://doc.geelark.com/web/#/602813392/101529680) |
| `automation` | `TikTokCommentInterceptionAsia` | `/open/v1/rpa/task/tiktokCommentInterceptionAsia` | [Source](https://doc.geelark.com/web/#/602813392/101529683) |
| `automation` | `TikTokDeleteComment` | `/open/v1/rpa/task/tiktokDeleteComment` | [Source](https://doc.geelark.com/web/#/602813392/101529804) |
| `automation` | `TikTokDeleteCommentAsia` | `/open/v1/rpa/task/tiktokDeleteCommentAsia` | [Source](https://doc.geelark.com/web/#/602813392/101529805) |
| `automation` | `FaceBookLogin` | `/open/v1/rpa/task/faceBookLogin` | [Source](https://doc.geelark.com/web/#/602813392/101528374) |
| `automation` | `FaceBookAutoComment` | `/open/v1/rpa/task/faceBookAutoComment` | [Source](https://doc.geelark.com/web/#/602813392/101528375) |
| `automation` | `FaceBookPublish` | `/open/v1/rpa/task/faceBookPublish` | [Source](https://doc.geelark.com/web/#/602813392/101528376) |
| `automation` | `FaceBookActiveAccount` | `/open/v1/rpa/task/faceBookActiveAccount` | [Source](https://doc.geelark.com/web/#/602813392/101528377) |
| `automation` | `FaceBookPubReels` | `/open/v1/rpa/task/faceBookPubReels` | [Source](https://doc.geelark.com/web/#/602813392/101528378) |
| `automation` | `FaceBookMessage` | `/open/v1/rpa/task/faceBookMessage` | [Source](https://doc.geelark.com/web/#/602813392/101528469) |
| `automation` | `FacebookReelsActive` | `/open/v1/rpa/task/facebookReelsActive` | [Source](https://doc.geelark.com/web/#/602813392/101529618) |
| `automation` | `InstagramPubReels` | `/open/v1/rpa/task/instagramPubReels` | [Source](https://doc.geelark.com/web/#/602813392/101528379) |
| `automation` | `InstagramWarmup` | `/open/v1/rpa/task/instagramWarmup` | [Source](https://doc.geelark.com/web/#/602813392/101528405) |
| `automation` | `InstagramMessage` | `/open/v1/rpa/task/instagramMessage` | [Source](https://doc.geelark.com/web/#/602813392/101528470) |
| `automation` | `InstagramPubReelsImages` | `/open/v1/rpa/task/instagramPubReelsImages` | [Source](https://doc.geelark.com/web/#/602813392/101528501) |
| `automation` | `InstagramLogin` | `/open/v1/rpa/task/instagramLogin` | [Source](https://doc.geelark.com/web/#/602813392/101528509) |
| `automation` | `InstagramEdit` | `/open/v1/rpa/task/instagramEdit` | [Source](https://doc.geelark.com/web/#/602813392/101529538) |
| `automation` | `InstagramAccountStatus` | `/open/v1/rpa/task/instagramAccountStatus` | [Source](https://doc.geelark.com/web/#/602813392/101529679) |
| `automation` | `InstagramFollowAccount` | `/open/v1/rpa/task/instagramFollowAccount` | [Source](https://doc.geelark.com/web/#/602813392/101529716) |
| `automation` | `InstagramAiComment` | `/open/v1/rpa/task/instagramAiComment` | [Source](https://doc.geelark.com/web/#/602813392/101529717) |
| `automation` | `YouTubePubShort` | `/open/v1/rpa/task/youtubePubShort` | [Source](https://doc.geelark.com/web/#/602813392/101528380) |
| `automation` | `YouTubePubVideo` | `/open/v1/rpa/task/youtubePubVideo` | [Source](https://doc.geelark.com/web/#/602813392/101528381) |
| `automation` | `YouTubeActiveAccount` | `/open/v1/rpa/task/youTubeActiveAccount` | [Source](https://doc.geelark.com/web/#/602813392/101528382) |
| `automation` | `YouTubeEdit` | `/open/v1/rpa/task/youtubeEdit` | [Source](https://doc.geelark.com/web/#/602813392/101529696) |
| `automation` | `GoogleLogin` | `/open/v1/rpa/task/googleLogin` | [Source](https://doc.geelark.com/web/#/602813392/101528383) |
| `automation` | `GoogleAppDownload` | `/open/v1/rpa/task/googleAppDownload` | [Source](https://doc.geelark.com/web/#/602813392/101528471) |
| `automation` | `GoogleAppBrowser` | `/open/v1/rpa/task/googleAppBrowser` | [Source](https://doc.geelark.com/web/#/602813392/101528472) |
| `automation` | `SheinLogin` | `/open/v1/rpa/task/sheinLogin` | [Source](https://doc.geelark.com/web/#/602813392/101528384) |
| `automation` | `XPublish` | `/open/v1/rpa/task/xPublish` | [Source](https://doc.geelark.com/web/#/602813392/101528392) |
| `automation` | `ListFlows` | `/open/v1/task/flow/list` | [Source](https://doc.geelark.com/web/#/602813392/101528402) |
| `automation` | `CreateCustomTask` | `/open/v1/task/rpa/add` | [Source](https://doc.geelark.com/web/#/602813392/101528403) |
| `automation` | `ImportFlow` | `/open/v1/task/flow/import` | [Source](https://doc.geelark.com/web/#/602813392/101529544) |
| `automation` | `ExportFlow` | `/open/v1/task/flow/export` | [Source](https://doc.geelark.com/web/#/602813392/101529545) |
| `automation` | `MultiPlatformVideoDistribution` | `/open/v1/rpa/task/multiPlatformVideoDistribution` | [Source](https://doc.geelark.com/web/#/602813392/101528365) |
| `automation` | `FileUpload` | `/open/v1/rpa/task/fileUpload` | [Source](https://doc.geelark.com/web/#/602813392/101528391) |
| `automation` | `ImportContacts` | `/open/v1/rpa/task/importContacts` | [Source](https://doc.geelark.com/web/#/602813392/101528497) |
| `automation` | `KeyboxUpload` | `/open/v1/rpa/task/keyboxUpload` | [Source](https://doc.geelark.com/web/#/602813392/101528498) |
| `automation` | `RedditWarmup` | `/open/v1/rpa/task/redditWarmup` | [Source](https://doc.geelark.com/web/#/602813392/101528460) |
| `automation` | `RedditVideo` | `/open/v1/rpa/task/redditVideo` | [Source](https://doc.geelark.com/web/#/602813392/101528461) |
| `automation` | `RedditImage` | `/open/v1/rpa/task/redditImage` | [Source](https://doc.geelark.com/web/#/602813392/101528462) |
| `automation` | `PinterestVideo` | `/open/v1/rpa/task/pinterestVideo` | [Source](https://doc.geelark.com/web/#/602813392/101528559) |
| `automation` | `PinterestImage` | `/open/v1/rpa/task/pinterestImage` | [Source](https://doc.geelark.com/web/#/602813392/101528560) |
| `automation` | `ThreadsVideo` | `/open/v1/rpa/task/threadsVideo` | [Source](https://doc.geelark.com/web/#/602813392/101528654) |
| `automation` | `ThreadsImage` | `/open/v1/rpa/task/threadsImage` | [Source](https://doc.geelark.com/web/#/602813392/101528655) |
| `shell` | `Execute` | `/open/v1/shell/execute` | [Source](https://doc.geelark.com/web/#/602813392/101527932) |
| `app` | `ListInstalled` | `/open/v1/app/list` | [Source](https://doc.geelark.com/web/#/602813392/101527847) |
| `app` | `ListAvailable` | `/open/v1/app/installable/list` | [Source](https://doc.geelark.com/web/#/602813392/101527848) |
| `app` | `Install` | `/open/v1/app/install` | [Source](https://doc.geelark.com/web/#/602813392/101527849) |
| `app` | `Start` | `/open/v1/app/start` | [Source](https://doc.geelark.com/web/#/602813392/101527850) |
| `app` | `Stop` | `/open/v1/app/stop` | [Source](https://doc.geelark.com/web/#/602813392/101527851) |
| `app` | `Uninstall` | `/open/v1/app/uninstall` | [Source](https://doc.geelark.com/web/#/602813392/101527852) |
| `app` | `Upload` | `/open/v1/app/upload` | [Source](https://doc.geelark.com/web/#/602813392/101527967) |
| `app` | `UploadStatus` | `/open/v1/app/upload/status` | [Source](https://doc.geelark.com/web/#/602813392/101527968) |
| `app` | `ListShop` | `/open/v1/app/shop/list` | [Source](https://doc.geelark.com/web/#/602813392/101528644) |
| `app` | `AddToTeam` | `/open/v1/app/add` | [Source](https://doc.geelark.com/web/#/602813392/101528645) |
| `app` | `ListTeam` | `/open/v1/app/teamApp/list` | [Source](https://doc.geelark.com/web/#/602813392/101528688) |
| `app` | `RemoveTeam` | `/open/v1/app/remove` | [Source](https://doc.geelark.com/web/#/602813392/101528689) |
| `app` | `SetAutoInstall` | `/open/v1/app/setStatus` | [Source](https://doc.geelark.com/web/#/602813392/101528690) |
| `app` | `SetAuthorization` | `/open/v1/app/auth/status` | [Source](https://doc.geelark.com/web/#/602813392/101528691) |
| `app` | `SetRoot` | `/open/v1/app/root` | [Source](https://doc.geelark.com/web/#/602813392/101528692) |
| `app` | `BatchOperate` | `/open/v1/app/operation/batch` | [Source](https://doc.geelark.com/web/#/602813392/101529391) |
| `app` | `SetKeepAlive` | `/open/v1/app/setKeepAlive` | [Source](https://doc.geelark.com/web/#/602813392/101529442) |
| `app` | `SetAutoStart` | `/open/v1/app/setAutoStart` | [Source](https://doc.geelark.com/web/#/602813392/101529443) |
| `app` | `SetHidden` | `/open/v1/app/setHide` | [Source](https://doc.geelark.com/web/#/602813392/101529759) |
| `adb` | `SetStatus` | `/open/v1/adb/setStatus` | [Source](https://doc.geelark.com/web/#/602813392/101527853) |
| `adb` | `GetInfo` | `/open/v1/adb/getData` | [Source](https://doc.geelark.com/web/#/602813392/101527854) |
| `files` | `GetUploadURL` | `/open/v1/upload/getUrl` | [Source](https://doc.geelark.com/web/#/602813392/101527836) |
| `files` | `UploadToPhone` | `/open/v1/phone/uploadFile` | [Source](https://doc.geelark.com/web/#/602813392/101528108) |
| `files` | `UploadResult` | `/open/v1/phone/uploadFile/result` | [Source](https://doc.geelark.com/web/#/602813392/101528109) |
| `files` | `UploadKeybox` | `/open/v1/phone/keyboxUpload` | [Source](https://doc.geelark.com/web/#/602813392/101528555) |
| `files` | `KeyboxResult` | `/open/v1/phone/keyboxUpload/result` | [Source](https://doc.geelark.com/web/#/602813392/101528556) |
| `library` | `GetUploadURL` | `/open/v1/material/getUploadUrl` | [Source](https://doc.geelark.com/web/#/602813392/101529515) |
| `library` | `Create` | `/open/v1/material/create` | [Source](https://doc.geelark.com/web/#/602813392/101528424) |
| `library` | `Delete` | `/open/v1/material/del` | [Source](https://doc.geelark.com/web/#/602813392/101528425) |
| `library` | `SetTag` | `/open/v1/material/tag/set` | [Source](https://doc.geelark.com/web/#/602813392/101528426) |
| `library` | `Search` | `/open/v1/material/search` | [Source](https://doc.geelark.com/web/#/602813392/101528427) |
| `library` | `CreateTag` | `/open/v1/material/tag/create` | [Source](https://doc.geelark.com/web/#/602813392/101528428) |
| `library` | `DeleteTag` | `/open/v1/material/tag/del` | [Source](https://doc.geelark.com/web/#/602813392/101528429) |
| `library` | `SearchTags` | `/open/v1/material/tag/search` | [Source](https://doc.geelark.com/web/#/602813392/101528430) |
| `webhook` | `SetURL` | `/open/v1/callback/set` | [Source](https://doc.geelark.com/web/#/602813392/101527973) |
| `webhook` | `GetURL` | `/open/v1/callback/get` | [Source](https://doc.geelark.com/web/#/602813392/101527974) |
| `oem` | `Customization` | `/open/v1/phone/customization` | [Source](https://doc.geelark.com/web/#/602813392/101528564) |
| `analytics` | `Data` | `/open/v1/analytics/data` | [Source](https://doc.geelark.com/web/#/602813392/101528665) |
| `analytics` | `AddAccounts` | `/open/v1/analytics/accounts/add` | [Source](https://doc.geelark.com/web/#/602813392/101528675) |
| `analytics` | `DeleteAccount` | `/open/v1/analytics/accounts/delete` | [Source](https://doc.geelark.com/web/#/602813392/101528676) |
| `analytics` | `ListAccounts` | `/open/v1/analytics/accounts/list` | [Source](https://doc.geelark.com/web/#/602813392/101528677) |
| `analytics` | `UpdateAccount` | `/open/v1/analytics/accounts/update` | [Source](https://doc.geelark.com/web/#/602813392/101529393) |
| `analytics` | `CreateTag` | `/open/v1/analytics/tags/create` | [Source](https://doc.geelark.com/web/#/602813392/101529754) |
| `analytics` | `ListTags` | `/open/v1/analytics/tags/list` | [Source](https://doc.geelark.com/web/#/602813392/101529755) |
| `analytics` | `UpdateTag` | `/open/v1/analytics/tags/update` | [Source](https://doc.geelark.com/web/#/602813392/101529756) |
| `analytics` | `DeleteTags` | `/open/v1/analytics/tags/del` | [Source](https://doc.geelark.com/web/#/602813392/101529757) |
| `livestream` | `Start` | `/open/v1/phone/videoPush/upload` | [Source](https://doc.geelark.com/web/#/602813392/101529812) |
| `livestream` | `Result` | `/open/v1/phone/videoPush/upload/result` | [Source](https://doc.geelark.com/web/#/602813392/101529813) |
| `livestream` | `Stop` | `/open/v1/phone/videoPush/stop` | [Source](https://doc.geelark.com/web/#/602813392/101529814) |
| `browser` | `Create` | `/api/v1/browser/create` | [Source](https://doc.geelark.com/web/#/602813392/101528629) |
| `browser` | `Update` | `/api/v1/browser/update` | [Source](https://doc.geelark.com/web/#/602813392/101528630) |
| `browser` | `Delete` | `/api/v1/browser/delete` | [Source](https://doc.geelark.com/web/#/602813392/101528631) |
| `browser` | `List` | `/api/v1/browser/list` | [Source](https://doc.geelark.com/web/#/602813392/101528632) |
| `browser` | `Start` | `/api/v1/browser/start` | [Source](https://doc.geelark.com/web/#/602813392/101528633) |
| `browser` | `Stop` | `/api/v1/browser/stop` | [Source](https://doc.geelark.com/web/#/602813392/101528634) |
| `browser` | `Transfer` | `/api/v1/browser/transfer` | [Source](https://doc.geelark.com/web/#/602813392/101528635) |
| `browser` | `APIStatus` | `/api/v1/status` | [Source](https://doc.geelark.com/web/#/602813392/101528659) |
| `browser` | `Status` | `/api/v1/browser/status` | [Source](https://doc.geelark.com/web/#/602813392/101529527) |
| `browser` | `ClearCache` | `/api/v1/browser/deleteCache` | [Source](https://doc.geelark.com/web/#/602813392/101529528) |
| `browser` | `MoveGroup` | `/api/v1/browser/moveGroup` | [Source](https://doc.geelark.com/web/#/602813392/101529529) |
| `browser` | `Cookies` | `/api/v1/browser/getCookie` | [Source](https://doc.geelark.com/web/#/602813392/101529530) |
| `browser` | `SetBookmarks` | `/api/v1/browser/setBookmark` | [Source](https://doc.geelark.com/web/#/602813392/101529535) |
| `browser` | `GetBookmarks` | `/api/v1/browser/getBookmark` | [Source](https://doc.geelark.com/web/#/602813392/101529536) |
| `browser` | `Clone` | `/api/v1/browser/clone` | [Source](https://doc.geelark.com/web/#/602813392/101529610) |
| `browser` | `Kernels` | `/api/v1/browser/getKernelsList` | [Source](https://doc.geelark.com/web/#/602813392/101529621) |
| `browser` | `UpdateKernels` | `/api/v1/browser/updateKernels` | [Source](https://doc.geelark.com/web/#/602813392/101529622) |
| `browser` | `ExtensionGroups` | `/open/v1/browser/extGroup/list` | [Source](https://doc.geelark.com/web/#/602813392/101529797) |
| `browserautomation` | `Search` | `/api/v1/browser/task/search` | [Source](https://doc.geelark.com/web/#/602813392/101529495) |
| `browserautomation` | `Cancel` | `/api/v1/browser/task/cancel` | [Source](https://doc.geelark.com/web/#/602813392/101529496) |
| `browserautomation` | `Restart` | `/api/v1/browser/task/restart` | [Source](https://doc.geelark.com/web/#/602813392/101529497) |
| `browserautomation` | `Detail` | `/api/v1/browser/task/detail` | [Source](https://doc.geelark.com/web/#/602813392/101529522) |
| `browserautomation` | `Flow` | `/api/v1/browser/task/flow` | [Source](https://doc.geelark.com/web/#/602813392/101529498) |
| `browserautomation` | `Add` | `/api/v1/browser/task/add` | [Source](https://doc.geelark.com/web/#/602813392/101529499) |
| `browserautomation` | `TikTokSearch` | `/api/v1/browser/task/tiktokSearch` | [Source](https://doc.geelark.com/web/#/602813392/101529576) |
| `browserautomation` | `TikTokComment` | `/api/v1/browser/task/tiktokComment` | [Source](https://doc.geelark.com/web/#/602813392/101529577) |
| `browserautomation` | `TikTokLike` | `/api/v1/browser/task/tiktokLike` | [Source](https://doc.geelark.com/web/#/602813392/101529578) |
| `browserautomation` | `FacebookPost` | `/api/v1/browser/task/facebookPost` | [Source](https://doc.geelark.com/web/#/602813392/101529579) |
| `browserautomation` | `FacebookFriends` | `/api/v1/browser/task/facebookFriends` | [Source](https://doc.geelark.com/web/#/602813392/101529580) |
| `browserautomation` | `FacebookLike` | `/api/v1/browser/task/facebookLike` | [Source](https://doc.geelark.com/web/#/602813392/101529581) |
| `browserautomation` | `FacebookHomepage` | `/api/v1/browser/task/facebookHomepage` | [Source](https://doc.geelark.com/web/#/602813392/101529582) |
| `browserautomation` | `InstagramLike` | `/api/v1/browser/task/instagramLike` | [Source](https://doc.geelark.com/web/#/602813392/101529583) |
| `browserautomation` | `InstagramSearch` | `/api/v1/browser/task/instagramSearch` | [Source](https://doc.geelark.com/web/#/602813392/101529584) |
| `browserautomation` | `YouTubeWatch` | `/api/v1/browser/task/youtubeWatch` | [Source](https://doc.geelark.com/web/#/602813392/101529585) |
| `browserautomation` | `XPost` | `/api/v1/browser/task/xPost` | [Source](https://doc.geelark.com/web/#/602813392/101529586) |
| `browserautomation` | `XLike` | `/api/v1/browser/task/xLike` | [Source](https://doc.geelark.com/web/#/602813392/101529587) |
| `browserautomation` | `RedditLike` | `/api/v1/browser/task/redditLike` | [Source](https://doc.geelark.com/web/#/602813392/101529588) |
| `browserautomation` | `CookieBot` | `/api/v1/browser/task/cookieBot` | [Source](https://doc.geelark.com/web/#/602813392/101529589) |
| `proxy` | `Add` | `/open/v1/proxy/add` | [Source](https://doc.geelark.com/web/#/602813392/101527926) |
| `proxy` | `Update` | `/open/v1/proxy/update` | [Source](https://doc.geelark.com/web/#/602813392/101527927) |
| `proxy` | `Delete` | `/open/v1/proxy/delete` | [Source](https://doc.geelark.com/web/#/602813392/101527928) |
| `proxy` | `List` | `/open/v1/proxy/list` | [Source](https://doc.geelark.com/web/#/602813392/101527929) |
| `proxy` | `Check` | `/open/v1/proxy/check` | [Source](https://doc.geelark.com/web/#/602813392/101528662) |
| `proxygroup` | `Create` | `/open/v1/proxyGroup/add` | [Source](https://doc.geelark.com/web/#/602813392/101529750) |
| `proxygroup` | `List` | `/open/v1/proxyGroup/search` | [Source](https://doc.geelark.com/web/#/602813392/101529751) |
| `proxygroup` | `Update` | `/open/v1/proxyGroup/update` | [Source](https://doc.geelark.com/web/#/602813392/101529752) |
| `proxygroup` | `Delete` | `/open/v1/proxyGroup/del` | [Source](https://doc.geelark.com/web/#/602813392/101529753) |
| `tag` | `Create` | `/open/v1/tag/add` | [Source](https://doc.geelark.com/web/#/602813392/101527936) |
| `tag` | `Delete` | `/open/v1/tag/delete` | [Source](https://doc.geelark.com/web/#/602813392/101527937) |
| `tag` | `List` | `/open/v1/tag/list` | [Source](https://doc.geelark.com/web/#/602813392/101527938) |
| `tag` | `Update` | `/open/v1/tag/update` | [Source](https://doc.geelark.com/web/#/602813392/101527939) |
| `group` | `Create` | `/open/v1/group/add` | [Source](https://doc.geelark.com/web/#/602813392/101527940) |
| `group` | `Delete` | `/open/v1/group/delete` | [Source](https://doc.geelark.com/web/#/602813392/101527941) |
| `group` | `List` | `/open/v1/group/list` | [Source](https://doc.geelark.com/web/#/602813392/101527942) |
| `group` | `Update` | `/open/v1/group/update` | [Source](https://doc.geelark.com/web/#/602813392/101527943) |
| `billing` | `Balance` | `/open/v1/pay/wallet` | [Source](https://doc.geelark.com/web/#/602813392/101529433) |
| `billing` | `Plans` | `/open/v1/pay/profiles/list` | [Source](https://doc.geelark.com/web/#/602813392/101529504) |
| `billing` | `ChangePlan` | `/open/v1/pay/plan/upgrade` | [Source](https://doc.geelark.com/web/#/602813392/101529505) |
| `billing` | `RenewPlan` | `/open/v1/pay/plan/continue` | [Source](https://doc.geelark.com/web/#/602813392/101529508) |
| `billing` | `PlanInfo` | `/open/v1/pay/plan/info` | [Source](https://doc.geelark.com/web/#/602813392/101529510) |
| `billing` | `Transactions` | `/open/v1/billing/transaction/detail` | [Source](https://doc.geelark.com/web/#/602813392/101529561) |
| `billing` | `BuyTime` | `/open/v1/pay/timeAddOn/buy` | [Source](https://doc.geelark.com/web/#/602813392/101529625) |

The [proxy-group list page](https://open.geelark.com/api/proxy-group-search) prints `api.geelark.com`. During the audit this hostname did not resolve in the test environment. The same path on `openapi.geelark.com` returned GeeLark's authentication error (`40003`) to an unauthenticated request. The SDK retains the standard cloud host; successful authenticated operation still needs account-level verification.

## Audit 2026-10-23

Reviewed all 219 live documentation pages, the 206 endpoint paths, request and response field tables, the Markdown bundle, and the [changelog](https://open.geelark.com/api/change-log). The latest entries were dated 2026-10-21 and 2026-10-20; their startup settings and per-phone Android version fields were already present. The downloadable bundle lagged behind the live pages.

The initial test suite passed despite the defects below. New regression tests reproduced the defects before the fixes.

| Area | Defect and correction |
| --- | --- |
| Billing | `giftMoney`, plan `price`, and transaction `amount` were integers, causing JSON decoding to fail for fractional values. They now use `float64`, matching the [balance](https://open.geelark.com/api/pay-wallet), [plan](https://open.geelark.com/api/pay-profiles-list), and [transaction](https://open.geelark.com/api/billing-transaction-detail) schemas. |
| Custom flows | `paramMap` accepted only `Title`, `Desc`, and `Video` from an example. [Cloud](https://open.geelark.com/api/task-rpa-add) and [browser](https://open.geelark.com/api/browser-task-add) requests now accept arbitrary named parameters, including zero and false. |
| Update lists | `omitempty` discarded explicitly empty tag lists and network blacklists. Custom JSON serialization now distinguishes nil from empty slices, allowing operations such as [clearing browser tags](https://open.geelark.com/api/edit-browser). |
| Task results | Query/history/detail models dropped failure codes and descriptions; query/history dropped execution duration, and query dropped the publication link. The documented fields are now retained. [Query](https://open.geelark.com/api/task-query), [detail](https://open.geelark.com/api/task-detail). |
| Log pagination | Millisecond `searchAfter` cursors used `[]int`, which cannot represent the documented timestamps on 32-bit systems. Request and response cursors now use `[]int64`. |
| Other response fields | Group and tag lookup failures are retained; browser responses accept the documented `memoryDevice` field alongside legacy `memeryDevice`. App version codes accept both numeric JSON examples and decimal strings from the [schema](https://open.geelark.com/api/app-installable-list). |
| Response validation | Empty objects, JSON null, and missing/null result codes could be reported as success; some 3xx statuses were also accepted. Success now requires a 2xx HTTP status and an explicit integer `code: 0`. Non-2xx statuses return `APIError`. |
| Rate limiting | Cancelled requests reserved future slots, so repeated cancellations could block later calls for a long time. Slots are now reserved only when granted; JSON serialization failures do not consume slots. |
| Upload timeouts | Signed uploads discarded a custom HTTP client's timeout. They now honor it while preserving context-controlled uploads when using the default client. |
| Base URLs | Empty query/fragment delimiters and empty hostnames were accepted and could produce requests to the wrong path. Constructors now reject them. |

Public type changes requiring caller updates:

- `billing.BalanceResponse.GiftMoney`, `billing.PlansResponseItem.Price`, and `billing.TransactionsResponseListItem.Amount`: `int` to `float64`.
- `automation.TaskDetailRequest.SearchAfter` and `TaskDetailResponse.SearchAfter`: `[]int` to `[]int64`.
- `automation.CreateCustomTaskRequest.ParamMap` and `browserautomation.AddRequest.ParamMap`: pointers to fixed structs to named `map[string]any` types. Use map literals with quoted parameter names; pointers are no longer required.

Validation includes 618 endpoint scenarios (206 methods × success, partial success, and HTTP failure), 356 existing documentation examples, and new transport/schema regressions. Partial-success tests decode `APIError.Data` for every method. Concurrency, cancellation, redirects, signed URL redaction, response size limits, authentication, and storage uploads are covered by local tests. Aggregate statement coverage is 90.2% across `./...`, including the example programs.

Passed: `go test -race -count=1 -coverpkg=./... ./...`, `go vet ./...`, and Staticcheck with Go 1.27.3, plus compatibility tests with Go 1.22.12 (`CGO_ENABLED=0` to avoid the old toolchain's incompatibility with the installed macOS SDK). Linux/386 test binaries and Windows/amd64 packages also cross-compiled successfully; this does not execute them on those operating systems.

The installed Go 1.27.2 standard library is affected by [GO-2026-6809](https://pkg.go.dev/vuln/GO-2026-6809). Govulncheck found no vulnerabilities when run with Go 1.27.3. The system Go installation was not changed; `GOTOOLCHAIN=go1.27.3` selects the patched toolchain for these checks.

Live account verification remains outstanding: no `GEELARK_TOKEN`, `GEELARK_APP_ID`, `GEELARK_API_KEY`, or local `.env` was available. Local HTTP tests establish request construction, decoding, and transport behavior; they do not establish successful cloud phone startup, RPA execution, billing operations, actual storage uploads, or desktop browser behavior. The Analytics update page and changelog also describe omission of `tagIds` differently; the SDK now preserves an explicit empty array, but the server's behavior for an omitted field should be checked on a test account.
