package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {

	header1 := http.Header{}
	header1.Add("", "")
	header2 := http.Header{}
	header2.Add("Authentication", "asadsasd")
	header3 := http.Header{}
	header3.Add("Authorization", "apikey thisisthekey")
	header4 := http.Header{}
	header4.Add("Authorization", "ApiKey thisisthekey")

	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"nodata":                           {input: header1, want: ""},
		"wrong header":                     {input: header2, want: ""},
		"correct header, wrong key format": {input: header3, want: ""},
		"correct header, correct key":      {input: header4, want: "thisisthekey"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err != nil {
				t.Log(err)
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
