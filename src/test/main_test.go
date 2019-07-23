package main

import (
	"testing"
)

func TestStore(t *testing.T) {
	t.Log(t.Name(), "Begin...")
	m := Monster{
		Name:  "红孩儿",
		Age:   119,
		Skill: "喷火",
	}
	res := m.Store()
	if res == true {
		t.Logf("Store() success %t", res)
	} else {
		t.Errorf("Error: Store() fail")
	}

	t.Log(t.Name(), "End...")
}
func TestReStore(t *testing.T) {
	t.Log(t.Name(), "Begin...")
	var m Monster
	m = m.ReStore()
	if m.Name == "红孩儿" {
		t.Logf("Restore() success")
	} else {
		t.Logf("Restore() fail")
	}

	t.Log(t.Name(), "End...")
}
