package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"

	"task/internal/types"
)

const fileName string = "data.csv"

func ReadFile() ([][]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func SaveFile(args []string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	data := types.Todo{
		Task:      strings.Join(args, " "),
		Done:      false,
		CreatedAt: time.Now(),
	}

	record := []string{
		data.Task,
		strconv.FormatBool(data.Done),
		data.CreatedAt.Format(time.RFC3339),
	}
	writer.Write(record)

	defer writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
