package csv

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func CreateCSVFileWithDatabaseData(filename string, data [][]string) (err error) {
	var file *os.File

	_, s := os.Stat(filename)
	switch s {
	case nil:
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
	default:
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func OpenAndEncodeCSV(filename string) (encodecsv string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to open CSV file: "+err.Error())
	}
	defer file.Close()

	// Read the contents of the CSV file
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to get file information: "+err.Error())
	}
	fileSize := fileInfo.Size()
	csvFileContents := make([]byte, fileSize)
	_, err = file.Read(csvFileContents)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to read CSV file: "+err.Error())
	}
	encodecsv = base64.StdEncoding.EncodeToString(csvFileContents)
	return
}

func CSVLoader(path string) [][]string {
	op, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cant open the csv : %s || %s", path, err.Error())
	}
	defer op.Close()
	reader := csv.NewReader(op)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Cant Read the csv file : %s", err.Error())
	}
	return rows
}
