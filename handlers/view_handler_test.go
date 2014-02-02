package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/gowiki/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("ViewHandler", func() {
	var handler func(w http.ResponseWriter, r *http.Request)

	BeforeEach(func() {
		handler = handlers.ViewHandler
	})
	Describe("Found page", func() {
		BeforeEach(func() {
			os.Mkdir("data", os.ModePerm)
			ioutil.WriteFile("data/FoundPage.txt", []byte("HelloWorld"), 0600)
		})
		AfterEach(func() {
			os.RemoveAll("data")
		})
		It("should return expect content", func() {
			req, _ := http.NewRequest("GET", "/view/FoundPage", nil)
			w := httptest.NewRecorder()

			handler(w, req)

			Expect(w.Body.String()).To(Equal("<h1>FoundPage</h1><div>HelloWorld</div>"))
		})
	})
	Describe("Not found page", func() {
		It("should redirect to /edit/NotFoundPage", func() {
			req, _ := http.NewRequest("GET", "/view/NotFoundPage", nil)
			w := httptest.NewRecorder()

			handler(w, req)

			Expect(w.Code).To(Equal(http.StatusFound))
			Expect(w.Header()["Location"]).To(Equal([]string{"/edit/NotFoundPage"}))
		})
	})
})
