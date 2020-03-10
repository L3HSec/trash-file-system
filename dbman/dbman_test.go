package dbman

import (
	"fmt"
	"testing"
	"time"

	"github.com/L3HSec/trash-file-system/common"
)

func TestDB(t *testing.T) {
	man := NewManager("tfs.db")
	err := man.AddFile(&common.File{
		ID:       0xdeadbeef,
		FileName: "testfile",
		FileSize: 0x10,
		Comment:  "comment",
		Expire:   time.Now().Add(time.Hour),
	})
	if err != nil {
		t.Fatal("cannot add")
	}

	err = man.AddFile(&common.File{
		ID:       0xcafebabe,
		FileName: "testfile2",
		FileSize: 0x23123,
		Comment:  "comment again",
		Expire:   time.Now().Add(time.Hour),
	})
	if err != nil {
		t.Fatal("cannot add")
	}

	err = man.DeleteFile(0xdeadbeef)
	if err != nil {
		t.Fatal("cannot delete")
	}

	list, err := man.ListFiles()
	if err != nil {
		t.Fatal("cannot list")
	}
	for _, v := range list {
		fmt.Println(v)
	}

	file, err := man.QueryFile(0xcafebabe)
	if err != nil || file.FileName != "testfile2" {
		t.Fatal("cannot query")
	}
}
