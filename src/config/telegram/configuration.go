package telegram

import (
	"bybitbot/src/controllers/handlers"
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/go-faster/errors"
	"github.com/gotd/td/examples"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"

	updhook "github.com/gotd/td/telegram/updates/hook"
)

var handler tg.UpdateDispatcher // handler
var logger *zap.Logger          // logger
var gaps *updates.Manager
var flow auth.Flow               //auth
var tele_client *telegram.Client //client

func CreateTelegramConfig() {
	keys = LoadTelegramEnvironment()

	SetTelegramClient()
	RunTelegramClient()
}

func SetTelegramClient() {
	SetHandlerConfig()
	SetAuthenticationConfig()
	SetMessageHandlers()
	gaps = updates.New(updates.Config{
		Handler: handler,
		//Logger:  logger.Named("gaps"),
	})

	// Environment variables:
	// 	APP_ID:         app_id of Telegram app.
	// 	APP_HASH:       app_hash of Telegram app.
	// 	SESSION_FILE:   path to session file
	// 	SESSION_DIR:    path to session directory, if SESSION_FILE is not set
	client, err := telegram.ClientFromEnvironment(telegram.Options{
		Logger:        logger,
		UpdateHandler: gaps,
		Middlewares: []telegram.Middleware{
			updhook.UpdateHook(gaps.Handle),
		},
	})
	if err != nil {
		fmt.Printf("error creating telegram client: %v", err)
	}

	tele_client = client

}

func RunTelegramClient() {
	ctx := context.Background()
	// Async run
	go tele_client.Run(ctx, func(ctx context.Context) error {
		// Perform auth if no session is available.
		if err := tele_client.Auth().IfNecessary(ctx, flow); err != nil {
			return errors.Wrap(err, "auth")
		}

		// Fetch user info.
		user, err := tele_client.Self(ctx)
		if err != nil {
			return errors.Wrap(err, "call self")
		}

		return gaps.Run(ctx, tele_client.API(), user.ID, updates.AuthOptions{
			OnStart: func(ctx context.Context) {
				//logger.Info("Gaps started")
			},
		})
	})
}

func SetHandlerConfig() {
	//obtain logger
	log, _ := zap.NewDevelopment(zap.IncreaseLevel(zapcore.InfoLevel), zap.AddStacktrace(zapcore.FatalLevel))
	logger = log
	//obtain handler
	handler = tg.NewUpdateDispatcher()
}

func SetAuthenticationConfig() {
	flow = auth.NewFlow(examples.Terminal{}, auth.SendCodeOptions{})
}

func SetMessageHandlers() {
	// Setup message update handlers.
	handler.OnNewChannelMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
		//logger.Info("Channel message", zap.Any("message", update.Message))
		handlers.TelegramMessageHandler("Channel", update.Message)
		return nil
	})
	handler.OnNewMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
		//logger.Info("Message", zap.Any("message", update.Message))
		handlers.TelegramMessageHandler("User", update.Message)
		return nil
	})
}

func CloseTelegramClient() {
	logger.Sync()
}
