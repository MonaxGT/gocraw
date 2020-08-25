package main

import (
	"fmt"
	"log"

	"github.com/MonaxGT/gocraw"
)

func main() {
	pages, err := gocraw.GetRecLinks("https://xakep.ru/feed/")
	if err != nil {
		log.Fatal(err)
	}

	idXpath := "//div/article"
	titleXpath := "//title/text()"

	for _, v := range pages {
		conf, err := gocraw.LoadURL(v)
		if err != nil {
			log.Fatal(err)
		}

		res, err := conf.GetOneAttr(idXpath, "data-id")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)

		res, err = conf.GetAll(titleXpath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}

}
