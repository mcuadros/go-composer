package main

import (
    "fmt"
    "log"
    "net"
    "sync"
)

func client(sAddr *net.UDPAddr, wg *sync.WaitGroup) {
    cAddr, err := net.ResolveUDPAddr("udp", ":52741")
    if err != nil {
        log.Fatalln(err)
    }
    cConn, err := net.DialUDP("udp", cAddr, sAddr)
    if err != nil {
        log.Fatalln(err)
    }
    buf := []byte("hellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohello")
    n, err := cConn.Write(buf)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("client: wrote:", string(buf[0:n]))
    err = cConn.Close()
    if err != nil {
        log.Fatalln(err)
    }
    wg.Done()
}


func main() {
    sAddr, err := net.ResolveUDPAddr("udp", ":0")
    if err != nil {
        log.Fatalln(err)
    }

    cAddr, err := net.ResolveUDPAddr("udp", ":52741")
    if err != nil {
        log.Fatalln(err)
    }

    cConn, err := net.DialUDP("udp", sAddr, cAddr )
    if err != nil {
        log.Fatalln(err)
    }
    buf := []byte("hellohellohellohellohellohellohedddddddddddddddddddddddddddddddddddddssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssllohellohellohellohellohellohellohellohellohellohellohellohellohellohelloh1111ddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd33dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd3322ddddddddddddddd22212345678ddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd111")
    n, err := cConn.Write(buf)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("client: wrote:", string(buf[0:n]))
    err = cConn.Close()
    if err != nil {
        log.Fatalln(err)
    }
}