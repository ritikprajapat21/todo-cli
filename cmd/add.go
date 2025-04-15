/*
Copyright © 2025 RITIK PRAJAPAT <ritikprajapati084@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/internal/types"
	"task/internal/utils"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add a new task",
	Long:  `To add a new task`,
	Run:   addTodo,
}

func addTodo(cmd *cobra.Command, args []string) {
	data := types.Todo{
		Task:      strings.Join(args, " "),
		Done:      false,
		CreatedAt: time.Now(),
	}

	err := utils.SaveTodo(data)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Task added")
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
