package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	httpGateway "marine/internal/gateways/http"
	"marine/internal/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
	}
	userService := usecase.NewUserService()
	scenarioService := usecase.NewScenarioService()
	materialService := usecase.NewMaterialService()
	useCases := httpGateway.UseCases{
		User:     userService,
		Scenario: scenarioService,
		Material: materialService,
	}
	r := httpGateway.NewServer(useCases)
	server.Handler = r

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, _ := errgroup.WithContext(context.Background())
	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)
	eg.Go(func() error {
		s := <-sigQuit
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Println(err.Error())
		}
		return fmt.Errorf("captured signal: %v", s)
	})

	go func() {
		if err := r.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error during server shutdown: %v", err)
		}
	}()
	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the server: %v", err) // gracefully shutting down the server: captured signal: interrupt
	}

}
