package config_test

import (
	"testing"

	"github.com/cappyzawa/terraform-registry/config"
)

func TestParse(t *testing.T) {
	t.Parallel()
	actual, err := config.Parse("../testdata/config.yaml")
	if err != nil {
		t.Errorf("err should not occur: %v", err)
	}
	if len(actual.Providers) != 1 {
		t.Errorf("length of providers should be 1, but it is %v", len(actual.Providers))
	}
	if actual.Providers[0].Namespace != "cappyzawa" {
		t.Errorf("provider namespace should be cappyzawa, but it is %v", actual.Providers[0].Namespace)
	}
}
