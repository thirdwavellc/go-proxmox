package proxmox

type MyStruct struct {
	Name string
	ID   int
}

func ExamplePrintDataStruct() {
	myStruct := MyStruct{"Test", 1}
	PrintDataStruct(myStruct)
	// Output:
	// Name: Test
	// ID: 1
}

func ExamplePrintDataSlice() {
	struct1 := MyStruct{"One", 1}
	struct2 := MyStruct{"Two", 2}
	struct3 := MyStruct{"Three", 3}
	dataSlice := []MyStruct{struct1, struct2, struct3}
	PrintDataSlice(dataSlice)
	// Output:
	// Name: One
	// ID: 1
	//
	// Name: Two
	// ID: 2
	//
	// Name: Three
	// ID: 3
	//
}
