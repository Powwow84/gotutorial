package main

import (
	// "bufio"
	"fmt"
	"strconv"
	// "log"
	// "os"
	"reflect"
)

var pl = fmt.Println

func main() {
	// pl("hello go")

	// pl ("whats your name")
	// reader := bufio.NewReader(os.Stdin) calls the reader from bufio which prompts an input which is then read
	// name, err := reader.ReadString('\n') the \n means it will stop reading when there a a new line i.e enter
	// if err == nil {
	// 	pl("Hello", name)
	// } else {
	// 	log.Fatal(err)
	// }

	// ******** Variables
	
	// normally variables define with name that begins with letter then can contain letters
	// names that start with caps is conisdered exported and can be accessed outside the package

	// var vName string = "Pao"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "hello"
	// 	// variables are mutable by default meaning values can change as long as data types are the same
	// v4 := 2.4
	// // v4 := 5.4  if you did this it would catch an error, you cant use the colon for reassigning
	// v4 = 5.4

	// *********** data types
	// int, float64, bool, string, rune
	// defautl type 0, 0.0, false, ""

	// pl(reflect.TypeOf((25)))
	// pl(reflect.TypeOf("string"))

	// casting

	
	// cV1 := 1.5 // Changing float to int
	// cV2 := int(cV1) //casting changes the varialbe type

	// cV3 := "5000000000"
	// cv4, err := strconv.Atoi(cV3)  // output converts to string -  Asci to Integer
	// pl(cv4, err, reflect.TypeOf(cV4))

	// cV5 := 5000000000
	// cV6 := strconv.Itoa(cV5)  // output converts to string -  Integer to asci
	// pl(cV6)

	// cV7 := "3.14" //changes string into an int float 64
	// if cV8, err := strconv.ParseFloat(cV7, 64); err == nil { 
	// 	pl(cV8)
	// }

	// cV9 := fmt.Sprintf("%f", 3.14) // formatted printing. converts the float into a string


	//----------- conditionals

	// > < >= <= == !=
	// logical operators && || !

	// iAge := 8
	// if (iAge >= 1) && (iAge <= 18) {
	// 	pl("Important  Birthday")
	// } else if (iAge == 21) ||  (iAge == 50) {
	// 	pl("Important Birthday")
	// } else if iAge >= 65 {
	// 	pl("Important Birthday")
	// } else {
	// 	pl("Not an Important Birthday")
	// }

	// pl("!true =", !true)
}