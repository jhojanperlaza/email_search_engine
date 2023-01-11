package main

import (
	"github.com/jhojanperlaza/email_search_engine/tree/main/indexer_v2/functions"
	"os"
)

func ProfilingTest(name_dir string, curt_path string) int {
	//Declare environment name of bd
	os.Setenv("name_bd", "bd_v3")
	functions.Browser_dirs(name_dir, curt_path)
	return 1
}
