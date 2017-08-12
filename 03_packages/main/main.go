package main

import (
	"fmt"
	"github.com/mauricelacerda/GolangTraining/03_packages/nameutil"
)

func main()  {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, " + nameutil.MyName)
}