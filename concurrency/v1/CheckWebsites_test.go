package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	} else if url == "http://justfortest.com"{
		return true
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
		"http://justfortest.com",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
		"http://justfortest.com":	  true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	// Map values are deeply equal when all of the following are true:
	// they are both nil or both non-nil, they have the same length,
	// and either they are the same map object or their corresponding keys
	// (matched using Go equality) map to deeply equal values.

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
