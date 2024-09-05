package main

import "testing"

func TestIsPrime(t *testing.T) {
	testValue := 0
	result, msg := IsPrime(testValue)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false[%s]", testValue, msg)
	}

	testValue = 7
	result, msg = IsPrime(testValue)
	if !result {
		t.Errorf("with %d as test parameter, got true, but expected false[%s]", testValue, msg)
	}
}
