package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const ERROR_MENT = "Select protocol -> https or socket"

func main() {
    if len(os.Args) < 2 {
        fmt.Println(ERROR_MENT)
        return
    } else {
        // port := ":" + GetPortNumber()
        port := ":8000"
        if len(os.Args) == 3 {
            port = ":" + os.Args[2]
        }

        switch os.Args[1] {
        case "socket":
            RunSocket(port)
        case "https":
            RunHttps(port)
        default:
            fmt.Println(ERROR_MENT)
        }
    }
}

func GetPortNumber() string {
    db, err := sql.Open("sqlite3", DB_PATH)
    if err != nil {
        panic(err)
    }

    rows, _ := db.Query("SELECT DocName FROM RECD_JOBLOG WHERE NO=4")
    var port string
    for rows.Next() {
        rows.Scan(&port)
    }
    db.Close()

    return port
}