package main

import (
	"context"
	"e-commerce/cmd/app/container"
	"e-commerce/cmd/app/router"
	_ "e-commerce/cmd/docs"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"
)

func main() {
	container := container.BuildContainer()
	engine := router.InitGinEngine(container)
	server := &http.Server{
		Addr:    ":8888",
		Handler: engine,
	}
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	glog.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		glog.Fatal("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		glog.Info("Timeout of 5 seconds.")
	}
	glog.Info("Server exiting")
	fmt.Print(container)
}
