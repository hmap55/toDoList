package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hmap55/toDoList/entities"
)

func AddTask(task entities.Task) {

	var tasks []entities.Task

	file, err := os.OpenFile("List.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al intentar crear o abrir el archivo: ", err)
		return
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println("Error al intentar obetener informacion: ", err)
		return
	}

	if info.Size() > 0 {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&tasks)
		if err != nil {
			fmt.Println("error al leer el archivo ", err)
			return
		}
		task.Id = tasks[len(tasks)-1].Id + 1
	}

	tasks = append(tasks, task)

	file.Seek(0, 0)
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("error procesando el archivo")
		return
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("error al guardar nueva tarea", err)
		return
	}

	fmt.Println("Tarea guardada con exito")

}

func ListTasks() {

	var tasks []entities.Task

	file, err := os.OpenFile("List.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error accediendo al archivo")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("error en lectura de archivo")
	}

	fmt.Printf("%-5s %-50s %-15s\n", "ID", "DESCRIPTION", "CREATED")

	for _, list := range tasks {
		if !list.Done {
			fmt.Printf("%-5d %-50s %-15s\n", list.Id, list.Description, list.Create)
		}
	}

}

func ListTasksWithFlag() {

	var tasks []entities.Task

	file, err := os.OpenFile("List.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error accediendo al archivo")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("error en lectura de archivo")
	}

	fmt.Printf("%-5s %-50s %-25s %-30s \n", "ID", "DESCRIPTION", "CREATED", "DONE")

	for _, list := range tasks {

		fmt.Printf("%-5d %-50s %-25s %-30v\n", list.Id, list.Description, list.Create, list.Done)

	}

}

func UpdateTasks(id int, operation string) {

	var tasks []entities.Task
	isPresent := false

	file, err := os.OpenFile("List.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error al intentar leer el archivo")
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("error convirtiendo el archivo")
	}

	for i, task := range tasks {
		if task.Id == id {
			if operation == "del" {
				tasks = append(tasks[:i], tasks[i+1:]...)
				isPresent = true

			}
			if operation == "upd" {
				tasks[i].Done = true
				isPresent = true
			}

		}
	}

	if !isPresent {
		fmt.Println("el id ingresado no se encuentra")
		return
	} else {
		file.Seek(0, 0)
		err = file.Truncate(0)
		if err != nil {
			fmt.Println("error procesando el archivo")
			return
		}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(tasks)
		if err != nil {
			fmt.Println("no se pudo escribir en el archivo")
			return
		}
		if operation == "del" {
			fmt.Printf("id %d eliminado con exito \n", id)
		} else {
			fmt.Printf("id %d completado con exito \n", id)
		}

	}
}
