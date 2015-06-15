package csvparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	DEBUG bool = false // Set it to true if you want to enable debug messages from this package.
)

// User should implement this interface to process items of each line.
type DataProcessor interface {
	OnDone(rows [][]string)                              // Callback function on all rows processed. Rows: all rows processed.
	OnError(err error)                                   // Callback function on error.
	ProcessLineItems(items []string, currentLine uint64) // Callback function to process each row(line) of CSV file.
}

// CSVParser struct
type CSVParser struct {
	file      string        // File name
	sep       string        // Separator
	processor DataProcessor // DataProcessor interface
}

// NewCSVParser() creats a CSVParser
//
//   Params:
//       file: File name.
//       sep: Separator of input file. Default CSV separator is ','.
//       processor: User should implement DataProcessor interface to provide the function to process items of each line.
//   Return:
//       *CSVParser
func NewCSVParser(file, sep string, processor DataProcessor) (p *CSVParser) {
	p = &CSVParser{file, sep, processor}
	return p
}

// Start() starts parsing the CSV file.
func (p *CSVParser) Start() error {
	var n uint64 = 0
	if DEBUG {
		fmt.Printf("p.file = %v", p.file)
	}
	file, err := os.Open(p.file)
	if err != nil {
		if DEBUG {
			fmt.Printf("os.Open(%s) err: %s\n", p.file, err)
		}
		p.processor.OnError(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := ""
	items := []string{}
	rows := [][]string{}
	for scanner.Scan() {
		line = scanner.Text()
		items = strings.Split(line, p.sep)
		n++
		p.processor.ProcessLineItems(items, n)
		rows = append(rows, items)
	}

	if err := scanner.Err(); err != nil {
		if DEBUG {
			fmt.Printf("scanner.Err(): %s\n", err)
		}
		p.processor.OnError(err)
		return err
	}

	p.processor.OnDone(rows)
	return nil
}
