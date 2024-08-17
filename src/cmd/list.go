/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hmap55/toDoList/service"
	"github.com/spf13/cobra"
)

var Source bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lista las tareas creadas solo con su atributo de creacion",
	Long: `Lista las tareas almacenadas en el sistema solo teniendo en cuenta las que estan faltnates 
	por completar y sin mostrar el atributo de estado`,
	Run: func(cmd *cobra.Command, args []string) {

		if Source {
			service.ListTasksWithFlag()
		} else {
			service.ListTasks()
		}

		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.Flags().BoolVarP(&Source, "all", "a", false, "activar listado detallado")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
