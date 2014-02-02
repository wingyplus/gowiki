package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wingyplus/gowiki/handlers"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("EditHandler", func() {
	var handler func(w http.ResponseWriter, r *http.Request)

	BeforeEach(func() {
		handler = handlers.EditHandler
	})
	Context("Not found page", func() {
		It("show Title and empty text in textarea", func() {
			req, _ := http.NewRequest("GET", "/edit/PageNotFound", nil)
			w := httptest.NewRecorder()

			handler(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal(`<h1>PageNotFound</h1>
<form action="/save/PageNotFound" method="POST">
	<textarea name="body" rows="20" cols="80"></textarea>
	<button type="submit">Save</button>
</form>`))
		})
	})
})
