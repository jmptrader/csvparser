
# CSVParser

#### About  
CSVParser parses the CSV file and will pass the items of each line to user's own function to process the data.  
To do this, apps that use CSVParser should implement the `DataProcessor` interface:  

    // User should implement this interface to process items of each line.
    type DataProcessor interface {
        OnDone(rows [][]string)
        OnError(err error)
        ProcessLineItems(items []string, currentLine uint64)
    }


#### Example  

    package main

    import (
        "fmt"
        "github.com/northbright/csvparser"
    )

    // Implement DataProcessor interface:
    // ProcessLineItems(), OnDone(), OnError()
    type MyDataProcessor struct {
    }

    // Put your own business logic here
    func (*MyDataProcessor) ProcessLineItems(items []string, currentLine uint64) {
        fmt.Printf("%v: %v\n", currentLine, items)
    }

    func (*MyDataProcessor) OnDone(rows [][]string) {
        fmt.Printf("MyDataProcessor: OnDone()\n")
    }

    func (*MyDataProcessor) OnError(err error) {
        fmt.Printf("MyDataProcessor: OnError(): %v\n", err)
    }

    func main() {
        p := csvparser.NewCSVParser("./1.csv", ",", &MyDataProcessor{})
        p.Start()  // start parsing
    }

#### LICENSE
MIT License
