# dbmovie250
A spider to collect douban top250  movies. 
![screen shot 2018-01-04 at 2 20 58 pm](https://user-images.githubusercontent.com/1459834/34552164-d5ec5bf2-f15a-11e7-8591-83a75b41385d.png)

### Usage

1. Use `go get github.com/binatify/dbmovie250`
2. Or [binary](https://github.com/binatify/dbmovie250/releases/tag/v0.1) release.

  ```bash
  $ dbmovie250 -h

  Usage of dbmovie250:
    -csv string
        the data saved csv name. (default "./data.csv")
  ```

when you run `dbmovie250 -csv xxx`, you will get data in file `xxx.csv`.

### Data struct

``` golang
type Movie struct {
	Name     string
	Director string
	Starring string

	Year     string
	Area     string
	Category string
	Star     string

	Discription string
	Poster      string
	Link        string
}

```

