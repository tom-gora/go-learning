// Package lesson7 is a package that contains lesson 7 code
package lesson7

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	// "strings"
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

type typeExtractor interface {
	getType() (typeAsString string)
}

// Hotel is one type of possible venue
type Hotel struct {
	Name              string
	FloorAreaInMeters float64
	RentPricePerMeter float64
}

// Restaurant is one type of possible venue
type Restaurant struct {
	Name              string
	FloorAreaInMeters float64
	RentPricePerMeter float64
}

// Cinema is one type of possible venue
type Cinema struct {
	Name              string
	FloorAreaInMeters float64
	RentPricePerMeter float64
}

func (h Hotel) hostEvent() {
	fmt.Printf("%s \"%s\" is hosting a speed dating session tonight.\n", h.getType(), h.Name)
}

func (r Restaurant) hostEvent() {
	fmt.Printf("%s \"%s\" is hosting a ball tonight.\n", r.getType(), r.Name)
}

func (c Cinema) hostEvent() {
	fmt.Printf("%s \"%s\" is hosting a horror marathon tonight.\n", c.getType(), c.Name)
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

func (h Hotel) getType() string {
	hType := reflect.TypeOf(h).String()
	return strings.Split(hType, ".")[1]
}

func (r Restaurant) getType() string {
	rType := reflect.TypeOf(r).String()
	return strings.Split(rType, ".")[1]
}

func (c Cinema) getType() string {
	cType := fmt.Sprintf("%T", c)
	return strings.Split(cType, ".")[1]
}

// in getters 2 ways of extracting type as string: using fmt and reflect
func (h Hotel) getNameAndType() (string, string) {
	return h.Name, h.getType()
}

func (r Restaurant) getNameAndType() (string, string) {
	return r.Name, r.getType()
}

func (c Cinema) getNameAndType() (string, string) {
	return c.Name, c.getType()
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

func (gc GoodCop) interogate(s []string) (formattedStr string) {
	random := rand.Intn(len(s))
	return fmt.Sprintf("%s says: \"%s\"", gc.Name, s[random])
}

// BadCop is a type of possible cop
type BadCop struct {
	Name string
}

func (gc BadCop) interogate(s []string) (formattedStr string) {
	random := rand.Intn(len(s))
	return fmt.Sprintf("%s says: \"%s\"", gc.Name, s[random])
}

type interogator interface {
	interogate(interrogationSentences []string) (formattedStr string)
}

// Terminator is a type of possible cop
type Terminator struct {
	Name string
}

func (t Terminator) interogate(s []string) string {
	random := rand.Intn(len(s))
	return fmt.Sprintf("%s says: \"%s\"", t.Name, s[random])
}

// SendCopIntoRoom is a function exporting module private func results
// rewritten to do type assertion switch rather than simple if checks
func SendCopIntoRoom(i interogator, gs []string, bs []string, ts []string) {
	switch v := i.(type) {
	case GoodCop:
		fmt.Println(v.interogate(gs))
	case BadCop:
		fmt.Println(v.interogate(bs))
	case Terminator:
		fmt.Println(v.interogate(ts))
	default:
		fmt.Println("The interrogation has been cancelled.")
	}
}
