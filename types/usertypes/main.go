// Package usertypes contains types for user data
package usertypes

// INFO: TYPES FOR USER DATA IN LESSON 5

/*
// Response is a struct for user data received from api
type Response struct {
	Results []UserData `json:"results"`
}

// UserData is a struct for user data received from api
type UserData struct {
	Name Name `json:"name"`
	Dob  Dob  `json:"dob"`
}

// Name is a struct for user name data received from api
type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// Dob is a struct for user dob data received from api
type Dob struct {
	Age int `json:"age"`
}
*/

// NOTE:can also be done concisely with anonymous structs since there is going to be only
// one instance of this struct in a small program

// NOTE: Pitfalls I fell into:
// - do not forget key names in json mapping tags need to be in quotes
// - it's not C-like syntax. [] goes before not after slice name
// - struct not stuct XD
// NOTE: Or separate but made a bit more concise with usage of embedded structs
// mind the fact that UserData[0] will expose FirstName, LastName and Age directly on top level

/*
// Response is a struct for user data received from api
type Response struct {
  UserData `json:"results"`
}

// UserData is a struct for user data with embedded fields for Name and Dob
type UserData []struct {
  Name `json:"name"`
  Dob  `json:"dob"`
}

// Name is a struct for user name data received from the API
type Name struct {
  FirstName string `json:"first"`
  LastName  string `json:"last"`
}

// Dob is a struct for user dob data received from the API
type Dob struct {
  Age int `json:"age"`
}
*/

// Response is a struct for user data received from api
// good for use to build a json structure
type Response struct {
	UserData []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Dob struct {
			Age int `json:"age"`
		} `json:"dob"`
	} `json:"results"`
}
