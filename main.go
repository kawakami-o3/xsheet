package main

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	files := os.Args[1:]

	for _, f := range files {
		xls, err := excelize.OpenFile(f)
		if err != nil {
			fmt.Printf("%s,error\n", f)
			fmt.Println(err)
			continue
		}

		for i:=1 ; i<=len(xls.GetSheetMap()); i++ {
			fmt.Printf("%s,%s\n", f, xls.GetSheetName(i))
		}
	}
}
