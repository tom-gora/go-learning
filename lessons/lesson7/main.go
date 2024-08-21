// Package lesson7 is a package that contains lesson 7 code
package lesson7

import (
	"fmt"
	"math/rand"
)

// Run is a function printing lesson introduction
func Run() {
	fmt.Print("Hello from lesson 7\nThis time we are doing basics of interfaces:\n\n")
}

// VenueDetailPrinter interface - a collection of methods that you want all types to be able to call
type VenueDetailPrinter interface {
	hostEvent()
	getNameAndType() (string, string)
	getRent() float64
}

// Hotel is one type of possible venue
type Hotel struct {
	Type              string
	Name              string
	FloorAreaInMeters float64
	RentPricePerMeter float64
}

// Restaurant is one type of possible venue
type Restaurant struct {
	Type              string
	Name              string
	FloorAreaInMeters float64
	RentPricePerMeter float64
}

// Cinema is one type of possible venue
type Cinema struct {
	Type              string
	Name              string
	FloorAreaInMeters float64
	RentPricePerMeter float64
}

func (h Hotel) hostEvent() {
	fmt.Printf("%s \"%s\" is hosting a speed dating session tonight.\n", h.Type, h.Name)
}

func (r Restaurant) hostEvent() {
	fmt.Printf("%s \"%s\" is hosting a ball tonight.\n", r.Type, r.Name)
}

func (c Cinema) hostEvent() {
	fmt.Printf("%s \"%s\" is hosting a horror marathon tonight.\n", c.Type, c.Name)
}

func (h Hotel) getRent() float64 {
	return h.RentPricePerMeter * h.FloorAreaInMeters
}

func (r Restaurant) getRent() float64 {
	return r.RentPricePerMeter * r.FloorAreaInMeters
}

func (c Cinema) getRent() float64 {
	return c.RentPricePerMeter * c.FloorAreaInMeters
}

func (h Hotel) getNameAndType() (string, string) {
	return h.Name, h.Type
}

func (r Restaurant) getNameAndType() (string, string) {
	return r.Name, r.Type
}

func (c Cinema) getNameAndType() (string, string) {
	return c.Name, c.Type
}

// PrintDetails is a function exporting module private func results
func PrintDetails(vdp VenueDetailPrinter) {
	vdp.hostEvent()
	rentPerMonth := vdp.getRent()
	n, t := vdp.getNameAndType()
	fmt.Printf("The %s %s costs the owner £%.2f per month.\n", t, n, rentPerMonth)
}

// second part just to try
// NOTE: COPS

// GoodCop is a type of possible cop
type GoodCop struct {
	Name string
}

// Interogate is a good cop action
func (gc GoodCop) Interogate(s []string) {
	random := rand.Intn(len(s))
	fmt.Println(gc.Name + " says: \"" + s[random] + "\"")
}

// BadCop is a type of possible cop
type BadCop struct {
	Name string
}

// Interogate is a bad cop action
func (gc BadCop) Interogate(s []string) {
	random := rand.Intn(len(s))
	fmt.Println(gc.Name + " says: \"" + s[random] + "\"")
}

type interogator interface {
	Interogate(s []string)
}

// RandomizeInterogation is a function exporting module private func results
func RandomizeInterogation(i interogator, gs []string, bs []string) {
	_, ok := i.(GoodCop)
	if ok {
		i.Interogate(gs)
	} else {
		i.Interogate(bs)
	}
}
