package csljson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// func TestParseFile(t *testing.T) {
// 	tests := []struct {
// 		path           string
// 		wantErr        bool
// 		title          string //tests info.title
// 		citationPrefix string
// 	}{
// 		{"styles/apa-6th-edition.csl", false, "American Psychological Association 6th edition", "("},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.path, func(t *testing.T) {
// 			s, err := ParseFile(tt.path)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if s.Info.Title != tt.title {
// 				t.Errorf("ParseFile() wanted = %v, got %v", tt.title, s.Info.Title)
// 			}
// 			if s.Citation.Layout.Prefix != tt.citationPrefix {
// 				t.Errorf("ParseFile() wanted = %v, got %v", tt.citationPrefix, s.Citation.Layout.Prefix)
// 			}
// 		})
// 	}
// }

//ExampleParseFile tests parser against sample csljson files

func ExampleParseFile2() {
	s, err := ParseFile("test/previews.json")
	if err != nil {
		fmt.Printf("ParseFile() error = %v", err)
		return
	}
	fmt.Println(s.Items[0].DOI)
	fmt.Println(s.Items[0].Author[0].Given)
	fmt.Println(s.Items[0].Issued.DateParts[0][0])

	// output:
	// 10.1016/j.cub.2016.05.047
	// Rumi
	// 2016
}

// Variant json testing
type CustomerAddress struct {
	Variable Variant `json:"variable"`
}

var js = []byte(`[
  {
    "variable": 55555
  },
  {
    "variable": "55555"
	},
	{
    "variable": true
	},
	{
    "variable": null
	},
	{
		"variable": ""		
  }
]`)

func TestVariantJson(t *testing.T) {
	var addresses []CustomerAddress
	err := json.Unmarshal(js, &addresses)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("addr : %#v\n", addresses)
}
