package csvparser

import (
	"testing"
)

type MyDataProcessor struct {
}

func (*MyDataProcessor) ProcessLineItems(items []string, currentLine uint64) {
	logger.Printf("%v: %v", currentLine, items)
}

func (*MyDataProcessor) onDone() {
	logger.Printf("MyDataProcessor: onDone()")
}

func (*MyDataProcessor) onError(err error) {
	logger.Printf("MyDataProcessor: onError(): %v", err)
}

func TestCSVParser(t *testing.T) {
	logger.Printf("NewCSVParser()...")

	p := NewCSVParser("./1.csv", ",", &MyDataProcessor{})

	p.Start()
}
