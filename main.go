package main

import (
	"fmt"
	"github.com/jhojanperlaza/email_search_engine/functions"
	"io/ioutil"
	"os"
)

func main() {

	// Get the name of data base and current directory
	name_bd := os.Args
	if len(name_bd) < 2 {
		panic("unspecified database")
	}
	current_Path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	current_Path += "/" + name_bd[1]

	// Get the currents files
	files, err := ioutil.ReadDir(current_Path)
	if err != nil {
		panic(err)
	}
	if len(files) < 1 {
		panic("No files found")
	}

	//lists of file and directory names
	var list_files []string
	var list_dirs []string

	for _, file := range files {

		if file.IsDir() {
			list_dirs = append(list_dirs, file.Name())
		} else {
			list_files = append(list_files, file.Name())
		}
	}

	if len(list_files) >= 1 {
		functions.To_ndjson(list_files, current_Path)
	}

	for _, dir := range list_dirs {
		functions.Browser_dirs(dir, current_Path)
	}

	functions.Post_zincsearch()
	fmt.Println("Database indexing done successfully!!!")
}
