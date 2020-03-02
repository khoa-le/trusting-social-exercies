package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func WriteFile(fileName string, data [][]string) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	err = csvWriter.WriteAll(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	csvWriter.Flush()
	return nil
}

func OpenAndReadFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file")
	}

	rows, err := readFile(file)
	if err != nil {
		fmt.Println("Failed to read file")
	}
	return rows
}
func readFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rows, nil
}
