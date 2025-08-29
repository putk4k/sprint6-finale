package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос HandlerRoot")
	http.ServeFile(w, r, "index.html")

}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос HandlerUpload")
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Вызов функции конвертации")
	converted, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)

	filename := fmt.Sprintf("result_%s%s", time.Now().UTC().Format("20060102_150405"), ext)
	log.Println("Создание временного файла")
	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}

	defer out.Close()
	log.Println("Начало записи в файл")
	if _, err := out.WriteString(converted); err != nil {
		http.Error(w, "Ошибка при записи в файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Результат конвертации:\n" + converted))
}
