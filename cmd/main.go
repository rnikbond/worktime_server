package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	fmt.Println("Server start ...")

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-done

	fmt.Println("Server stop ...")

}
