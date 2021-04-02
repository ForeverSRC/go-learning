package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const storageRoot = "STORAGE_ROOT"

var methods = map[string]func(w http.ResponseWriter, r *http.Request){
	http.MethodPut: put,
	http.MethodGet: get,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	method := methods[m]
	if method == nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	method(w, r)
}

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(getObjName(r))
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(getObjName(r))
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func getObjName(r *http.Request) string {
	return os.Getenv(storageRoot) + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2]
}
