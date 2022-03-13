package util

import (
	"github.com/qkgo/scaff/pkg/util/filesystem"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestRecursionListFile(t *testing.T) {
	testDirName := "TestDir"
	testDir := t.TempDir()
	testDeepDir := path.Join(testDir, testDirName, testDirName, testDirName, testDirName)
	os.MkdirAll(testDeepDir, os.ModePerm)
	results1 := filesystem.ExistsDir(testDeepDir)
	t.Logf("%+v", results1)
	ioutil.WriteFile(path.Join(testDeepDir, "TestFile"), []byte("TestContent"), os.ModePerm)
	ioutil.WriteFile(path.Join(testDeepDir, "TestFile2"), []byte("TestContent2"), os.ModePerm)
	results, _ := RecursionListFile(testDir, nil, nil)
	t.Logf("%+v", results)
	assert.Equal(t, len(results), 6)
}
