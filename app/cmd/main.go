package main

import (
	"context"
	"github.com/yvv4git/eger/app/internal/config"
	"github.com/yvv4git/eger/app/internal/tracer"
	"github.com/yvv4git/eger/app/transport"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const timeOutPrintConsole = 10 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Gracefully shutdown
	go func() {
		termCh := make(chan os.Signal, 1)
		signal.Notify(termCh, os.Interrupt, syscall.SIGINT)
		<-termCh
		log.Println("Signal was received")
		cancel()
	}()

	// Blocking operation
	processing(ctx)
	log.Println("Server is stopped")
}

func processing(ctx context.Context) {
	cfg := config.NewConfig()
	webSrv := transport.NewWebServer(cfg.WebSrv)
	closerJaeger := tracer.InitTracer(cfg.JaegerConf)
	defer func() {
		errCls := closerJaeger.Close()
		if errCls != nil {
			log.Println("Tracer server hasn't stopped successfully.")
		}
	}()

	go func() {
		log.Println("Starting web server")
		err := webSrv.ListenAndServe()
		if err != nil {
			log.Fatal("Web server err: ", err)
			return
		}
	}()

	for {
		select {
		case <-time.After(timeOutPrintConsole):
			log.Println("We are waiting for a few seconds")
		case <-ctx.Done():
			log.Println("Stop app")
			err := webSrv.Shutdown(ctx)
			if err != nil {
				log.Fatal("Problems with stopping the server: ", err)
			}
		}
	}
}
