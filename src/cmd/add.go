/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/hmap55/toDoList/entities"
	"github.com/hmap55/toDoList/service"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Agrega una nueva tarea a la lista de tareas",
	Long:  `Crea una nueva tarea en el archivo que consigna las tareas y lo guarda en el mismo`,
	Run: func(cmd *cobra.Command, args []string) {
		id := 1
		task := entities.NewTask(id, strings.Join(args, " "), time.Now().Format("2006-01-02 15:04:05"), false)
		service.AddTask(*task)
		fmt.Println("add called")

	},
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
