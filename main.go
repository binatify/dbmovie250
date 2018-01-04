package main

import (
	"flag"

	"github.com/binatify/dbmovie250/spider"
	"github.com/binatify/dbmovie250/storage"
)

var (
	csvPath string
)

func main() {
	flag.StringVar(&csvPath, "csv", "./data.csv", "the data saved csv name.")
	flag.Parse()

	stg := storage.NewStgStorage(csvPath)
	stg.Save(spider.Run())
}
