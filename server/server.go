package server

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/envoy49/go-spotify-cli/config"

	"github.com/envoy49/go-spotify-cli/routes"
	"github.com/sirupsen/logrus"
)

const (
	serverPort = ":4949"
)

func Server(ctx context.Context, cfg *config.Config) {
	// Create a new server instance each time
	server := &http.Server{Addr: serverPort}

	routes.SetupRoutes(cfg)

	// Start the server in a goroutine
	go func() {
		logrus.Println("Opened server to get an auth token on " + config.ServerUrl)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logrus.WithError(err).Error("Error starting the server")
		}
	}()

	// Listen for the context being canceled
	<-ctx.Done()

	// Create a deadline to wait for
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.WithError(err).Error("Error shutting down the server")
	}
}

func StartServer(cfg *config.Config, route string) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	go Server(ctx, cfg)

	if err := waitForServerReady(); err != nil {
		logrus.WithError(err).Error("Error waiting for the server to start")
		return cancel
	}

	resp, err := http.Get(config.ServerUrl + route)
	if err != nil {
		logrus.WithError(err).Error("Error making the GET request to: " + route)
		return cancel
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.WithError(err).Error("Error closing request to :" + route)
		}
	}()

	return cancel
}

func waitForServerReady() error {
	parsedURL, err := url.Parse(config.ServerUrl)
	if err != nil {
		return err
	}

	deadline := time.Now().Add(2 * time.Second)
	var lastErr error

	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", parsedURL.Host, 50*time.Millisecond)
		if err == nil {
			return conn.Close()
		}
		lastErr = err
		time.Sleep(50 * time.Millisecond)
	}

	return lastErr
}
