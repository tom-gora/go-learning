// main package entry point
package main

// learning to import local modules
import (
	"fmt"
	/*
		"hello-world/lessons/lesson1"
		"hello-world/lessons/lesson2"
		"hello-world/lessons/lesson3"
		"hello-world/lessons/lesson4"
		"hello-world/lessons/lesson5"
		"hello-world/api"
		"hello-world/lessons/lesson6"
	*/
	"hello-world/lessons/lesson7"
)

func main() {
	// --------------------------------------------------------------------------------------------------
	// LESSON 7
	lesson7.Run()

	// NOTE: mplementation corrected! now logic is private to module and only interface is exported
	// it then can operate on public structs instantiated here as some examples
	// and a method is called in a loop to call a public func making use of the interface's private methods
	//___
	// crucial thing I had to understand and classic example square and circle implementing area() func
	// does NOT make it easy to understand, is that interface is just some set of defined behaviors
	// the shape area suggests interface is some more abstract group containing types,
	// like some higher level eancapsualting struct (shapes contain rects, circles and others)
	// but it is actually better to derive struct name from a verb - shared behavior.
	// I pass behavior declaration (interface) into some function and then any types that implement it
	// can perform its respective implementation within the function

	// NOTE: 2: with that the linguistic confusion cleared on one end but got moved to the other end of this logic.
	// I note to myself - In the func signature like  i.e. `func PrintArea(ac AreaCalculator)` I pass in some verb derived abstract thing (interface),
	// but when actual func call happens a struct instance is being passed `r := rect{ some_props } 	PrintArea(r)
	// -> that was a big part of my confusion!!!!!!! Thus I am making the point to stick in my best practices this note:
	// I think the best Interface descriptor contains concept of what structs it would partake to and also behavior.
	// Ie FilePrinter, VenueCostCalculator, ShapeAreaCalculator etc it makes it most cleraly obvious to a person reading the code
	// what types would a function receiving this interface likely deal with and what it would do with them.

	venue1 := lesson7.Restaurant{
		Name:              "Madzia's Dream Feast",
		FloorAreaInMeters: 1200,
		RentPricePerMeter: 300.99,
	}

	venue2 := lesson7.Hotel{
		Name:              "The Grand Budapest",
		FloorAreaInMeters: 12000,
		RentPricePerMeter: 1000.99,
	}

	venue3 := lesson7.Cinema{
		Name:              "CinemaShitty",
		FloorAreaInMeters: 7500,
		RentPricePerMeter: 549.99,
	}

	venuesList := []lesson7.VenueDetailPrinter{venue1, venue2, venue3}

	for _, v := range venuesList {
		lesson7.PrintDetails(v)
		fmt.Println()
	}

	// part two cops:
	// NOTE: Good code practiced: - explicitly naming interface params, using redable type assertion switch

	goodSentences := []string{
		"Hey, we're not here to make your life difficult. We just want to get to the truth.",
		"Listen, I know you're in a tough spot. Help us out, and we can help you.",
		"I get it, you're scared. But the sooner you tell us what happened, the sooner you can get out of here.",
		"You're a good person, I can tell. Don't let one mistake ruin your life.",
		"We know you're not a bad guy. Let's just clear this up so you can go home.",
	}

	badSentences := []string{
		"You're not fooling anyone. We know you're involved. So why don't you just cut the crap?",
		"If you think you’re walking out of here without telling us everything, you’re dead wrong.",
		"I'm done playing nice. Start talking, or you're gonna wish you had.",
		"You think we don’t have enough to lock you up? One word from me, and you're finished.",
		"Keep lying, and you'll regret it. We can do this the easy way or the hard way.",
	}

	terminatorSentences := []string{
		"Hasta la vista, baby... and your alibi.",
		"I'll be back for more information about this incident... repeatedly.",
		"Your story is like a virus - it's been terminated.",
		"You're not very good at this game... I'm going to have to reboot you.",
		"The machine has spoken... and it says you're lying.",
	}

	cop1 := lesson7.BadCop{Name: "Terry"}
	cop2 := lesson7.GoodCop{Name: "Brandon"}
	cop3 := lesson7.Terminator{Name: "Arnie"}

	fmt.Println("\nGood cop goes in:")
	lesson7.SendCopIntoRoom(cop2, goodSentences, badSentences, terminatorSentences)
	fmt.Println("\nBad cop goes in:")
	lesson7.SendCopIntoRoom(cop1, goodSentences, badSentences, terminatorSentences)
	fmt.Println("\nFuck the good cops, let's send the terminator in!")
	lesson7.SendCopIntoRoom(cop3, goodSentences, badSentences, terminatorSentences)

	// NOTE:  API DETOUR
	// api.Run()

	// --------------------------------------------------------------------------------------------------
	// LESSON 6
	// lesson6.Run()
	/*
		myCar := struct {
		Make  string
		Model string
		Age   int
		}{Make: "Dacia", Model: "Sandero Stepway", Age: 2}
		fmt.Printf("My car is a %s %s and it is %d years old\n", myCar.Make, myCar.Model, myCar.Age)
	*/

	// NOTE: What I learned is name - composite literals. Basically in-place filling in a data structure on the same line it is declared
	// used for structs but also more commonly for slices, maps and arrays
	// scope those with {}
	/*
		someNumbers := []int{1, 4, 6, 32, 6}
		str := "\n"
		for _, n := range someNumbers {
		str += fmt.Sprintf("%d, ", n)
		}
		str = str[:len(str)-2]
		fmt.Println(str)
	*/

	// NOTE: A little detour to understand how slicing works by trying out
	// also implemented with returned custom errors to learn good practice
	/*
		sliceAndDiceTest := "TrevorTheRelentless"
		sliceAndDiceTest, err := lesson6.DoASlice(sliceAndDiceTest, "start", 5)
		if err != nil {
		fmt.Println(err)
		}
		fmt.Println(sliceAndDiceTest)
		sliceAndDiceTest, err = lesson6.DoASlice(sliceAndDiceTest, "end", 6)
		if err != nil {
		fmt.Println(err)
		}
		fmt.Println(sliceAndDiceTest)
	*/

	// NOTE: Declared some method on a custom struct
	/*
		myCat := lesson6.Cat{
				Name:      "Mooncake",
				Age:       4,
				Lovliness: 0.2,
				HasEaten:  false,
			}

			myCat.Feed()
			myCat.Feed()
			myCat.Feed()
			myCat.Feed()
	*/
	// --------------------------------------------------------------------------------------------------
	// LESSON 5
	/*
		lesson5.Run()
		num := 10
		lesson5.Increment(num)
		fmt.Println(num)
		// 
		// NOTE:  Passed by value so a copy is passed into the func and num in this scope remains unchanged
		//          & _
		//        To get result value from func need to be stored as brand new copy (variable)
		//        
		num = lesson5.IncrementWithReturn(num)
		fmt.Println(num)
	*/
	// NOTE:: While learning multiple returns why not learn fetching api data at the same time?
	/*
		fmt.Printf("My name is %s %s and I am %d years old", fName, lName, age)
		fName, lName, age := lesson5.MultipleReturns()
	*/
	// NOTE: Sooooo from what I learned over there in this function called above:
	// __
	// Declare structs for each bit of json I want to extract
	// (learned to extract type declarations into separate modules to keep logic clear away from declarative bits)
	// and they will stack like matryoishka reflecting json structure.
	// Rather than rebuilding entire response json as a massive struct User,
	// just break it into precise encapsulated types that I need and use this fancy json:"key" syntax
	// to extract json keys for each struct property. Then at the end just assign variables
	// by accessing nested types almost akin to traversing javascript object structure.

	// NOTE: Also learned that mapping with interfaces in go is literally hellish spawn comparable only
	// to single page static blogs made with jquery with embedded java applets from 20 years ago
	// still run in IE in governmets around the world. Man I doubt modern java is half as verbose
	// convoluted and fucked in the brain as this method of performing abstract unholy rituals just to map some json
	// data to variables. GO you were supposed to be the chosen one. You were supposed to bring the simplicity
	// to the force!

	// --------------------------------------------------------------------------------------------------
	// LESSON 4
	// lesson4.Run()

	// NOTE: IF statements can be written with an additional initialization section where variables can be setup
	// scoped ONLY to this condition (similar to say traditional var i in loop blocks)
	// makes it safer and nicer as data is block scoped and GC'd after evaluation

	//                    So in the func we pass max-min and then add min to the output to correct for the lower bound
	//                  󱞢 This is using rand.Intn() that takes one int and returns int  from 0 to that int( not including)
	/*  if i := lesson4.ReturnRandomInt(1, 69); i == 42 {
	fmt.Println("The universe is saved!")
	} else {
	fmt.Println("The universe is doomed!")
	}
	*/
	// NOTE: and `i` is not accessible anywhere but this scope! Amazing!
	// ye the code below won't even run because i is undefined
	/*
		test, err := fmt.Sprint(i)
		if err != nil {
		fmt.Println("Error:", err)
		return
		}
		fmt.Println(test)
	*/
	/*
		NOTE: go does not do while loop. `for` is used do declare continuous loops
		it's a bit different syntax and paradigm but basically as I understand it:
		- just use for as equivalent to `while(true)` and open scoped block right after
		- declare breakign conditions as guard statement(s) (can use the init syntax as above)
		- write the code that needs to run while condition(s) is/are not met
	*/
	// fmt.Printf("---------------------------------------------\nAnd looping below:\n")

	// NOTE: Below will not work because i is scoped to the guard clause and not accessible
	// elsewhere in the loop
	/*
		for {
		if i := lesson4.ReturnRandomInt(1, 69); i == 42 {
		fmt.Printf("Attempt %d. The universe is saved!", i)
		break
		}

		fmt.Printf("Attempt %d. The universe is doomed!", i)
		}
		NOTE: INSTEAD DO*/
	// also ansii escape secquences work as expeced YASS!

	/*
		counter := 1
		for {
		i := lesson4.ReturnRandomInt(1, 69)
		if i == 42 {
		fmt.Printf("\033[32mPicked number %d\033[0m\n", i)
		fmt.Printf("Attempt %d. The universe is \033[35msaved!\033[0m\n", counter)
		break
		}
		fmt.Printf("\033[32mPicked number %d\033[0m\n", i)
		fmt.Printf("Attempt %d. The universe is doomed!\n", counter)
		counter++
		}
	*/

	// --------------------------------------------------------------------------------------------------
	// LESSON 3

	/*
		lesson3.Run()
		simple assignemts as var
		testFloat := 15.876
		consts are a thing here
		const testFloat float64 = 15.876
		aboveToInt := lesson3.FloatToIntTruncation(testFloat)
		fmt.Println(fmt.Sprintf(
		"\nFloat %.3f converted to int %d",
		testFloat,
		aboveToInt,
		))

	*/
	// const constantA = 10
	// const constantB = 20
	//
	// // NOTE: this ain't javascript bro. functions are only evaluated at runtime and  cannot produce constants
	//
	// const constSum = lesson3.FunWithConsts(constantA, constantB)
	// const constSum = constantA + constantB
	//
	// NOTE: so - constants can be computed in scope but cannot be evaluated by remotely called function

	/*
		fmt.Printf("Constant value - computed \"inline\": %v\n", constSum)
		regularIntA := 10
		regularIntB := 20
		// NOTE: this can be computed by a function

		regularSum := lesson3.Sum(regularIntA, regularIntB)
		fmt.Printf(
		"And this is result of regular adding of vars by a function: %v\n",
		regularSum,
		)
	*/
	// NOTE: same with strings

	/*
		regularStr1 := "A świstak siedzi"
		regularStr2 := "i koduje te jebane sreberka"

		regularConcatenatedStr := lesson3.Concat(regularStr1, regularStr2)
		fmt.Println(regularConcatenatedStr + "\n")
	*/
	// NOTE:  and interpolation into a var with sprintf

	/*
		name := "Jimmy"
		height := 4.67
		interpolatedStr := fmt.Sprintf(
		"My name is %s and I am %.2f feet tall.",
		name,
		height,
		)
		fmt.Println(interpolatedStr)
	*/

	// --------------------------------------------------------------------------------------------------
	// LESSON 2

	// assigning variables on one line with type inference
	// averageWill, taskDescription := .002, "is the rate of how strong my average will to go work in Tesco is."
	// fmt.Println(lesson2.Run(averageWill, taskDescription))
	/*
		numAsString := "12354353"
		sum := 0
		arrASCII := lesson2.TryAndSeeIfFlexibleTypeCastingIsAThingInGo(numAsString)
		arrInts := make([]int, len(arrASCII))
	*/

	/*
		for _, c := range arrASCII {
		c, err := strconv.Atoi(string(rune(c)))
		if err != nil {
		fmt.Println("Error:", err)
		return
		}
		arrInts = append(arrInts, c)
		}
		for _, i := range arrInts {
		sum = sum + i
		}
	*/

	/*
		fmt.Printf(
		"I converted %T (\"%s\") to an array (%T) of %Ts (ascii codes)\nand then converted them as runes back to strings to convert to individual %Ts\nand finally added each digit to get a sum of type %T equal to %d",
		numAsString,
		numAsString,
		arrASCII,
		arrASCII[0],
		arrInts[0],
		sum,
		sum,
		)
		fmt.Print("\n\n")
	*/

	// yea skip past blah blah what are viariables data types blah blah
	// let's play a bit with how is a go project organized namespaces there, modules elsewhere and here we have packages

	// --------------------------------------------------------------------------------------------------
	// LESSON 1

	// call imported module
	/*
		lesson1.Run()
		lesson1.DoSomethingElse()
		divisionOutcome := lesson1.Divide(1322, 23)
		fmt.Println(divisionOutcome)
		fmt.Println(math.Round((divisionOutcome * 100) / 100))
	*/
	fmt.Print("\n\n              \n         \033[32mLearning GO")
	fmt.Print("\n         ,_---~~~~~----._         \n  _,,_,*^____      _____``*g*\"*, \n / __/ /'     ^.  /      \\ ^@q   f \n[  @f | @))    |  | @))   l  0 _/  \n \\'/   \\~____ / __ \\_____/    \\   \n  |           _l__l_           I   \n  }          [______]           I  \n  ]            | | |            |  \n  ]             ~ ~             |  \n  |                            |   \n   |                           |   \033[0m")
}
