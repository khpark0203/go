package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

func main() {
    conn, err := net.Dial("tcp", "10.15.111.117:8000")
    if nil != err {
        log.Println(err)
    }

    go func() {
        data := make([]byte, 4096)

        for {
            n, err := conn.Read(data)
            if err != nil {
                log.Println(err)
                return
            }

            log.Println("Server send : " + string(data[:n]))
        }
    }()

    for {
        data := make(map[string]interface{})
        data["name"] = "kim"
        data["age"] = 25
    
        doc, _ := json.Marshal(data)
    
        conn.Write(doc)
        time.Sleep(time.Second * 2)
    }
}