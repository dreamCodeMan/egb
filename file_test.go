package egb

import "testing"

func Test_FileGetString(t *testing.T) {
	_, err := FileGetString("invalid_file")
	if err == nil {
		t.Fail()
	}
	str, err := FileGetString("./testdata/testfile.txt")
	if err != nil {
		t.Error(err)
	}
	if str != "Hello World!" {
		t.Fail()
	}
}

func Test_FileSetString(t *testing.T) {
	randstr := RandString(6)
	err := FileSetString("./testdata/setfile.txt", randstr)
	if err != nil {
		t.Error(err)
	}
	str, err := FileGetString("./testdata/setfile.txt")
	if err != nil {
		t.Error(err)
	}
	if str != randstr {
		t.Fail()
	}
}

func Test_FileAppendString(t *testing.T) {
	randstr1 := RandString(6)
	randstr2 := RandString(6)
	FileSetString("./testdata/appendfile.txt", randstr1)
	err := FileAppendString("./testdata/appendfile.txt", randstr2)
	str, err := FileGetString("./testdata/appendfile.txt")
	if err != nil {
		t.Error(err)
	}
	if str != randstr1 + randstr2 {
		t.Fail()
	}
}

func Test_FileExists(t *testing.T) {
	if !FileExists("./testdata/testfile.txt") {
		t.Fail()
	}
	if FileExists("invalid_file") {
		t.Fail()
	}
}

func Test_FileIsDir(t *testing.T) {
	if FileIsDir("./testdata/testfile.txt") {
		t.Fail()
	}
	if !FileIsDir(".") {
		t.Fail()
	}
}

func Test_FileFind(t *testing.T) {
	dirs := []string{"."}
	filename := "./testdata/testfile.txt"
	path, found := FileFind(dirs, filename)
	if !found {
		t.Fail()
	}
	if path != "testdata/testfile.txt" {
		t.Fail()
	}
	filename = "invalid_file"
	path, found = FileFind(dirs, filename)
	if found {
		t.Fail()
	}
}

func Test_FileGetPrefix(t *testing.T) {
	filename := "xxx.txt"
	correct := "xxx"
	result := FileGetPrefix(filename)
	if result != correct {
		t.Fail()
	}
}

func Test_FileGetExt(t *testing.T) {
	filename := "xxx.txt"
	correct := "txt"
	result := FileGetExt(filename)
	if result != correct {
		t.Fail()
	}
}

func Test_ListDir(t *testing.T) {
}