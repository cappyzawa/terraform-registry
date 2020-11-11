package provider_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cappyzawa/terraform-registry/config"
	"github.com/cappyzawa/terraform-registry/handler/provider"
	"github.com/go-chi/chi"
	"go.mercari.io/go-httpdoc"
)

func TestVersionHandlerServeHTTP(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		namespace    string
		_type        string
		expectStatus int
	}{
		{
			name:         "existing provider: cappyzawa/concourse",
			namespace:    "cappyzawa",
			_type:        "concourse",
			expectStatus: http.StatusOK,
		},
		{
			name:         "non existing provider: foo/bar",
			namespace:    "foo",
			_type:        "bar",
			expectStatus: http.StatusNotFound,
		},
	}

	document := &httpdoc.Document{
		Name: "Provider versions",
	}
	defer func() {
		if err := document.Generate("../../docs/provider/versions.md"); err != nil {
			t.Fatalf("err: %v", err)
		}
	}()

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			ts := testServer(document, test.name)
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/providers/%s/%s/versions", ts.URL, test.namespace, test._type), nil)
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

func testServer(doc *httpdoc.Document, description string) *httptest.Server {
	cfg, _ := config.Parse("../../testdata/config.yaml")
	pvh := &provider.VersionsHandler{
		Providers: cfg.Providers,
	}

	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return httpdoc.Record(next, doc, &httpdoc.RecordOption{
			Description:    description,
			ExcludeHeaders: []string{"Content-Length", "User-Agent", "Accept-Encoding"},
		})
	})
	r.Get("/v1/providers/{namespace}/{type}/versions", pvh.ServeHTTP)

	return httptest.NewServer(r)
}
