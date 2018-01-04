package storage

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/binatify/dbmovie250/spider"
)

type CsvStorage struct {
	Path string
}

func NewStgStorage(file string) *CsvStorage {
	return &CsvStorage{
		Path: file,
	}
}

func (stg *CsvStorage) Save(movies []*spider.Movie) (err error) {
	file, err := os.Create(stg.Path)
	if err != nil {
		log.Fatalln("Cannot create file", err)
	}

	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"电影名称", "导演", "演员", "年份", "地区", "类型", "评分", "简介", "海报封面", "豆瓣详情链接"}
	w.Write(header)

	for _, movie := range movies {
		row := []string{
			movie.Name,
			movie.Director,
			movie.Starring,
			movie.Year,
			movie.Area,
			movie.Category,
			movie.Star,
			movie.Discription,
			movie.Poster,
			movie.Link,
		}

		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	return
}
