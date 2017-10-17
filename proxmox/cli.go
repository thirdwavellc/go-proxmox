package proxmox

import (
	"log"
	"os"
	"reflect"
)

// PrintError prints the error in a standardized format, and exits with
// return status 1.
func PrintError(err error) {
	log.Println("There was an error...")
	log.Printf("Error: %v", err)
	os.Exit(1)
}

// PrintDataSlice uses an empty interface/reflection to print any slice
// of any struct type. This allows the data returned from the Proxmox API
// to be printed in a standardized format.
func PrintDataSlice(data interface{}) {
	d := reflect.ValueOf(data)

	for i := 0; i < d.Len(); i++ {
		dataItem := d.Index(i)
		typeOfT := dataItem.Type()

		for j := 0; j < dataItem.NumField(); j++ {
			f := dataItem.Field(j)
			log.Printf("%s: %v\n", typeOfT.Field(j).Name, f.Interface())
		}
		log.Printf("\n")
	}
}

// PrintDataStruct uses an empty interface/reflection to print any single
// struct in a standardized format.
func PrintDataStruct(data interface{}) {
	d := reflect.ValueOf(data)
	typeOfT := d.Type()

	for j := 0; j < d.NumField(); j++ {
		f := d.Field(j)
		log.Printf("%s: %v\n", typeOfT.Field(j).Name, f.Interface())
	}
	log.Printf("\n")
}
