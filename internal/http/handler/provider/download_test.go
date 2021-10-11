package provider_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cappyzawa/terraform-registry/config"
	"github.com/cappyzawa/terraform-registry/internal/http/handler/provider"
	"github.com/go-chi/chi"
	"go.mercari.io/go-httpdoc"
)

func TestDownloadHandlerServeHTTP(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		namespace    string
		_type        string
		version      string
		os           string
		arch         string
		expectStatus int
	}{
		{
			name:         "existing provider: cappyzawa/concourse:0.1.0(darwin/amd64)",
			namespace:    "cappyzawa",
			_type:        "concourse",
			version:      "0.1.0",
			os:           "darwin",
			arch:         "amd64",
			expectStatus: http.StatusOK,
		},
		{
			name:         "non existing provider: foo/bar:0.1.0(darwin/amd64)",
			namespace:    "foo",
			_type:        "bar",
			version:      "0.1.0",
			os:           "darwin",
			arch:         "amd64",
			expectStatus: http.StatusNotFound,
		},
		{
			name:         "non existing provider version: cappyzawa/concourse:11.11.0(darwin/amd64)",
			namespace:    "cappyzawa",
			_type:        "concourse",
			version:      "11.11.0",
			os:           "darwin",
			arch:         "amd64",
			expectStatus: http.StatusNotFound,
		},
		{
			name:         "non existing provider os: cappyzawa/concourse:0.1.0(windows/amd64)",
			namespace:    "cappyzawa",
			_type:        "concourse",
			version:      "0.1.0",
			os:           "windows",
			arch:         "amd64",
			expectStatus: http.StatusNotFound,
		},
	}

	document := &httpdoc.Document{
		Name: "Provider versions",
	}
	defer func() {
		if err := document.Generate("../../../../docs/provider/download.md"); err != nil {
			t.Fatalf("err: %v", err)
		}
	}()

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			ts := testDownloadServer(document, test.name)
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/providers/%s/%s/%s/download/%s/%s", ts.URL, test.namespace, test._type, test.version, test.os, test.arch), nil)
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("err: %v", err)
			}
			defer res.Body.Close()
			if res.StatusCode != test.expectStatus {
				t.Errorf("status code should be %v, but it is %v", test.expectStatus, res.StatusCode)
			}
		})
	}
}

func testDownloadServer(doc *httpdoc.Document, description string) *httptest.Server {
	cfg, _ := config.Parse("../../../../testdata/config.yaml")
	ph := provider.NewHandler(
		provider.Providers(cfg.Providers),
		provider.Logger(log.New(os.Stderr, "", 0)),
	)

	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Use(func(next http.Handler) http.Handler {
		return httpdoc.Record(next, doc, &httpdoc.RecordOption{
			Description:    description,
			ExcludeHeaders: []string{"Content-Length", "User-Agent", "Accept-Encoding"},
		})
	})
	r.Get("/v1/providers/{namespace}/{type}/{version}/download/{os}/{arch}", ph.Download)

	return httptest.NewServer(r)
}
