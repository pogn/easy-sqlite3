package dbsearch

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DBPATH = "./sqlite3_db.db"
const CREATE_TABLE_SQL = "CREATE TABLE AA (name text);"
// var DATA_STRUCT = ""

func InitDB() *sql.DB { // DB생성 (db포인터, file 위치)
	os.Remove(DBPATH)
	
	file, err := os.Create(DBPATH) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite3 db created")
   
	db, err := sql.Open("sqlite3", DBPATH)
	if err != nil {
		log.Println("Could not init db - error occured")
		panic(err)
	}
	if db == nil {
		log.Println("Could not init db - db nil")
		panic(err)
	}
	return db
}

func CreateTable(db *sql.DB) (bool) { // Table 생성 ()
	_,err := db.Exec(CREATE_TABLE_SQL)
	if err != nil {
		log.Println("Fail - create table")
		panic(err)
	}
	return true
}



/*
func InsertEntry(db *sql.DB, []struct) (bool) {  // input : Struct List 
	// struct내의 데이터 type을 main에서 전달받거나, 추론해야한다. 
	// 데이터타입 추론 관련 참고 링크
	// . https://www.geeksforgeeks.org/how-to-find-the-type-of-struct-in-golang/
	// . https://intrepidgeeks.com/tutorial/go-attribute-process-part-4-type-transformation-and-reasoning
}

func SetDataStruct () { // struct data type 
	var DATA_STRUCT =""
}

func ImportCSVtoTable { // CSV import 

}

func ExportCSVtoTable { // CSV export

}

func _createSQL() {

}

func _insertSQL() {

}

func _


*/