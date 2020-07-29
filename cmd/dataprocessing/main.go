package main

import (
    "github.com/jyotishp/go-orders/pkg/storage"
    "log"
)

func main() {
    csvFilePath := "sample.csv"
    outputDir := "outputs"
    csvReader := storage.NewCsvReader(csvFilePath, true)
    dh := storage.NewDataHandler(outputDir)
    data, done := csvReader.ReadLine()
    if done {
        log.Fatalln("Got an empty file")
    }
    dh.Init(data)

    for !done {
        data, done = csvReader.ReadLine()
        dh.Write(data)
    }
    dh.Close()
}
