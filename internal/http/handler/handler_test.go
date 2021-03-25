package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cappyzawa/terraform-registry/internal/http/handler"
	"github.com/go-chi/chi"
	"go.mercari.io/go-httpdoc"
)

func TestWellKnownHandlerServeHTTP(t *testing.T) {
	document := &httpdoc.Document{
		Name: "WellKnown",
	}
	defer func() {
		if err := document.Generate("../../../docs/well_known.md"); err != nil {
			t.Fatalf("err: %v", err)
		}
	}()

	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Use(func(next http.Handler) http.Handler {
		return httpdoc.Record(next, document, &httpdoc.RecordOption{
			Description:    "",
			ExcludeHeaders: []string{"Content-Length", "User-Agent", "Accept-Encoding"},
		})
	})
	h := handler.New()
	r.Get("/.well-known/terraform.json", h.WellKnown)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/.well-known/terraform.json", ts.URL), nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("status code should be %v, but it is %v", http.StatusOK, res.StatusCode)
	}
}
