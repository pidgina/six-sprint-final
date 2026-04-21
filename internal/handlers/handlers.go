package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", r.Method),
			http.StatusMethodNotAllowed)
		return
	}

	textFile, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(textFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func DownloaderHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", r.Method),
			http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	byteFile, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	result := []byte(service.Convert(string(byteFile)))

	time1 := time.Now().UTC().Format("02-01-2006 15-04-05")

	err = os.WriteFile(time1+filepath.Ext(handler.Filename), result, 0644)
	if err != nil {
		http.Error(w, "ошибка при заполнении файла", http.StatusInternalServerError)
		return
	}

	w.Write(result)
}
