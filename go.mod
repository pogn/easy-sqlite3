module src

go 1.18

require github.com/mattn/go-sqlite3 v1.14.12 // indirect
require "src.com/dbsearch" v0.0.0

replace "src.com/dbsearch" v0.0.0 => ./dbsearch
