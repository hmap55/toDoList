/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hmap55/toDoList/service"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "comando que al enviar el id de la tarea la marca como completada",
	Long:  `completar tarea segun el id ingresado, marcar como culminada una tarea`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(strings.Join(args, " "))
		if err != nil {
			fmt.Println("ingrese un valor numerico")
			return
		}
		service.UpdateTasks(id, "upd")
		fmt.Println("complete called")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
