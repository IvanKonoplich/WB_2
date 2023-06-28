package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	wget()
}

func wget() {
	flag.Parse() //получаем введенную ссылку
	if len(flag.Args()) <= 0 {
		log.Fatal("введите ссылку на ресурс")
	}
	url := flag.Arg(0)
	request, err := http.NewRequest("GET", url, nil) //создаем запрос
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //добавляем к запросу контекст, на случаи сбоев
	defer cancel()
	request.WithContext(ctx)
	response, err := http.DefaultClient.Do(request) //выполняем запрос
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 { //проверяем что запрос сработал корректно
		log.Fatal("Status Code:", response.Status)
	}
	//проверяем что пришло по запросу
	var fileName string
	if strings.Contains(response.Header.Get("Content-Type"), "text/html") {
		fileName = "index.html"
	} else {
		fileName = path.Base(request.URL.Path)
	}

	file, err := os.Create(fileName) //создаем файл для результата
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body) //пишем в файл
	if err != nil {
		log.Fatalln(err)
	}
}
