package logtest

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

func slogTests() {
	ctx := context.Background()

	slog.Info("Starting server on port 8080") // want `log message "Starting server on port 8080" should start with a lowercase letter`
	slog.Error("Failed to connect")           // want `log message "Failed to connect" should start with a lowercase letter`

	slog.Info("запуск сервера")      // want `log message .* should be in English only`
	slog.Error("ошибка подключения") // want `log message .* should be in English only`

	slog.Info("server started!")         // want `log message "server started!" should not contain special punctuation`
	slog.Error("connection failed!!!")   // want `log message "connection failed!!!" should not contain special punctuation`
	slog.Warn("something went wrong...") // want `log message "something went wrong\.\.\." should not contain special punctuation`

	password := "secret123"
	slog.Info("user login", "pass", password) // want `log argument "password" may contain sensitive data`

	slog.Info("user password is set")             // want `log message contains potentially sensitive keyword "password"`
	slog.InfoContext(ctx, "user password is set") // want `log message contains potentially sensitive keyword "password"`
	slog.ErrorContext(ctx, "Failed to connect")   // want `log message "Failed to connect" should start with a lowercase letter`
	slog.WarnContext(ctx, "ошибка подключения")   // want `log message .* should be in English only`
	slog.DebugContext(ctx, "server started!")     // want `log message "server started!" should not contain special punctuation`

	apiToken := "tok-ctx"
	slog.InfoContext(ctx, "user login", "token", apiToken) // want `log argument "apiToken" may contain sensitive data`

	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
	slog.Debug("request completed")
	slog.Warn("retrying in 5 seconds")
	slog.InfoContext(ctx, "starting server on port 8080")
}

func zapTests() {
	logger := zap.NewNop()

	logger.Info("Starting server on port 8080") // want `log message "Starting server on port 8080" should start with a lowercase letter`
	logger.Error("Failed to connect")           // want `log message "Failed to connect" should start with a lowercase letter`
	logger.Fatal("Connection lost")             // want `log message "Connection lost" should start with a lowercase letter`

	logger.Info("запуск сервера")      // want `log message .* should be in English only`
	logger.Error("ошибка подключения") // want `log message .* should be in English only`

	logger.Info("server started!")         // want `log message "server started!" should not contain special punctuation`
	logger.Error("connection failed!!!")   // want `log message "connection failed!!!" should not contain special punctuation`
	logger.Warn("something went wrong...") // want `log message "something went wrong\.\.\." should not contain special punctuation`

	token := "tok-abc"
	logger.Info("user login", token) // want `log argument "token" may contain sensitive data`

	logger.Info("user password reset") // want `log message contains potentially sensitive keyword "password"`

	logger.Info("starting server on port 8080")
	logger.Error("failed to connect to database")
	logger.Debug("request completed successfully")
	logger.Warn("retrying connection in 5 seconds")
}
