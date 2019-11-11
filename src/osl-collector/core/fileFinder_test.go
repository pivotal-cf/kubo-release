package core_test

import (
	jsonReader "osl-collector/core"
	"testing"
)

func TestFindFiles(t *testing.T) {
	jsonFiles := jsonReader.FindFiles("testNested", jsonReader.DefaultOslPackageFile)

	if len(jsonFiles) != 1 {
		t.Errorf("Should have found 1 test json file but found %d", len(jsonFiles))
	}
	expected := "testNested/testDir2/" + jsonReader.DefaultOslPackageFile
	if jsonFiles[0] != expected {
		t.Errorf("Should have found %s but found %s\n", expected, jsonFiles[0])
	}

}