package page

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	return ioutil.WriteFile("data/"+p.Title+".txt", p.Body, 0600)
}

func Find(title string) (*Page, error) {
	body, err := ioutil.ReadFile("data/" + title + ".txt")
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
