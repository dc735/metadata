package metadata

import (
	//	"fmt"
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"
)

var testLocation Location

func init() {

	testLocation = Location{
		Id:       "location",
		Name:     "A Location Name",
		Latitude: &[]float32{-41.5}[0],
		Services: []string{"Test 1 Service", "Test 2 Service"},
		Tags:     []string{"ABC", "DEF"},
		Links: []Link{
			Link{
				Id:       "somewhere",
				Key:      &[]string{"Key 1"}[0],
				Role:     &[]string{"Role 1"}[0],
				Polarity: &[]string{"Polarity 1"}[0],
			},
			Link{
				Id:       "else",
				Key:      &[]string{"Key 2"}[0],
				Role:     &[]string{"Role 2"}[0],
				Polarity: &[]string{"Polarity 2"}[0],
			},
		},
		Access: &[]string{"Some Access Info\nSome More Access Info"}[0],
		Notes:  &[]string{"Some Notes\nSome More Notes"}[0],
	}
}

func TestLocation_DecodeFile(t *testing.T) {
	t.Log("Check decoding location file.")
	{
		var l Location
		if _, err := toml.DecodeFile("testdata/location.toml", &l); err != nil {
			t.Fatal(err)
		}
		if l.String() != testLocation.String() {
			t.Errorf("location file text mismatch: [\n%s\n]", SimpleDiff(l.String(), testLocation.String()))
		}
	}
}

func TestLocation_ReadFile(t *testing.T) {
	t.Log("Compare loaded location file.")
	{
		b, err := ioutil.ReadFile("testdata/location.toml")
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != testLocation.String() {
			t.Errorf("location file text mismatch: [\n%s\n]", SimpleDiff(string(b), testLocation.String()))
		}
	}
}

func TestLocation_LoadFile(t *testing.T) {
	t.Log("Check loading location file.")
	{
		l, err := LoadLocation("testdata/location.toml")
		if err != nil {
			t.Fatal(err)
		}
		if l == nil {
			t.Fatalf("location file load problem")
		}
		if l.String() != testLocation.String() {
			t.Errorf("location file decode mismatch: [\n%s\n]", SimpleDiff(l.String(), testLocation.String()))
		}
	}
}

func TestLocation_LoadFiles(t *testing.T) {
	t.Log("Check loading location files.")
	{
		l, err := LoadLocations("testdata", "location.toml")
		if err != nil {
			t.Fatal(err)
		}
		if len(l) != 1 {
			t.Fatalf("location files load problem")
		}
		if l[0].String() != testLocation.String() {
			t.Errorf("location file decode mismatch: [\n%s\n]", SimpleDiff(l[0].String(), testLocation.String()))
		}
	}
}
