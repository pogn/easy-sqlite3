package main


import (
	"fmt"
	ds "src.com/dbsearch"
	
)

func main() {
	db_1 := ds.InitDB()
	defer db_1.Close()

	ds.CreateTable(db_1)
	//ds.InsertEntry(db_1, )

	fmt.Println("done")

}

/*
func makeData(){
	sensor ss1 = {"junginpc", "1.1.1.1", "", }
}
*/