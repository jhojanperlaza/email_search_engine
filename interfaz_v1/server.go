package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func FunRequest(keyCharacters string) map[string]string {
	//Assigned to: CN=Sandra F Brawner
	//Sandra_Brawner_Dec2000
	query := fmt.Sprintf(`{"query": {"match": {"_all": "%s" }}, "size":10}`, keyCharacters)
	req, err := http.NewRequest("POST", "http://localhost:4080/es/bd_v3/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	x := make(map[string]interface{})

	json.Unmarshal(body, &x)
	DataToFront := DataProcessing(x, keyCharacters)
	return DataToFront
}

func DataProcessing(data map[string]interface{}, keyCharacters string) map[string]string {

	var emailsData string
	DataReturn := make(map[string]string)
	DictAux := data["hits"]
	LenData := 0

	for key, value := range DictAux.(map[string]interface{}) {
		if key == "hits" {
			DictAux = value
			LenData = len(value.([]interface{}))
			break
		}
	}
	fmt.Println(LenData)
	if LenData <= 0 {
		return map[string]string{"error.error.error": "****************no match found****************"}
	}

	for i := 0; i <= LenData-1; i++ {
		for key, value := range DictAux.([]interface{})[i].(map[string]interface{}) {
			if key == "_source" {
				for key, value := range value.(map[string]interface{}) {
					emailsData, _ = value.(string)
					//magic
					if strings.Contains(emailsData, keyCharacters) {
						emailsData = strings.Replace(emailsData, "\n", "<br>", -1)
						emailsData = strings.Replace(emailsData, keyCharacters, "<mark>"+keyCharacters+"</mark>", -1)
						DataReturn[key] = emailsData
					}
				}
			}
		}
	}
	return DataReturn
}

func main() {
	router := chi.NewRouter()
	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	router.Post("/api/searchQuery", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		s := string(body)
		DataToFront := FunRequest(s)
		jsonStr, _ := json.Marshal(DataToFront)
		w.Write([]byte(jsonStr))
	})

	FileServer(router)
	fmt.Println("Seeker is running in http://localhost:3000")
	panic(server.ListenAndServe())
}

// FileServer is serving static files.
func FileServer(router *chi.Mux) {
	root := "./interfaz_seeker/dist"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
