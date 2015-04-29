
# CSVParser

#### About  
CSVParser parses the CSV file and will pass the items of each line to user's own function to process the data.  
To do this, apps that use CSVParser should implement the `DataProcessor` interface:  

    // User should implement this interface to process items of each line.
    type DataProcessor interface {
        onDone()
        onError(err error)
        ProcessLineItems(items []string, currentLine uint64)
    }


#### Example  

    package main

    import (
        "fmt"
        "github.com/northbright/csvparser"
    )

    // Implement DataProcessor interface:
    // ProcessLineItems(), onDone(), onError()
    type MyDataProcessor struct {
    }

    // Put your own business logic here
    func (*MyDataProcessor) ProcessLineItems(items []string, currentLine uint64) {
        fmt.Printf("%v: %v", currentLine, items)
    }

    func (*MyDataProcessor) onDone() {
        fmt.Printf("MyDataProcessor: onDone()")
    }

    func (*MyDataProcessor) onError(err error) {
        fmt.Printf("MyDataProcessor: onError(): %v", err)
    }

    func main() {
        p := csvparser.NewCSVParser("./1.csv", ",", &MyDataProcessor{})
        p.Start()  // start parsing
    }

#### LICENSE
MIT License