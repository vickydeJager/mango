package vin

import (
	"reflect"
	"testing"
)

func TestGetIllegalMap_Success(t *testing.T) {
	input := []string{"I", "O", "Q"}
	expected := []int{9, 15, 17}
	actual := getIllegalMap(input...)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func TestGetWeight_Success(t *testing.T) {
	input := "K"
	expected := 10
	actual := getWeight(input, "I")

	if expected != actual {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
