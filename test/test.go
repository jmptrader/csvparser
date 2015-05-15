package main

import (
	"github.com/northbright/csvparser"
	"github.com/northbright/fnlog"
	"log"
)

var (
	logger *log.Logger
)

type MyDataProcessor struct {
}

func (*MyDataProcessor) ProcessLineItems(items []string, currentLine uint64) {
	logger.Printf("%v: %v", currentLine, items)
}

func (*MyDataProcessor) OnDone() {
	logger.Printf("MyDataProcessor: OnDone()")
}

func (*MyDataProcessor) OnError(err error) {
	logger.Printf("MyDataProcessor: OnError(): %v", err)
}

func main() {
	logger.Printf("NewCSVParser()...")

	p := csvparser.NewCSVParser("./1.csv", ",", &MyDataProcessor{})

	p.Start()
}

func init() {
	logger = fnlog.New("")
}
