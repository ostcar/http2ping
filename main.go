package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	pingWait := flag.Int("wait", 5, "Time in seconds to wait between pings")

	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		log.Fatal("No url given")
	}

	client, body, err := connect(url)
	if err != nil {
		log.Fatalf("Can not connect to server: %v", err)
	}
	defer body.Close()

	go func() {
		scanner := bufio.NewScanner(body)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Can not read responce body: %v", err)
		}
		log.Printf("Connection closed")
	}()

	// Wait for a short time to wait for showing the first data on stdout.
	time.Sleep(time.Second)

	for {
		go func() {
			fmt.Println("Sending Ping...")
			if err := client.Ping(context.Background()); err != nil {
				log.Fatalf("Ping error: %v", err)
			}
			fmt.Println("Pong received :)")
		}()
		time.Sleep(time.Duration(*pingWait) * time.Second)
	}
}

func connect(addr string) (*http2.ClientConn, io.ReadCloser, error) {
	t := new(http2.Transport)

	url, err := url.Parse(addr)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing addr: %w", err)
	}

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"h2"},
	}

	tconn, err := tls.Dial("tcp", url.Host, tlsconfig)
	if err != nil {
		return nil, nil, fmt.Errorf("connect tls: %w", err)
	}

	client, err := t.NewClientConn(tconn)
	if err != nil {
		return nil, nil, fmt.Errorf("create new http2 client: %w", err)
	}

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("creatung request: %w", err)
	}

	resp, err := client.RoundTrip(req)
	if err != nil {
		return nil, nil, fmt.Errorf("start connection: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Server responded with status %s", resp.Status)
	}

	return client, resp.Body, nil
}
