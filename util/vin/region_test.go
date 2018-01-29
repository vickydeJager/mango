package vin

import "testing"

func TestGetRegion_Valid_SouthAfrica(t *testing.T) {
	input := "AD"
	actual := getRegion(input)
	expectedCont := "Africa"
	expectedCountry := "South Africa"

	if actual.Continent != expectedCont {
		t.Errorf("Continent Error: Expected %s got %s", expectedCont, actual.Continent)

		if actual.Country != expectedCountry {
			t.Errorf("Region Error: Expected %s got %s", expectedCountry, actual.Country)
		}
	}
}

func TestGetRegion_IllegalChars(t *testing.T) {
	input := "QI"
	actual := getRegion(input)

	if actual.Continent != "" || actual.Country != "" {
		t.Errorf("Expected empty values, got %s %s", actual.Continent, actual.Country)
	}
}