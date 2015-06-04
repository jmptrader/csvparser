package csvparser_test

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

func (*MyDataProcessor) OnDone(rows [][]string) {
	logger.Printf("MyDataProcessor: OnDone()")
	for i, v := range rows {
		logger.Printf("%v: %v", i, v)
	}
}

func (*MyDataProcessor) OnError(err error) {
	logger.Printf("MyDataProcessor: OnError(): %v", err)
}

func Example() {
	logger.Printf("NewCSVParser()...")

	p := csvparser.NewCSVParser("./data/1.csv", ",", &MyDataProcessor{})

	p.Start()
	// Output:
}

func init() {
	logger = fnlog.New("")
}
