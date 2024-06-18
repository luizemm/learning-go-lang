package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	/*
		No go, se tiver uma variável que não está sendo utilizada, o go não compila.
		Porém tem momentos que vai ocorrer uma situação que não queremos
		utilizar um dos valores que foi retornado por uma função, por exemplo.
		Para isso é utilizado o underline "_"
	*/
	res, _ := http.Get("https://google.com")
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(string(body))
}
