package descriptors

import "github.com/louisevanderlith/mango/util/vin"

type WMICategory int

type WMIGroup map[string][]WMIDescriptor

func NewWMIGroup(groupName string) *WMIGroup {
	group, ok := groups[groupName]

	if ok {
		return group
	}

	group[groupName] = []WMIDescriptor{}
	return group[groupName]
}

func (g *WMIGroup) Add(manufacturerCode, manufacturerName string, category WMICategory, descriptor vin.Descriptor) {
	descrip := buildDescriptor(manufacturerCode, manufacturerName, category, descriptor)
	g = append(g, descrip)
}

var groups WMIGroup

func init() {
	groups = make(WMIGroup)
	buildGroups()
}

func buildGroups() {
	groupsAlfa()
	groupsAudi()
	groupsBMW()
	groupsFerrari()
	groupsFord()
	groupsGM()
	groupsHonda()
	groupsHyundai()
	groupsIsuzu()
	groupsKIA()
	groupsLamborghini()
	groupsLandRover()
	groupsMaserati()
	groupsMercedes()
	groupsMitsubishi()
	groupsSEAT()
	groupsSkoda()
	groupsSubaru()
	groupsToyota()
	groupsVW()
	groupsVolvo()
}

type WMIDescriptor struct {
	Key          string
	Manufacturer string
	Category     WMICategory
	Descriptor   vin.Descriptor
}

const (
	NotSpecified WMICategory = iota
	PassengerCar
	MPV
	Truck
	Bus
	Motorcycle
	ATV
	IncompleteCar
	BasicChassis
)

func GetDescriptor(countryCode, manufacturerCode string) vin.Descriptor {
	var result vin.Descriptor

	return result
}

func buildDescriptor(manufacturerCode, manufacturerName string, category WMICategory, descriptor vin.Descriptor) WMIDescriptor {
	result := WMIDescriptor{
		Key:          manufacturerCode,
		Manufacturer: manufacturerName,
		Category:     category,
		Descriptor:   descriptor,
	}

	return result
}
