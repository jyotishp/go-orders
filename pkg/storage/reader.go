package storage

import (
    "encoding/csv"
    "log"
    "os"
)

type CsvReader struct {
    FilePath string
    File *os.File
    Csv *csv.Reader
}

func NewCsvReader(filepath string, header bool) *CsvReader {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatalf("Failed to open CSV: %v", err.Error())
    }
    r := csv.NewReader(file)
    reader := &CsvReader{
        FilePath: filepath,
        File:     file,
        Csv:      r,
    }
    if header {
        reader.ReadLine()
    }
    return reader
}

func (r *CsvReader) ReadLine() ([]string, bool) {
    data, err := r.Csv.Read()
    if data == nil {
        return data, true
    }
    if err != nil {
        log.Fatalf("Failed to read CSV: %v", err.Error())
    }
    return data, false
}

func (r *CsvReader) Close() {
    r.File.Close()
}