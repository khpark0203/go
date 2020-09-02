package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	posixMQ "github.com/syucream/posix_mq"
)

const (
    FOREVER            = 0
    STRING_MAX_LEN     = 30
    SEND_MESSAGE_PATH  = "/server_to_sys"
    RECV_MESSAGE_PATH  = "/sys_to_server"
    DB_PATH            = "/app/nvram/karas.db"
)

type QMessage struct {
    Cmd		uint32
    Data1	[STRING_MAX_LEN]uint8
    Data2	[STRING_MAX_LEN]uint8
    Data3	[STRING_MAX_LEN]uint8
}

type RecvMsg struct {
    cmd     uint32
    data1	string
    data2	string
    data3	string
}

func SendMessageQ(msg QMessage) {
    oflag := posixMQ.O_WRONLY
    mq, err := posixMQ.NewMessageQueue(SEND_MESSAGE_PATH, oflag, 0666, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer mq.Close()

    buf := &bytes.Buffer{}
    binary.Write(buf, binary.LittleEndian, msg)

    mq.Send(buf.Bytes(), 0)
}

func RecvMessageQ(loop int, Callback func(msg RecvMsg)) {
    oflag := posixMQ.O_RDONLY | posixMQ.O_CREAT
    fmt.Println("Start receiving messages")
    for {
        mq, _ := posixMQ.NewMessageQueue(RECV_MESSAGE_PATH, oflag, 0666, nil)
        defer mq.Close()

        rxdata, _, _ := mq.Receive()
        Callback(MsgCpy(rxdata))

        if loop != FOREVER {
            break
        }
    }
}


func MsgCpy(data []byte) RecvMsg {
    byteMsg := QMessage{}
    msg := RecvMsg{}
    buf := &bytes.Buffer{}
    binary.Write(buf, binary.LittleEndian, data)
    binary.Read(buf, binary.LittleEndian, &byteMsg)
    msg.cmd = byteMsg.Cmd
    msg.data1 = ConvertString(byteMsg.Data1)
    msg.data2 = ConvertString(byteMsg.Data2)
    msg.data3 = ConvertString(byteMsg.Data3)
    return msg
}

func ConvertString(data [STRING_MAX_LEN]uint8) string {
    var str string
    for i := 0; i < len(data); i++ {
        str = str + string(data[i])
    }
    return str
}

func ConvertByteArray(cp string) [STRING_MAX_LEN]uint8 {
    var data [STRING_MAX_LEN]uint8
    maxLen := STRING_MAX_LEN - 1

    if maxLen + 1 > len(cp) {
        maxLen = len(cp)
    }

    for i := 0; i < maxLen; i++ {
        data[i] = cp[i]
    }

    data[maxLen] = 0
    return data
}