package csvparser_test

import (
	"fmt"
	"github.com/northbright/csvparser"
)

type MyDataProcessor struct {
}

func (*MyDataProcessor) ProcessLineItems(items []string, currentLine uint64) {
	fmt.Printf("%v: %v\n", currentLine, items)
}

func (*MyDataProcessor) OnDone(rows [][]string) {
	fmt.Printf("MyDataProcessor: OnDone()\n")
	for i, v := range rows {
		fmt.Printf("%v: %v\n", i, v)
	}
}

func (*MyDataProcessor) OnError(err error) {
	fmt.Printf("MyDataProcessor: OnError(): %v\n", err)
}

func Example() {
	fmt.Printf("NewCSVParser()...\n")
	p := csvparser.NewCSVParser("./data/1.csv", ",", &MyDataProcessor{})
	p.Start()
	// Output:
	// NewCSVParser()...
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
