package main_test

import "testing"

func testhelloWorld(t *testing.T) {
	if Helloworld() != "Hello World!!" {
		t.Fatal("Test fail")
	} else {
		t.Logf("OK")
	}
}
