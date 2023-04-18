package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FileUploadReponse struct {
	Message string `json:"message"`
	FileUrl string `json:"fileUrl"`
}

func ErrorReponse(w http.ResponseWriter, err error) http.ResponseWriter {
	response := FileUploadReponse{}
	response.Message = err.Error()
	response.FileUrl = ""
	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
	return w
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("FIle Uplaoded Syccessfully", r.Body)
	// fmt.Println("i am here ", r.MultipartForm.File)
	// if r.MultipartForm.File == nil {
	// 	fmt.Println("i AM NIL")
	// }
	response := FileUploadReponse{}

	file, fileHeader, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		// ErrorReponse(w, err)
		response.Message = err.Error()
		response.FileUrl = ""
		res, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	defer file.Close()

	folderPath := "./upload/"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			response.Message = err.Error()
			response.FileUrl = ""
			res, _ := json.Marshal(response)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)
		}
	}

	filepath := folderPath + fileHeader.Filename
	fmt.Println(filepath)

	newFile, err := os.Create(filepath)
	if err != nil {
		response.Message = err.Error()
		response.FileUrl = ""
		res, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		response.Message = err.Error()
		response.FileUrl = ""
		res, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	}

	response.Message = "FIle Uplaoded Syccessfully"
	response.FileUrl = filepath

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
