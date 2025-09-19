package main

import (
    "flag"
    "log"
    "net/url"
    "os"

    "github.com/gorilla/websocket"
)

func main() {
    addr := flag.String("addr", "localhost:8080", "http service address")
    flag.Parse()

    u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
    log.Printf("connecting to %s", u.String())

    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatalf("dial error: %v", err)
    }
    defer c.Close()

    // read messages until exit
    for {
        _, message, err := c.ReadMessage()
        if err != nil {
            log.Printf("read error: %v", err)
            os.Exit(0)
        }
        log.Printf("recv: %s", message)
    }
}
