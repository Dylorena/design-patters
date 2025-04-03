package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	filename := "output.csv"

	csv := NewCSVBuilder().setFilename(filename).AddRow([]string{"Name", "Age", "City"}).Build()

	csv.saveToFile()

}

type writerBuilder struct {
	w        *csv.Writer
	filename string
}

func NewCSVBuilder() *writerBuilder {
	return &writerBuilder{
		w: csv.NewWriter(os.Stdout),
	}
}

func (w *writerBuilder) setDelimiter(d rune) *writerBuilder {
	w.w.Comma = d
	return w
}

func (w *writerBuilder) setFilename(fn string) *writerBuilder {
	w.filename = fn
	return w
}

func (w *writerBuilder) AddRow(row []string) *writerBuilder {
	w.w.Write(row)
	return w
}

func (w *writerBuilder) Build() *writerBuilder {
	w.w.Flush()
	if err := w.w.Error(); err != nil {
		panic(err)
	}

	return w
}

func (w *writerBuilder) saveToFile() {
	file, err := os.Create(w.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("Success, export to %s", w.filename)

}
