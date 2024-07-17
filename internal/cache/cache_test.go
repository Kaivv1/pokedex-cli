package cache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	if cache.Cache == nil {
		t.Error("cache is nil")
	}
}
func TestGetAddCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)

	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://test.com",
			value: []byte("this is a test"),
		},
		{
			key:   "",
			value: []byte("welll duuuhh"),
		},
		{
			key:   "https://test.com/asd",
			value: []byte("this is a asd test"),
		},
		{
			key:   "key1",
			value: []byte("value1, value2"),
		},
	}

	for _, c := range cases {
		cache.Add(c.key, c.value)
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Error("key doesnt exist")
		}
		if string(actual) != string(c.value) {
			t.Errorf("%s doesn't match %s", string(actual), string(c.value))
		}
	}
}

func TestReapSuccess(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "key1"
	cache.Add(key, []byte("value1"))
	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get(key)
	if ok {
		t.Error("value should of been deleted")
	}
}
func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "key1"
	cache.Add(key, []byte("value1"))
	time.Sleep(interval / 2)
	_, ok := cache.Get(key)
	if !ok {
		t.Error("value is not deleted")
	}
}
