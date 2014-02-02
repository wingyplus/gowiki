package page_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/gowiki/page"
	"io/ioutil"
	"os"
	"testing"
)

func TestPage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Page Spec")
}

var _ = Describe("Page", func() {
	var p *page.Page

	BeforeEach(func() {
		os.Mkdir("data", os.ModePerm)
		p = &page.Page{Title: "Hello", Body: []byte("This is a book")}
	})
	AfterEach(func() {
		os.RemoveAll("data")
	})
	Describe("Save", func() {
		AfterEach(func() {
			os.Remove("data/Hello.txt")
		})
		It("should be nil when save success", func() {
			Expect(p.Save()).To(BeNil())
		})
		It("should save file to data/ and filename is {{.Title}}.txt", func() {
			p.Save()
			file, err := ioutil.ReadFile("data/Hello.txt")
			Expect(err).To(BeNil())
			Expect(file).To(Equal([]byte("This is a book")))
		})
	})
	Describe("Find", func() {
		Describe("Found", func() {
			BeforeEach(func() {
				p.Save()
			})
			AfterEach(func() {
				os.Remove("data/Hello.txt")
			})
			It("should be return Page object by given Title", func() {
				helloPage, err := page.Find("Hello")

				Expect(err).To(BeNil())
				Expect(helloPage.Title).To(Equal("Hello"))
				Expect(helloPage.Body).To(Equal([]byte("This is a book")))
			})
		})
		Describe("Not found", func() {
			It("should be return error", func() {
				helloPage, err := page.Find("Hello")

				Expect(helloPage).To(BeNil())
				Expect(err.Error()).To(Equal("open data/Hello.txt: no such file or directory"))
			})
		})
	})
})
