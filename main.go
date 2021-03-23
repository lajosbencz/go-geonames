package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/lajosbencz/go-geonames/web"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	listen := ":4000"
	dsn := "geonames:geonames@tcp(127.0.0.1:3316)/geonames?charset=utf8mb4&parseTime=True&loc=Local"

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := web.NewServer(ctx, listen, dsn); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
