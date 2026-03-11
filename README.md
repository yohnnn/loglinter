# loglinter

`loglinter` — Go-анализатор для проверки сообщений логов в `log/slog` и `go.uber.org/zap`.

## Что проверяет

1. Сообщение начинается со строчной буквы.
2. Сообщение написано только на английском.
3. Сообщение не содержит спецсимволы.
4. Сообщение и аргументы не содержат чувствительные данные.

## Поддерживаемые методы

- `log/slog`: `Info`, `Error`, `Warn`, `Debug`, `InfoContext`, `ErrorContext`, `WarnContext`, `DebugContext`
- `go.uber.org/zap`: `Info`, `Error`, `Warn`, `Debug`, `Fatal`, `Panic`

## Использование как standalone-анализатора

Сборка:

```bash
go build -o loglint ./cmd/loglint
```

Запуск:

```bash
./loglint ./...
```

## Использование как плагина для golangci-lint (v2)

1) Собрать плагин:

```bash
go build -buildmode=plugin -o loglinter.so ./plugin
```

1) Создать конфиг `.golangci.yml`:

```yaml
version: "2"

linters:
  default: none
  enable:
    - loglinter
  settings:
    custom:
      loglinter:
        type: goplugin
        path: ./loglinter.so
        description: Checks that log messages follow best practices
        original-url: https://github.com/yohnnn/loglinter
        settings:
          enable_sensitive: true
```

1) Запуск:

```bash
golangci-lint run
```

Если используется snap-версия `golangci-lint` и возникает ошибка `plugin: not implemented`, используйте бинарь, установленный через `go install`.

## Минимальная конфигурация

По умолчанию проверка чувствительных данных включена.

В коде поддержан простой конфиг `enable_sensitive`:

- `true` — проверка включена;
- `false` — проверка отключена.

Также поддерживаются кастомные паттерны чувствительных данных через `sensitive_patterns`.

Пример:

```yaml
linters:
  settings:
    custom:
      loglinter:
        settings:
          enable_sensitive: true
          sensitive_patterns:
            - jwt
            - bearer
            - sessionid
```

Линтер будет проверять и встроенные ключевые слова, и ваши паттерны.

## Автоисправление

Для правила lowercase добавлен `SuggestedFix`:

- если сообщение начинается с заглавной буквы, линтер предлагает как должна выглядеть строчка.


## Тесты

```bash
go test ./analyzer -v
```

## Пример нарушений

```go
slog.Info("Starting server")      // uppercase
slog.Error("ошибка подключения")  // non-English
slog.Warn("something went wrong!") // special chars
slog.Info("user login", password) // sensitive arg
```
