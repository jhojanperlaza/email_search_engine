/*This package contains all the necessary
functions to index files in nd_json format
and upload them to the ZincSearch database.*/
package functions

import (
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

/*function that takes the name of the files
present in the directory and creates the ndjson
format for the data*/
func To_ndjson(names_files []string, path string) {

	split_index := strings.Split(path, "/")
	var name_index string

	if len(split_index) >= 2 {
		name_index1 := split_index[len(split_index)-2]
		// index name cannot start with "_"
		name_index1 = strings.TrimPrefix(name_index1, "_")
		name_index = name_index1 + "." + split_index[len(split_index)-1]
	} else {
		name_index = split_index[len(split_index)-1]
		// index name cannot start with "_"
		name_index = strings.TrimPrefix(name_index, "_")
	}

	var cont int64 = 0
	for _, name_file := range names_files {

		MyFile, err := os.Stat(path + "/" + name_file)
		if err != nil {
			fmt.Println("File does not exist")
		}
		cont += MyFile.Size()
	}

	if cont > 700000 {
		chunkSlice(names_files, len(names_files)/2, path)
		return
	}

	//build dictionary
	dict := make(map[string]string)

	for _, name := range names_files {

		content, err := ioutil.ReadFile(path + "/" + name)
		HandleErr(err)
		//convert to string
		str_content := string(content)

		dict[name_index+"."+name] = str_content
	}

	to_json1, err := json.Marshal(dict)
	HandleErr(err)

	Post_zincsearch(to_json1)
}

//function that halves the number of files found in a directory
func chunkSlice(slice []string, chunkSize int, path string) {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	for _, chunk := range chunks {
		To_ndjson(chunk, path)
	}
}

//function that uploads the document in json format to Zincsearch
func Post_zincsearch(dict_post []byte) {

	h := http.Client{}
	path_post := "http://localhost:4080/api/" + os.Getenv("name_bd") + "/_doc"
	req, err := http.NewRequest("POST", path_post, strings.NewReader(string(dict_post)))
	HandleErr(err)

	req.SetBasicAuth("admin", "Complexpass#123")
	r, err := h.Do(req)
	HandleErr(err)
	defer r.Body.Close()
}
