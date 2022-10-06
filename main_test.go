package main

import "testing"

func TestAddition(t *testing.T) {
	expectedResult := 900

	actualResult := Addition(450, 400, 50)

	if actualResult != expectedResult {
		t.Errorf("Addition test failed. Expecyed %d, Received %d", expectedResult, actualResult)
	}

}

func TestSubtraction(t *testing.T) {
	expectedResult := 90
	actualResult := Subtraction(100, 10)

	if actualResult != expectedResult {
		t.Errorf("Subtraction test failed. Expecyed %d, Received %d", expectedResult, actualResult)
	}
}

func TestMultiplication(t *testing.T) {
	expectedResult := 81
	actualResult := Multiplication(9, 9)

	if actualResult != expectedResult {
		t.Errorf("Subtraction test failed. Expected %d, Received %d", expectedResult, actualResult)
	}
}

func TestDivision(t *testing.T) {
	expectedResult := 2
	actualResult := Division(10, 5)

	if actualResult != expectedResult {
		t.Errorf("Division test failed. Expected %d, actual %d", expectedResult, actualResult)
	}

}
