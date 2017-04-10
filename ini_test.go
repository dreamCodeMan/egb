package egb

import "testing"

var iniconfpath string = "./testdata/conf.ini"

func TestGetValue(t *testing.T) {
	conf := SetConfig(iniconfpath)
	name := conf.GetValue("database", "username")
	correct := "root"
	if name != correct {
		t.Errorf("GetValue\n Expect => %s\n Got => %s\n", correct, name)
	}
}

func TestSetValue(t *testing.T) {
	conf := SetConfig(iniconfpath)
	suc := conf.SetValue("database", "xxx", "xxx")
	if !suc {
		t.Errorf("SetValue Error")
	}
	name := conf.GetValue("database", "xxx")
	if name != "xxx" {
		t.Errorf("GetValue\n Expect => %s\n Got => %s\n", "xxx", name)
	}
}

func TestDeleteValue(t *testing.T) {
	conf := SetConfig(iniconfpath)
	suc := conf.SetValue("database", "xxx", "xxx")
	if !suc {
		t.Errorf("SetValue Error")
	}
	name := conf.GetValue("database", "xxx")
	if name != "xxx" {
		t.Errorf("GetValue\n Expect => %s\n Got => %s\n", "xxx", name)
	}
	suc = conf.DeleteValue("database", "xxx")
	if !suc {
		t.Errorf("DeleteValue Error")
	}
	name = conf.GetValue("database", "xxx")
	if name != "" {
		t.Errorf("GetValue\n Expect => %s\n Got => %s\n", "", name)
	}
}
