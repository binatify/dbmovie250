package spider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	Url string

	Movies []*Movie
}

func NewPage(page int) *Page {
	return &Page{
		Url: fmt.Sprintf("https://movie.douban.com/top250?start=%d", (page-1)*25),
	}
}

var (
	yearRex = regexp.MustCompile("\\d+")
)

func (this *Page) Perform() (err error) {
	doc, err := goquery.NewDocument(this.Url)
	if err != nil {
		return
	}

	movies := make([]*Movie, 0)

	doc.Find(".grid_view li").Each(func(_ int, item *goquery.Selection) {
		var (
			name, director, starring, discription string
			poster, link, area, category          string
			year, star                            string
		)

		pic := item.Find(".pic img")
		name, _ = pic.Attr("alt")
		poster, _ = pic.Attr("src")

		infoText := item.Find(".bd p").First().Text()
		year = yearRex.FindString(infoText)

		infoSplits := yearRex.Split(infoText, -1)

		infoHead := strings.Split(infoSplits[0], "主演:")
		director = strings.Replace(infoHead[0], "导演:", "", 1)
		if len(infoHead) == 2 {
			starring = infoHead[1]
		}

		infoTail := strings.Split(infoSplits[1], "/")
		area = infoTail[1]
		if len(infoTail) > 2 {
			category = infoTail[2]
		}

		star = item.Find(".star .rating_num").Text()

		discription = item.Find(".quote .inq").Text()
		link, _ = item.Find(".pic a").Attr("href")

		movies = append(movies, &Movie{
			Name:        name,
			Director:    director,
			Starring:    starring,
			Year:        year,
			Area:        area,
			Category:    category,
			Star:        star,
			Discription: discription,
			Poster:      poster,
			Link:        link,
		})
	})

	this.Movies = movies
	return nil
}
