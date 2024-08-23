package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	StaticDir string `json:"static_dir"`
}

func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal("Ошибка чтения конфигурационного файла:", err)
	}

	fs := http.FileServer(http.Dir(config.StaticDir))
	http.Handle("/", http.StripPrefix("/", fs))

	addr := config.Host + ":" + strconv.Itoa(config.Port)
	log.Printf("Сервер запущен на http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
