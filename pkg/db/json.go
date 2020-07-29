package db

import (
    "bufio"
    "github.com/tamerh/jsparser"
    "log"
    "os"
)

func JsonHandle(filepath, rootObj string) *jsparser.JsonParser {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatalf("Error opening JSON DB: %v", err.Error())
    }
    bufReadr := bufio.NewReaderSize(file, 65536)
    handle := jsparser.NewJSONParser(bufReadr, rootObj)
    return handle
}