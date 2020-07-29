package storage

import (
    "encoding/json"
    "log"
    "os"
)

type JsonWriter struct {
    File *os.File
    FilePath string
}

func NewJsonWriter(filepath string, resource string) *JsonWriter {
    RemoveFile(filepath)
    writer := &JsonWriter{FilePath: filepath, File: CreateFile(filepath)}
    writer.WriteString("{\"" + resource + "\":[")
    return writer
}

func CreateFile(filepath string) *os.File {
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err.Error())
    }
    return file
}

func RemoveFile(filepath string) {
    if _, err := os.Stat(filepath); err != nil {
        os.Remove(filepath)
    }
}

func (w *JsonWriter) WriteString(txt string) {
    _, err := w.File.WriteString(txt)
    if err != nil {
        log.Fatalf("Failed to write to file: %v", err.Error())
    }
}

func (w *JsonWriter) WriteResource(v interface{}) {
    jsonData, err := json.Marshal(v)
    if err != nil {
        log.Fatalf("Error encoding to JSON: %v", err.Error())
    }
    w.WriteString("," + string(jsonData))
}

func (w *JsonWriter) InitWriteResource(v interface{}) {
    jsonData, err := json.Marshal(v)
    if err != nil {
        log.Fatalf("Error encoding to JSON: %v", err.Error())
    }
    w.WriteString(string(jsonData))
}

func (w *JsonWriter) Close() {
    w.WriteString("]}")
    w.File.Close()
}