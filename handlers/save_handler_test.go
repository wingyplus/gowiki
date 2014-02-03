package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/gowiki/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
)

var _ = Describe("Save Handler", func() {
	var handler func(w http.ResponseWriter, r *http.Request)
	BeforeEach(func() {
		handler = handlers.SaveHandler
		os.Mkdir("data", os.ModePerm)
	})
	AfterEach(func() {
		os.RemoveAll("data")
	})
	Context("Save successful", func() {
		var req *http.Request
		var w *httptest.ResponseRecorder

		BeforeEach(func() {
			page := url.Values{}
			page.Set("body", "Hello World")
			req, _ = http.NewRequest("POST", "/save/PageFound", strings.NewReader(page.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w = httptest.NewRecorder()
		})
		AfterEach(func() {
			os.Remove("data/PageFound.txt")
		})
		It("should redirect to /view/PageFound when save file", func() {
			handler(w, req)

			Expect(w.Code).To(Equal(http.StatusFound))
			Expect(w.Header().Get("Location")).To(Equal("/view/PageFound"))
		})
		It(`should have file "PageFound.txt" when save page`, func() {
			handler(w, req)
			content, err := ioutil.ReadFile("data/PageFound.txt")

			Expect(err).To(BeNil())
			Expect(content).To(Equal([]byte("Hello World")))
		})
	})
})
