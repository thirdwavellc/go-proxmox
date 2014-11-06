package main

import (
	"fmt"
	"os"
	"reflect"
)

func PrintError(err error) {
	fmt.Println("There was an error...")
	fmt.Printf("Error: %v", err)
	os.Exit(1)
}

func PrintDataSlice(data interface{}) {
	d := reflect.ValueOf(data)

	for i := 0; i < d.Len(); i++ {
		dataItem := d.Index(i)
		typeOfT := dataItem.Type()

		for j := 0; j < dataItem.NumField(); j++ {
			f := dataItem.Field(j)
			fmt.Printf("%s: %v\n", typeOfT.Field(j).Name, f.Interface())
		}
		fmt.Printf("\n")
	}
}

func PrintDataStruct(data interface{}) {
	d := reflect.ValueOf(data)
	typeOfT := d.Type()

	for j := 0; j < d.NumField(); j++ {
		f := d.Field(j)
		fmt.Printf("%s: %v\n", typeOfT.Field(j).Name, f.Interface())
	}
	fmt.Printf("\n")
}
