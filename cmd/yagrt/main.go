package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/hbostann/yagrt"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println("Error while reading dir")
	}
	if _, err := os.Stat("outputs"); os.IsNotExist(err) {
		os.Mkdir("outputs", 0755)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".xml") {
			start := time.Now()
			yagrt.Render(f.Name())
			elapsed := time.Since(start)
			fmt.Printf("%v took %v \n", f.Name(), elapsed)
		}
	}
}
