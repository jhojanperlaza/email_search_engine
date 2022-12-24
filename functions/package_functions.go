/*This package contains all the necessary
functions to index files in nd_json format
and upload them to the ZincSearch database.*/
package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Function that prints on console errors in execution process
func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*Function that recursively traverses directories
and files to create the indexes of the database.*/
func Browser_dirs(name_dir string, curt_path string) {

	curt_path += "/" + name_dir

	// Get the currents files
	files, err := ioutil.ReadDir(curt_path)
	HandleErr(err)

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
		To_ndjson(list_files, curt_path)
	}

	for _, dir := range list_dirs {
		Browser_dirs(dir, curt_path)
	}

	if len(list_dirs) == 0 {
		return
	}
}

//function that writes the data to a file 'bd_mails.ndjson'
func write_file(dict1 []byte, dict2 []byte) {

	if _, err := os.Stat("bd_mails.ndjson"); err == nil {
		//File exists
		f, err := os.OpenFile("bd_mails.ndjson", os.O_APPEND|os.O_WRONLY, 0660)
		HandleErr(err)
		str := string(dict1)
		_, err = fmt.Fprint(f, str, "\n")
		HandleErr(err)
		str2 := string(dict2)
		_, err = fmt.Fprint(f, str2, "\n")
		HandleErr(err)

		defer f.Close()

	} else {
		//File does not exist
		f, err := os.Create("bd_mails.ndjson")
		HandleErr(err)
		str := string(dict1)
		_, err = fmt.Fprint(f, str, "\n")
		HandleErr(err)
		str2 := string(dict2)
		_, err = fmt.Fprint(f, str2, "\n")
		HandleErr(err)

		defer f.Close()
	}
}

/*function that takes the name of the files
present in the directory and creates the ndjson
format for the data*/
func To_ndjson(names_files []string, path string) {

	split_index := strings.Split(path, "/")
	name_index := split_index[len(split_index)-1]

	//build of the first dictionary for the documents bulk format
	dict1 := map[string]map[string]string{
		"index": {
			"_index": name_index,
		},
	}

	to_json, err := json.Marshal(dict1)
	HandleErr(err)

	//build the second dictionary
	dict2 := make(map[string]string)

	for _, name := range names_files {

		content, err := ioutil.ReadFile(path + "/" + name)
		HandleErr(err)
		//convert to string
		str_content := string(content)

		dict2[name] = str_content
	}

	to_json2, err := json.Marshal(dict2)
	HandleErr(err)

	write_file(to_json, to_json2)
}

func Post_zincsearch() {
	file_found, err := ioutil.ReadFile("bd_mails.ndjson")
	HandleErr(err)

	h := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulk", bytes.NewBuffer(file_found))
	HandleErr(err)

	req.SetBasicAuth("admin", "Complexpass#123")
	r, err := h.Do(req)
	HandleErr(err)

	defer r.Body.Close()
}
