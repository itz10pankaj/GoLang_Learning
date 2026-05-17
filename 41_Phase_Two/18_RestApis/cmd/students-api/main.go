package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itz10Pankaj/GoLang_Learning/41_Phase_Two/18_RestApis/internal/config"
)

func main() {
	// fmt.Println("Welcome to student APSI")4
	// load config
	cfg := config.MustLoad()

	// Router Set
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to studet Api"))
	})

	//set up server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server Started", slog.String("address", cfg.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start Server")
		}
	}()
	<-done

	slog.Info("Shutiing down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown", slog.String("error", err.Error()))
	}
	slog.Info("Server Shut down Successfully")

}
