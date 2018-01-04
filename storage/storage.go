package storage

import "github.com/binatify/dbmovie250/spider"

type Storage interface {
	Save([]*spider.Movie) (err error)
}
