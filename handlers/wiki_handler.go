package handlers

import (
	"fmt"
	"github.com/wingyplus/gowiki/page"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	p, err := page.Find(title)
	if err != nil {
		http.Redirect(w, req, "/edit/"+title, http.StatusFound)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
