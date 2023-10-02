package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var timeoutError = "Timeout while connecting to server"

var sTimeout string

func init() {
	flag.StringVar(&sTimeout, "timeout", "10s", "enter connection timeout")
}

var wg *sync.WaitGroup

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 1 {
		log.Fatal("Please enter server address and port")
	}

	host := args[0]
	port := args[1]
	timeout, err := time.ParseDuration(sTimeout)
	if err != nil {
		log.Fatal("Please enter correct timeout value")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		os.Exit(1)
	}()

	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)

	err = client.Connect()
	if err != nil {
		log.Fatal(timeoutError)
	}
	fmt.Fprintf(os.Stderr, "...Connected to %v\n", net.JoinHostPort(host, port))

	wg = &sync.WaitGroup{}
	wg.Add(2)
	go readRoutine(client)
	go writeRoutine(client)
	wg.Wait()
}

func readRoutine(telnetClient TelnetClient) {
	defer wg.Done()
	if err := telnetClient.Receive(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}

func writeRoutine(telnetClient TelnetClient) {
	defer wg.Done()
	if err := telnetClient.Send(); err != nil {
		fmt.Fprintf(os.Stderr, "...Connection was closed by peer\n")
		return
	}
	fmt.Fprintf(os.Stderr, "...EOF\n")
	telnetClient.Close()
}
