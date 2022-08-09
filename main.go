package main

import (
	"fmt"
	"path"
)

func main() {
	var _, file string
	_, file = path.Split("css/main.css")
	fmt.Println("file : ", file)

}
