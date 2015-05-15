package csvparser

import (
	"testing"
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

func TestCSVParser(t *testing.T) {
	logger.Printf("NewCSVParser()...")

	p := NewCSVParser("./1.csv", ",", &MyDataProcessor{})

	p.Start()
}
