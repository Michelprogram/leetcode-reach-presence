package main

import (
	"context"
	"leetcode-rich-presence/internal/config"
	"leetcode-rich-presence/internal/discord"
	"leetcode-rich-presence/internal/server"
	"leetcode-rich-presence/internal/server/handlers"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	queue := make(chan handlers.Message)

	conf, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		socket := server.NewServer(conf.Server.Port, queue)
		if err := socket.Start(); err != nil {
			slog.Error("Server error", "error", err)
			cancel()
		}
	}()

	go func() {

		dis, err := discord.NewDiscord(conf.Discord.ClientID, conf.Discord.ClientSecret)

		if err != nil {
			log.Fatalln(err)
		}

		if err := dis.ListenWithContext(ctx, queue); err != nil {
			slog.Error("Discord listener error", "error", err)
			cancel()
		}
	}()

	select {
	case <-sigChan:
		slog.Info("Shutdown signal received")
	case <-ctx.Done():
		slog.Info("Context cancelled")
	}

	cancel()

	time.Sleep(2 * time.Second)
	slog.Info("Application stopped")

}
