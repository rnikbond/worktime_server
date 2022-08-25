package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"worktime_server/internal/handler"
	"worktime_server/internal/server"
	"worktime_server/internal/service"
)

func main() {

	fmt.Println("Server start ...")

	ctx, cancel := context.WithCancel(context.Background())

	services := service.NewService(ctx)
	h := handler.NewHandler(services)
	serv := new(server.Server)

	go func() {
		if err := serv.Run(":8080", h.InitRouter()); err != nil {
			log.Fatalf("failed run server: %v\n", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-done

	if err := serv.Shutdown(ctx); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("failed shutdown server: %v\n", err)
		}
	}

	cancel()

	fmt.Println("Server stop ...")

}
