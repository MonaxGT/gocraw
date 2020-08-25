package gocraw

import (
	"errors"
	"fmt"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html"
)

type Config struct {
	conn *html.Node
}

const linksPath = "//item/link/text()"

// GetOneAttr return data from xpath and filter with attribute
func (c *Config) GetOneAttr(xpath string, attr string) (string, error) {
	a := htmlquery.FindOne(c.conn, xpath)
	return htmlquery.SelectAttr(a, attr), nil
}

// GetAll return data from xpath and concatenate data
func (c *Config) GetAll(xpath string) (string, error) {
	list, err := htmlquery.QueryAll(c.conn, xpath)
	if err != nil {
		return "", err
	}
	var str string
	for _, v := range list {
		str = fmt.Sprintf("%s,", v.Data)
	}
	return str, nil
}

func getAllLinks(rssURL string) ([]string, error) {
	var pages []string
	doc, err := xmlquery.LoadURL(rssURL)
	if err != nil {
		return nil, err
	}
	list := xmlquery.Find(doc, linksPath)
	for _, n := range list {
		pages = append(pages, n.Data)
	}
	return pages, nil
}

//GetRecLinks return all link from feed list
func GetRecLinks(rssURL string) ([]string, error) {
	pages, err := getAllLinks(rssURL)
	if err != nil {
		return nil, err
	}
	if len(pages) == 0 {
		return nil, errors.New("can't find any links")
	}
	return pages, nil
}

// LoadURL connect to URL and return constructor
func LoadURL(URL string) (*Config, error) {
	conn, err := htmlquery.LoadURL(URL)
	if err != nil {
		return nil, err
	}
	return &Config{
		conn: conn,
	}, nil
}
