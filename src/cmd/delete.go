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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "comando para borrar tarea segun el id especificado",
	Long:  `junto con el id elimina la tarea designada y crea un listado nuevo`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(strings.Join(args, " "))
		if err != nil {
			fmt.Println("ingrese un valor numerico")
			return
		}
		service.UpdateTasks(id, "del")

		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
