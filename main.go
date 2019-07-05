package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	goji "goji.io"
	"goji.io/pat"
)

type Student struct {
	Name string "json:name"
	Age  int    "json:age"
}

var allstudentmap = map[string]Student{}

func main() {

	allstudentmap["joey"] = Student{"joey", 39}
	allstudentmap["tom"] = Student{"tom", 20}

	rootmux := goji.NewMux()
	rootmux.HandleFunc(pat.Get("/Students"), allstudentHandler)
	rootmux.HandleFunc(pat.Get("/Student/:name/:age"), StudentHandler)
	rootmux.HandleFunc(pat.Put("/NewStudent"), newStudentHandler)
	http.ListenAndServe(":8089", rootmux)

	time.Sleep(8 * time.Hour)
}

func allstudentHandler(w http.ResponseWriter, r *http.Request) {
	allstudents := []Student{}

	for _, v := range allstudentmap {
		allstudents = append(allstudents, v)
	}

	studentjson, _ := json.Marshal(allstudents)

	fmt.Fprintf(w, string(studentjson))
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	studentName := pat.Param(r, "name")
	studentAge := pat.Param(r, "age")

	fmt.Println(studentAge)
	student := allstudentmap[studentName]

	studentjson, _ := json.Marshal(student)

	fmt.Fprintf(w, string(studentjson))
}

func newStudentHandler(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)

	var newStudent Student

	json.Unmarshal(b, &newStudent)

	allstudentmap[newStudent.Name] = newStudent

	fmt.Fprintf(w, string(b))

}
