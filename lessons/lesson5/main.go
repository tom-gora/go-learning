// Package lesson5 contains functions for lesson 5
package lesson5

import (
	"encoding/json"
	"fmt"
	types "hello-world/types/usertypes"
	"io"
	"net/http"
)

// Run print welcomem msg
func Run() {
	fmt.Println("Hello from lesson 5")
}

// Increment increment in place no return
func Increment(x int) {
	x++
}

// IncrementWithReturn with return allowing assigning to var
func IncrementWithReturn(x int) int {
	return x + 1
}

// MultipleReturns fetches data from api and returns multiple values
func MultipleReturns() (string, string, int) {
	// NOTE: For executing most basic api call (no tokens, no keys, no any headers)
	// 1. Declare url
	apiURL := "https://randomuser.me/api/"
	// 2. Create request (http package) params are: http verb, url, request body (if any)
	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Print("Error making request: ", err.Error())
	}

	// 3. exec the request using DefaultClient instance of the  Client struct in http module
	// which does the request itself. think like inbuilt curl
	// it's method Do() takes one param - declared instance of Request
	// returns response as struct Response
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Print("Error fetching response: ", err.Error())
	}
	// close the connection, tidy up resources
	//(note it is implemented by Body not Response struct itself. Design choice I guess)
	// mnemonics >: you request body from John Wick, you receive a body in response. Now body itself needs tidying up
	defer response.Body.Close()

	// i/o std lib pakage provides a general reader method to read and store response.body bytes
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print("Error reading response: ", err.Error())
	}

	// declare instance of prepared custom struct for storing required data
	var res types.Response
	// NOTE: to init empty struct add empty scope to its name {}
	// this makes it possible to use shorthand syntax for initializing
	// 
	// res := usertypes.Response{}

	// NOTE: !!! and I just learned and understood the following:
	// go does passing by value by default creating data copies,
	// so if I explicitly need to use pointer to DO modify whatever a variable represents
	// (by passing by reference), I need to prefix it with `&`
	// also doing some thinking I realized it is an impure "void-like" function. Other than err
	// it returns nothing and it mutates the state of one of its input args (2nd)
	jsonErr := json.Unmarshal([]byte(body), &res)
	if jsonErr != nil {
		fmt.Print(":Error processing response: ", err.Error())
	}

	// 4. extract data from response body and store in vars then return them
	fName := res.UserData[0].Name.First
	lName := res.UserData[0].Name.Last
	age := int(res.UserData[0].Dob.Age)

	return fName, lName, age
}
