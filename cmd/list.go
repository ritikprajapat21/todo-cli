/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"task/internal/utils"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "To list all the tasks",
	Long: `To list all the tasks`,
	Run:   listTodo,
}

func listTodo(cmd *cobra.Command, args []string) {
	todos, err := utils.ReadFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, todo := range todos {
		prefix := "   "
		if todo.Done {
			prefix = "✔️ "
		}
		fmt.Printf("%v%d:\t %v\t %s\n", prefix, i+1, todo.Task, todo.CreatedAt)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
