package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//1. 基于errgroup实现一个http server的启动和关闭，以及linux signal信号的注册和处理，要保证能够一个退出，全部注销退出

//启动
func serverStart(stop <-chan struct{}) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rsp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rsp, "hello world~")
	})
	go func() {
		<-stop
		serverStop()
	}()
	log.Println("server start...")
	if err := http.ListenAndServe("0.0.0.0:8000", mux); err != nil {
		log.Fatal(err)
	}
}

//关闭
func serverStop() {
	log.Println("server stop...")
	os.Exit(1)
}

//信号量监听
func semctl(stop chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	s := <-c
	log.Printf("get exit signal: %s\n", s.String())
	stop <- struct{}{}
}

func main() {
	g := new(errgroup.Group)
	stop := make(chan struct{}, 1)
	g.Go(func() error {
		serverStart(stop)
		return nil
	})
	g.Go(func() error {
		semctl(stop)
		return nil
	})
	g.Wait()
}
