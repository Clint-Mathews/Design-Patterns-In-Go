package main

import "fmt"

// The Adapter design pattern in Go is a structural pattern that allows incompatible interfaces to work together seamlessly

// Plug type by countries for reference
// Type A - Canada, United States, Japan, and Mexico
// Type B - Canada, United States, and Mexico
// Type C - widely used throughout Asia, Europe, and South America
// etc...

// Target interface
type universalTravelAdapter interface {
	plugInfo()
}

// Concrete prototype implementation
type typeA struct {
	plugTypeCountries string
}

func (plug *typeA) plugInfo() {
	fmt.Printf("Plug type is availble in countries: %s\n", plug.plugTypeCountries)
}

// Client
type client struct{}

func (c *client) plugInfo(u universalTravelAdapter) {
	u.plugInfo()
}

// Adaptee which does not follow Target interface
type typeB struct {
	plugTypeCountries string
	voltage           float64
}

func (plug *typeB) plugInfoTypeB() {
	fmt.Printf("Plug type is availble in countries: %s\n", plug.plugTypeCountries)
}

func (plug *typeB) getVolatge() float64 {
	return plug.voltage
}

// Adpater
type universalTravelAdapterAdapter struct {
	plug typeB
}

func (u *universalTravelAdapterAdapter) plugInfo() {
	u.plug.plugInfoTypeB()
}

// Main
func main() {
	fmt.Println("Adapter-Design-Pattern")

	// Using the power outlet directly without using client for verifying
	powerOutlet := &typeA{plugTypeCountries: "Canada, United States, Japan, and Mexico"}
	powerOutlet.plugInfo()

	// Using power outlet via client
	client := &client{}
	client.plugInfo(powerOutlet)

	// Use plug type 2, extended requirement
	typeBpowerOutlet := typeB{plugTypeCountries: "Canada, United States, and Mexico"}

	// Using type 2 typeBpowerOutlet directly would throw error, since plugInfo is not implemented by typeB!
	// client.plugInfo(typeBpowerOutlet)
	// So we can use the Adapter here
	adapter := &universalTravelAdapterAdapter{plug: typeBpowerOutlet}
	client.plugInfo(adapter)
}
