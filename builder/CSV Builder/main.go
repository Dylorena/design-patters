package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	err := NewCSVBuilder().
		setFilename("dados.csv").
		setDelimiter(';').
		SetHeader("Nome", "Idade", "Profissão").
		AddRow("Alice", "25", "Engenheira").
		AddRow("Bob", "30", "Desenvolvedor").
		AddRow("Carlos", "28", "Designer").
		Build().
		SaveToFile()

	if err != nil {
		fmt.Println("Error to create filename.", err)
	}

	err = NewCSVBuilder().
		setFilename("abril-2025.csv").
		SetHeader("Data", "Descrição", "Valor").
		AddRow("12/04", "Mercado", "537,88").
		AddRow("13/04", "Lanche", "15,90").
		AddRow("14/04", "Energia", "100").
		Build().
		SaveToFile()

	if err != nil {
		fmt.Println("Error to create filename.", err)
	}

	err = NewCSVBuilder().
		SetHeader("Data", "Descrição", "Valor").
		AddRow("12/04", "Mercado", "537,88").
		AddRow("13/04", "Lanche", "15,90").
		AddRow("14/04", "Energia", "100").
		Build().
		SaveToFile()

	if err != nil {
		fmt.Println("Error to create filename.", err)
	}

}

type mapOfString [][]string

type writerBuilder struct {
	filename  string
	delimiter rune
	header    []string
	rows      [][]string
	built     mapOfString
}

func NewCSVBuilder() *writerBuilder {
	return &writerBuilder{
		filename:  "output.csv",
		delimiter: ',',
	}
}

func (w *writerBuilder) setDelimiter(d rune) *writerBuilder {
	w.delimiter = d
	return w
}

func (w *writerBuilder) setFilename(fn string) *writerBuilder {
	w.filename = fn
	return w
}

func (w *writerBuilder) AddRow(values ...string) *writerBuilder {
	w.rows = append(w.rows, values)
	return w
}

func (w *writerBuilder) SetHeader(columns ...string) *writerBuilder {
	w.header = columns
	return w
}

func (w *writerBuilder) Build() *writerBuilder {

	w.fillHeader()

	w.fillRows()

	return w
}

func (w *writerBuilder) fillHeader() {
	if len(w.header) > 0 {
		w.built = append(w.built, w.header)
	}
}

func (w *writerBuilder) fillRows() {
	for i, row := range w.rows {
		if len(w.header) > 0 && len(row) != len(w.header) {
			fmt.Printf("Something don't work well with row %d, jump to next row", i+1)
			continue
		}
		w.built = append(w.built, row)
	}
}

func (w *writerBuilder) SaveToFile() error {
	file, err := os.Create(w.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = w.delimiter

	err = writer.WriteAll(w.built)
	if err != nil {
		return err
	}

	writer.Flush()
	fmt.Printf("Success, export to: %s\n", w.filename)
	return nil
}
