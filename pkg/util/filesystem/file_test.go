package filesystem

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestExists(t *testing.T) {
	scanPath := os.TempDir()
	writeData := "expectDatas"
	inlineDir := path.Join(scanPath, writeData)
	writeData1 := "expectData"
	inlineFile1 := path.Join(inlineDir, writeData1)
	err := os.WriteFile(inlineFile1, []byte(writeData), os.ModePerm)
	if err != nil {
		t.Logf("precondition initial failed : %v", err)
		return
	}
	assert.Equal(t, true, Exists(inlineFile1))
	assert.Equal(t, false, Exists(inlineFile1+"1"))

}

func TestCopyIOFile(t *testing.T) {
	scanPath := os.TempDir()
	writeData := "expectDatas"
	inlineDir := path.Join(scanPath, writeData)
	writeData1 := "expectData"
	inlineFile1 := path.Join(inlineDir, writeData1)
	inlineFile2 := path.Join(inlineDir, "expectData2")
	err := os.WriteFile(inlineFile1, []byte(writeData), os.ModePerm)
	if err != nil {
		t.Logf("precondition initial failed : %v", err)
		return
	}
	err = CopyIoFile(inlineFile1, inlineFile2)
	if err != nil {
		t.Logf("precondition initial failed : %v", err)
		return
	}

	assert.Equal(t, true, Exists(inlineFile2))
	assert.Equal(t, false, Exists(inlineFile2+"1"))

}
