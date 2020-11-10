package config_test

import (
	"testing"

	"github.com/cappyzawa/terraform-registry/config"
)

func TestParse(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input     string
		existErr  bool
		expectLen int
	}{
		"config file exists": {
			input:     "../testdata/config.yaml",
			existErr:  false,
			expectLen: 1,
		},
		"config file does not exist": {
			input:     "nonexist.yaml",
			existErr:  true,
			expectLen: 0,
		},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := config.Parse(test.input)
			if err != nil && !test.existErr {
				t.Errorf("error should not occur: %v", err)
			}
			if err == nil && test.existErr {
				t.Error("error should occur")
			}
			if actual != nil && len(actual.Providers) != test.expectLen {
				t.Errorf("length of provider should be %v, but it is %v", test.expectLen, len(actual.Providers))
			}
		})
	}
}
