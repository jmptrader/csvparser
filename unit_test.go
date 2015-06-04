package csvparser

import (
	"testing"
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

func TestCSVParser(t *testing.T) {
	logger.Printf("NewCSVParser()...")

	p := NewCSVParser("./data/1.csv", ",", &MyDataProcessor{})

	p.Start()
}
