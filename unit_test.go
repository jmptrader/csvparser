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
	// NewCSVParser()...
	// p.file = ./data/1.csv
	// 1: [iphone4s 8GB  3  2015/04/03]
	// 2: [iphone5s 16GB  1  2015/05/05]
	// 3: [iphone4s 8GB  20  2015/05/06]
	// 4: [iphone5s 32GB  8  2015/05/06]
	// 5: [iphone5s 64GB  1  2015/05/06]
	// 6: [iphone4s 8GB  1  2015/05/08]
	// MyDataProcessor: OnDone()
	// 0: [iphone4s 8GB  3  2015/04/03]
	// 1: [iphone5s 16GB  1  2015/05/05]
	// 2: [iphone4s 8GB  20  2015/05/06]
	// 3: [iphone5s 32GB  8  2015/05/06]
	// 4: [iphone5s 64GB  1  2015/05/06]
	// 5: [iphone4s 8GB  1  2015/05/08]

}

func init() {
	logger = fnlog.New("", true, false, false)
}
