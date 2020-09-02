package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func RunSocket(port string) {
    go RecvMessageQ(FOREVER, GetData)
    l, err := net.Listen("tcp", port)
    if nil != err {
        log.Println(err)
    }
    defer l.Close()

    for {
        conn, err := l.Accept()
        if nil != err {
            log.Println(err)
            continue
        }
        defer conn.Close()
        go ConnHandler(conn)
    }
}

func ConnHandler(conn net.Conn) {
    recvBuf := make([]byte, 4096)
    conn.Write([]byte("asdf"))
    for {
        n, err := conn.Read(recvBuf)
        if nil != err {
            log.Println(err)
            return
        }
        if n > 0 {
            data := recvBuf[:n]
            log.Println(string(data))
            ParseDataAndSend(data, conn)
        }
    }
}

func ParseDataAndSend(data []byte, conn net.Conn) {
    var recv map[string]interface{}
    send := make(map[string]interface{})
    json.Unmarshal(data, &recv)

    if recv["name"] == "kim" {
        send["answer"] = "no"
    }

    tx, _ := json.Marshal(send)
    conn.Write(tx)

    msg := QMessage{}
    msg.Cmd = 1
    msg.Data1 = ConvertByteArray("asdfasdfsadfsadfasdfasdfasdfasdfasdfasdf")
    msg.Data2 = ConvertByteArray("asdfasdfsadfsadfasdfasdfasdfasdfasdfasdf")
    msg.Data3 = ConvertByteArray("asdfasdfsadfsadfasdfasdfasdfasdfasdfasdf")
    SendMessageQ(msg)
}

func GetData(msg RecvMsg) {
    fmt.Println("WOWOWOWOWOWOWOWO")
    fmt.Println(msg)
}