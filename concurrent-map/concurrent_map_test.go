package concurrent_map

import (
	"testing"
	"strconv"
)

func TestNew(t *testing.T) {
	m := New()
	if m == nil {
		t.Error("map is null.")
	}
	if m.Count() != 0 {
		t.Error("new map should be empty.")
	}
}

func TestConcurrentMap_Set(t *testing.T) {
	m := New()
	m.Set("key1", "value1")
	m.Set("key2", "value2")
	if m.Count() != 2 {
		t.Error("map set error")
	}
}

func TestConcurrentMap_MultiSet(t *testing.T) {
	m := New()
	data := make(map[string]interface{})
	data["a"] = "a"
	data["b"] = "b"
	m.MultiSet(data)
	if m.Count() != 2 {
		t.Error("map multiset error")
	}
}

func TestConcurrentMap_SetIfAbsent(t *testing.T) {
	m := New()
	m.Set("key", "value")
	m.Set("key1", "value1")
	m.SetIfAbsent("key", "newvalue")
	if m.Count() != 2 {
		t.Error("map SetIfAbsent error")
	}
}

func TestConcurrentMap_Get(t *testing.T) {
	m := New()
	m.Set("a", "a")
	val, ok := m.Get("a")
	if !ok {
		t.Error("map get error")
	}
	if val.(string) != "a" {
		t.Error("map get error")
	}
}

func TestConcurrentMap_Count(t *testing.T) {
	m := New()
	m.Set("a", "a")
	if m.Count() != 1 {
		t.Error("map count error")
	}
}

func TestConcurrentMap_Delete(t *testing.T) {
	m := New()
	m.Set("a", "a")
	m.Delete("a")
	_, ok := m.Get("a")
	if ok {
		t.Error("map delete error")
	}
}

func TestConcurrentMap_Pop(t *testing.T) {
	m := New()
	m.Set("a", "a")
	val, ok := m.Pop("a")
	if !ok {
		t.Error("map pop error")
	}
	if val.(string) != "a" {
		t.Error("map pop error")
	}
	_, ok = m.Get("a")
	if ok {
		t.Error("map pop error")
	}
}

func TestConcurrentMap_Items(t *testing.T) {
	m := New()
	for i := 0; i < 100; i++ {
		m.Set(strconv.Itoa(i), i)
	}
	if m.Count() != 100 {
		t.Error("map items error")
	}
}
