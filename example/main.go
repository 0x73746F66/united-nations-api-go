package main

import (
	"fmt"
	"unapi"
)

func main() {
	comp, _ := unapi.FetchOranizations()
	fmt.Printf("%+v\n", comp)

	for i := range comp {
		//db, _ := unapi.FetchDatabases(comp[i])
		//fmt.Printf("%+v\n", db)
		//datasets, _ := unapi.FetchDatasets(comp[i])
		unapi.FetchDatasets(comp[i])
		//fmt.Printf("%+v\n", datasets)
	}
}
