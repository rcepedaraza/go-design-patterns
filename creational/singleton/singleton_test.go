package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	//expectedCounter := counter1
}

func Test_singleton_AddOne(t *testing.T) {
}
