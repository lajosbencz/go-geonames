package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"

	"github.com/lajosbencz/go-geonames/web"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var showHelp bool
	var listen, dsnAddr, dsnUser, dsnPassword, dsnDatabase string

	flag.Usage = func() {
		_, name := path.Split(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", name)
		flag.PrintDefaults()
	}

	flag.StringVar(&listen, "listen", "127.0.0.1:4000", "HTTP listen address")
	flag.StringVar(&dsnAddr, "db_addr", "tcp(127.0.0.1:3306)", "Database address")
	flag.StringVar(&dsnUser, "db_user", "", "Username")
	flag.StringVar(&dsnPassword, "db_password", "", "Password")
	flag.StringVar(&dsnDatabase, "db_name", "geonames", "Database")
	flag.Parse()

	flag.BoolVar(&showHelp, "help", false, "Show usage info")
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	dsn := dsnAddr + "/" + dsnDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	if dsnUser != "" {
		if dsnPassword != "" {
			dsn = ":" + dsnPassword + "@" + dsn
		} else {
			dsn = "@" + dsn
		}
		dsn = dsnUser + dsn
	}
	fmt.Println(listen, dsn)
	return

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	srv, err := web.NewServer(ctx, listen, dsn)
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+s\n", err)
		}
	}()

	log.Println("Server listening on " + listen)
	<-ctx.Done()

	log.Printf("server stopped")

	// ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer func() {
	// 	cancel()
	// }()
	// if err = srv.Shutdown(ctxShutDown); err != nil {
	// 	log.Fatalf("server Shutdown Failed:%+s", err)
	// }
	log.Printf("server exited")
}
