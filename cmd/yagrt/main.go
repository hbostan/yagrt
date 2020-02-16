package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/hbostann/yagrt"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Please provide the scene xml file")
	// }

	// yagrt.Render(os.Args[1])
	files, err := ioutil.ReadDir("./scenes")
	if err != nil {
		fmt.Println("Error while listing dir")
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".xml") {
			fmt.Printf("%v took ", f.Name())
			start := time.Now()
			yagrt.Render("./scenes/" + f.Name())
			elapsed := time.Since(start)
			fmt.Printf("%v \n", elapsed)
		}
	}
}
