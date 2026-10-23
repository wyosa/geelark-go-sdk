# GeeLark Go SDK

[English](README.md)

Неофициальный Go SDK для [GeeLark API](https://open.geelark.com/api), устроенный в стиле `duoplus-go-sdk`: корневые клиенты, отдельные пакеты сервисов и явные типы запросов и ответов. Реализовано 206 путей; полный [список методов](docs/coverage.md) составлен по документации на 23 октября 2026 года.

## Быстрый старт

```go
client, err := geelark.NewClient(os.Getenv("GEELARK_TOKEN"))
if err != nil { log.Fatal(err) }

phones, err := client.CloudPhones.List(ctx, cloudphone.ListRequest{Page: 1, PageSize: 10})
if err != nil { log.Fatal(err) }
for _, phone := range phones.Items { fmt.Println(phone.ID, phone.SerialName) }
```

Примеры запускаются из корня репозитория:

```bash
GEELARK_TOKEN=your-token go run ./examples/basic
GEELARK_TOKEN=your-token go run ./examples/browser
GEELARK_APP_ID=your-app-id GEELARK_API_KEY=your-api-key go run ./examples/key-auth
```

Требуется Go 1.22+. Путь модуля: `github.com/wyosa/geelark-go-sdk`. Установка в проект:

```bash
go get github.com/wyosa/geelark-go-sdk@latest
```

Для эксплуатации используйте поддерживаемую версию Go с актуальными исправлениями. При аудите 23 октября 2026 года в установленном Go 1.27.2 найдена [GO-2026-6809](https://pkg.go.dev/vuln/GO-2026-6809); на Go 1.27.3 сканирование прошло без находок.

## Клиенты

- `geelark.NewClient(token)` — облачное API с Bearer-токеном.
- `geelark.NewKeyClient(appID, apiKey)` — облачное API с подписью запроса.
- `geelark.NewBrowserClient(token)` — локальное API браузера. Для него приложение GeeLark должно быть запущено, а пользователь — авторизован.

Облачные методы обращаются к `https://openapi.geelark.com`; локальные методы браузера — к `http://localhost:40185`. Метод получения категорий расширений браузера использует облачный сервер. Настройки: `WithBaseURL`, `WithRemoteBaseURL`, `WithHTTPClient`, `WithMinInterval`.

Каждый вызов отправляет `POST` с JSON и новым `traceId` UUIDv4. Код `0` означает успех; остальные коды возвращаются как `*geelark.APIError`. При частичном успехе (`40006`) ответ остаётся в `APIError.Data`; его можно разобрать через `DecodeData`.

Каждый экземпляр клиента SDK выдерживает интервал 310 мс между запросами к одному пути по умолчанию. Разные экземпляры не используют общий ограничитель. Повторные попытки автоматически не выполняются. Значимые нулевые значения передаются указателями, например `common.Ptr(0)` и `common.Ptr(false)`.

Параметры пользовательских облачных и браузерных RPA-сценариев передаются через `ParamMap` типа `map[string]any`: поддерживаются любые имена параметров, строки, массивы строк, числа и логические значения. В запросах обновления тегов облачного телефона, браузера и Analytics, а также сетевого чёрного списка `nil` пропускает поле, а `[]string{}` отправляет пустой массив. Например, `browser.UpdateRequest{ID: "browser-id", TagIDs: []string{}}` очищает теги браузера.

`Files.Upload` и `Library.Upload` получают подписанный URL и загружают файл без заголовков авторизации GeeLark. Для сохранения файла в библиотеке затем вызовите `Library.Create`. Пакет `webhook` содержит типы событий и функции проверки заголовков обратного вызова.

По умолчанию HTTP-таймаут API составляет 30 секунд. Загрузка в хранилище ограничивается контекстом; задавайте для неё deadline. Таймаут клиента, переданного через `WithHTTPClient`, применяется и к API, и к загрузке файлов.

## Проверка

```bash
go test -race -coverpkg=./... ./...
go vet ./...
```

Тесты проверяют все 206 путей, авторизацию, загрузку файлов, ошибки и примеры JSON из документации. Основным источником служат [живые страницы GeeLark](https://open.geelark.com/api/change-log); опубликованная [OpenAPI-спецификация](https://github.com/GeeLark/geelark-openapi/blob/main/docs/OPENAPI_COVERAGE.md) охватывает только часть методов.

Для каждого метода проверяются успех, частичный успех и HTTP-ошибка на локальном тестовом сервере. Результаты аудита и изменения публичных типов описаны в [отчёте](docs/coverage.md#audit-2026-10-23). Операции с реальным аккаунтом и приложением GeeLark не проверялись: ключей API в окружении не было.
