package spider

import (
	"log"
	"sync"
)

func Run() (movies []*Movie) {
	var (
		wg        sync.WaitGroup
		totalPage = 10
		pages     = make([]*Page, totalPage)
	)

	for pIndex := 1; pIndex <= totalPage; pIndex++ {
		wg.Add(1)

		page := NewPage(pIndex)
		pages[pIndex-1] = page

		go func() {
			defer wg.Done()

			err := page.Perform()
			if err != nil {
				log.Panic(err)
			}
		}()
	}

	wg.Wait()

	for _, page := range pages {
		movies = append(movies, page.Movies...)
	}

	return
}
