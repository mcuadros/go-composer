package main

import (
    "fmt"
    "log"
    "net"
    "sync"
)

type Server struct {
    addr *net.UDPAddr
    conn *net.UDPConn
    wg *sync.WaitGroup
    error error
}

func (s *Server) Connect() bool {
    s.addr, s.error = net.ResolveUDPAddr("udp", ":52741")
    if s.error != nil {
        log.Fatalln(s.error)
    }

    s.conn, s.error = net.ListenUDP("udp", s.addr)
    if s.error != nil {
        log.Fatalln(s.error)
    }

    fmt.Println("listening on ", s.conn.LocalAddr().String())
    return true
}

func (s *Server) Start() bool {
    for {
        buf := make([]byte, 1024)
        n, err := s.conn.Read(buf)
        if err != nil {
            log.Fatalln(err)
        }

        fmt.Println("server: read:", string(buf[0:n]))
    }

    return true
}

func main() {
    server := &Server{}
    server.Connect()
    server.Start()
}