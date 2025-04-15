package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"task/internal/types"
)

const fileName string = "data.csv"

func ReadFile() ([]types.Todo, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return []types.Todo{}, err
	}

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		return []types.Todo{}, err
	}

	var todos []types.Todo

	for _, record := range records {
		done, err := strconv.ParseBool(record[1])
		if err != nil {
			return []types.Todo{}, err
		}
		createdAt, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			return []types.Todo{}, err
		}

		todo := types.Todo{
			Task:      record[0],
			Done:      done,
			CreatedAt: createdAt,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func SaveTodo(todo types.Todo) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	record := []string{
		todo.Task,
		strconv.FormatBool(todo.Done),
		todo.CreatedAt.Format(time.RFC3339),
	}

	writer.Write(record)

	defer writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

func SaveTodoAll(todos []types.Todo) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	var records [][]string

	for _, todo := range todos {
		record := []string{
			todo.Task,
			strconv.FormatBool(todo.Done),
		todo.CreatedAt.Format(time.RFC3339),
		}
		records = append(records, record)
	}

	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
