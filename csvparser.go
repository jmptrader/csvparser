package csvparser

import (
	"bufio"
	"github.com/northbright/fnlog"
	"log"
	"os"
	"strings"
)

var (
	logger *log.Logger
)

// User should implement this interface to process items of each line.
type DataProcessor interface {
	onDone()
	onError(err error)
	ProcessLineItems(items []string, currentLine uint64)
}

// CSVParser struct
// file: File name
// sep: Separator
// processor: DataProcessor interface
type CSVParser struct {
	file      string
	sep       string
	processor DataProcessor
}

// New a CSVParser
// Params:
//     file: File name.
//     sep: Separator of input file. Default CSV separator is ','.
//     processor: User should implement DataProcessor interface to provide the function to process items of each line.
// Return:
//     *CSVParser
func NewCSVParser(file, sep string, processor DataProcessor) (p *CSVParser) {
	p = &CSVParser{file, sep, processor}
	return p
}

// Start parsing CSV file
// Returns:
//     error
func (p *CSVParser) Start() error {
	var n uint64 = 0
	file, err := os.Open(p.file)
	if err != nil {
		logger.Printf("os.Open(%s) err: %s\n", p.file, err)
		p.processor.onError(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := ""
	items := []string{}
	for scanner.Scan() {
		line = scanner.Text()
		items = strings.Split(line, p.sep)
		n++
		p.processor.ProcessLineItems(items, n)
	}

	if err := scanner.Err(); err != nil {
		logger.Printf("scanner.Err(): %s\n", err)
		p.processor.onError(err)
		return err
	}

	p.processor.onDone()
	return nil
}

func init() {
	logger = fnlog.New("")
}