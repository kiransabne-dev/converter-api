package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-chi/chi"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func ConvertFileRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", home)
	router.Post("/", uploadFile)
	return router
}

func uploadFile(res http.ResponseWriter, req *http.Request) {
	log.Println("File Upload Endpoint")
	req.ParseMultipartForm(10 << 20)

	file, handler, err := req.FormFile("File")
	if err != nil {
		log.Println("upload err -> ", err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//above file upload,
	// create a temp file in server and copy the data in it
	f, err := os.OpenFile("/home/kiran/Downloads/uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("copy file -> ", err)
	}
	defer f.Close()
	io.Copy(f, file)

	//convert docx to pdf
	arg0 := "soffice"
	arg1 := "--invisible" //This command is optional, it will help to disable the splash screen of LibreOffice.
	arg2 := "--convert-to"
	arg3 := "pdf:writer_pdf_Export"
	path := "/home/kiran/Downloads/uploads/" + handler.Filename
	nout, err := exec.Command(arg0, arg1, arg2, arg3, path).Output()
	if err != nil {
		fmt.Println("nout err -> ", err)
	}
	fmt.Println(nout)

}

func home(res http.ResponseWriter, req *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "painting"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("home page controller")

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}
