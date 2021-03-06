package enums

import (
	"strings"
)

type OptimizeType int

const (
	Logo OptimizeType = iota
	Banner
	Ad
	Thumb
)

var optimizeTypes = [...]string{
	"Logo",
	"Banner",
	"Ad",
	"Thumb"}

func (r OptimizeType) String() string {
	return roleTypes[r]
}

func GetOptimizeType(name string) OptimizeType {
	var result OptimizeType

	for k, v := range optimizeTypes {
		if strings.ToUpper(name) == strings.ToUpper(v) {
			result = OptimizeType(k)
			break
		}
	}

	return result
}
