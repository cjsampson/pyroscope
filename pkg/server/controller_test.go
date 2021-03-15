package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	"github.com/pyroscope-io/pyroscope/pkg/config"
	"github.com/pyroscope-io/pyroscope/pkg/storage"
)

var _ = Describe("server package", func() {
	Context("Controller", func() {
		It("does what its suppose to", func() {
			cfg := config.NewForTests("storage/path")
			s, err := storage.New(cfg)
			if err != nil {
				log.Fatalf("does what its suppose to err: %v\n", err.Error())
			}

			controller := New(cfg, s)
			mux := controller.CreateMux()

			req, _ := http.NewRequest(http.MethodGet, "/ingest", nil)
			rr := executeRequest(mux, req)

			fmt.Println("---------------")
			fmt.Printf("rr:%#v\n", rr)
			fmt.Println("---------------")
		})
	})
})

func executeRequest(mux *http.ServeMux, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}
