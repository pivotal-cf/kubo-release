package core_test

import (
	"os"
	"osl-collector/core"
	"osl-collector/core/oslStructs"
	"path"
	"reflect"
	"strings"
	"testing"
)

func TestReadFiles_EmptyArray(t *testing.T) {
	output := core.ReadFiles([]string{})

	if len(output) != 0 {
		t.Errorf("Should have not found any json files, but found %d\n", len(output))
	}
}

func TestReadFiles_HappyCase(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	output := core.ReadFiles([]string{
		path.Join(dir, "testNested/testDir1/temp.json"),
		path.Join(dir, "testNested/testDir2/osl-package.json"),
	})

	if len(output) != 2 {
		t.Errorf("Should have found 2 json files but found %d\n", len(output))
	}
	firstJson := strings.TrimSpace(string(output[0]))
	if firstJson != "{}" {
		t.Errorf("Should have found {} in the first file but found %s\n", firstJson)
	}
	secondJson := strings.TrimSpace(string(output[1]))
	if secondJson != "[]" {
		t.Errorf("Should have found [] in the second file but found %s\n", secondJson)
	}
}

func TestParseOSLData(t *testing.T) {
	data := core.ParseOSLData([][]byte{
		[]byte(`
{ "packages": [
  {
    "name": "Name",
    "version": "1.0.0",
    "url": "https://foo.bar",
    "license": "BSD3.0"
  }
]}
`),
	})

	if len(data) != 1 {
		t.Errorf("Should have parsed 1 data but found %d\n", len(data))
	}
	expected := oslStructs.OSLData{Packages: []oslStructs.OSLPackage{
		{"Name", "1.0.0", "https://foo.bar", "BSD3.0"}},
	}
	if !reflect.DeepEqual(expected, data[0]) {
		t.Errorf("Expected %+v but got %+v", expected, data[0])
	}
}
