package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ebrianne/adguard-exporter/config"
	"github.com/ebrianne/adguard-exporter/internal/metrics"
	"github.com/ebrianne/adguard-exporter/internal/adguard"
	"github.com/ebrianne/adguard-exporter/internal/server"
)

const (
	name = "adguard-exporter"
)

var (
	s *server.Server
)

func main() {
	conf := config.Load()

	metrics.Init()

	initAdguardClient(conf.AdguardProtocol, conf.AdguardHostname, conf.AdguardPort, conf.AdguardUsername, conf.AdguardPassword, conf.Interval)
	initHttpServer(conf.Port)

	handleExitSignal()
}

func initAdguardClient(protocol, hostname string, port uint16, username, password string, interval time.Duration) {
	client := adguard.NewClient(protocol, hostname, port, username, password, interval)
	go client.Scrape()
}

func initHttpServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	s.Stop()
	fmt.Println(fmt.Sprintf("\n%s HTTP server stopped", name))
}
