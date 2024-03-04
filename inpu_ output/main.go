package main

import "frontendmasters.com/go/io/data"

var url = `https://fronted.com`

// global
func main() {
	//function scoped variable
	const y = 2
	var x int = 2
	var message string = "delhi,welcome"
	var x2 = 10
	var string2 = "hello"
	x3 := 10
	print(data.MaxSpeed)
	string3 := "tt"
	{
		var ee = 100
		print(ee)
	}
	Iiiii()
	print("hello", y, x, message, string2, x2, x3, string3, url)
}
