package vin

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/louisevanderlith/mango/util/vin/common"
)

func TestLoadSectionCodes_Correct(t *testing.T) {
	input := "JHLRD68405C200888"
	expected := common.VINSections{
		WMICode:           "JHL",
		FeatureCode:       "RD684",
		CheckDigit:        "0",
		YearCode:          "5",
		AssemblyPlantCode: "C",
		SequenceCode:      "200888",
		FullVIN:           input,
	}

	actual, _ := common.LoadVINSections(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expecting %v, got %v", expected, actual)
	}
}

func TestGetInfo_Correct(t *testing.T) {
	input := "JHLRD68405C200888"
	actual, _ := GetInfo(input)

	fmt.Println("ACT", actual)

	t.Fail()
}
