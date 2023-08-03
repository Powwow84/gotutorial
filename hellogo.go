package main

import (
	// "bufio"
	"fmt"
	// "strconv"
	// "strings"
	// // "unicode/utf8"

	// "log"
	// "os"
	// // "reflect"
	// "time"
	// // "math"
	// "math/rand"
)

var pl = fmt.Println

func sayHello(){
	pl("Hello")
}

func getSum(x int, y int) int {
	sum := x + y
	pl(sum)
	return sum // if you have a return type you need to have the return statement in there
}

func getTwo(x int) (int, int) { //if you need to return two return types back from a function you need to wrap the return type in ()
	return x+1, x+2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

func getMore(nums ...int) int { //syntax allows for an undeclared number of parameters -> getMore(1,2,3,4,5,6,7)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int { // using an array as a parameter. note pass by value will not change the contents of the array outside of the function when used inside of a function
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func changeVal(myPtr *int) { //to change the value of a passed by value you need to use a pointer  '*'  and the address '&' changeVal(&f3). to store a pbv 'var variable *int = some_value'
	*myPtr = 12
}



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


	// ---------------------------strings

	// sv1 := "A word"
	// replacer := strings.NewReplacer("A", "Another" )
	// sV2 := replacer.Replace(sV1)
	// pl(sv2)

	// pl("Length :", len(sV2))
	// pl("Contains Another :", strings.Contains(sV2,"Another"))
	// pl("o index:", strings.Index(sV2, "o"))
	// pl("Replace :", strings.Replace(sV2, "o", "O", -1)) // -1 replaces everything, if you had a 2 it would replace the first two
	// sV3 := "\nSome Words\n" // the  \ lead into an escape sequence, lets you use things insdie of quotes like """ " 
	// sV3 = strings.TrimSpace(sV3) //trims whitespace
	// pl("Split :", strings.Split("a-b-c-d", "-"))
	// pl("lower :", strings.ToLower(sV2))
	// pl("upper :", strings.ToUpper(sV2))
	// pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	// pl("Suffix :", strings.HasSuffix("tacocat", "cat"))
	
	// ---------------runes
	// characters are called runes, runes are unicode that represents characters

	// rStr := "abcdefg"
	// pl("Rune COunt :", utf8.RuneCountInString(rStr))
	// for i, runeVal := range rStr{
	// 	fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal ) //<- formated printing. Note whenever you use Printf you need to use \n to break it into a new line if you want a new line
	// 	// returns this
	// 	// Rune COunt : 7
	// 	// 0 : U+0061 'a' : a
	// 	// 1 : U+0062 'b' : b
	// 	// 2 : U+0063 'c' : c
	// 	// 3 : U+0064 'd' : d
	// 	// 4 : U+0065 'e' : e
	// 	// 5 : U+0066 'f' : f
	// 	// 6 : U+0067 'g' : g
	// }


	// -------------------Time
	// now := time.Now()
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())

	// // -----------------operators

	// pl("5+4 =" , 5+4)
	// pl("5-4 =" , 5-4)
	// pl("5*4 =" , 5*4)
	// pl("5/4 =" , 5/4)
	// pl("5%4 =" , 5%4)
	// mInt := 1
	// mInt++
	// mInt += 1
	// mInt--
	// mInt -=

	// -----------------Mathmathical functions
		// Most commonly used

	// pl("Abs(-10) =", math.Abs(-10))
	// pl("Pow(4, 2) =", math.Pow(4, 2))
	// pl("Cbrt(8) =", math.Cbrt(8))
	// pl("Ceil(4.4) =", math.Ceil(4.4))
	// pl("floor(4.4) =", math.Floor(4.4))
	// pl("Round(4.4) =", math.Round(4.4))
	// pl("Log2(8) =", math.Log2(8))
	// pl("Log10(100) =", math.Log10(100))
	// // get the log of e to the power of 2
	// pl("Log(7.389) =", math.Log(math.Exp(2)))
	// pl("Max(5,4) =", math.Max(5,4))
	// pl("Min(5,4) =", math.Min(5,4))
	// // converts 90 degrees into radians
	// r90 := 90 * math.Pi / 180
	// // converts radians into degrees
	// d90 := r90 * (180 / math.Pi)
	// fmt.Printf("%f radians = %f degrees\n", r90, d90)
	// pl("Sin(90) =", math.Sin(r90))
	// there are also functions dor cos, tan, acos, asin
	// Atan, AsinH, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos, Htpos

	// Abs(-10) = 10 (absolute value)
	// Pow(4, 2) = 16 (power of)
	// Cbrt(8) = 2 (cube)
	// Ceil(4.4) = 5 (ceiling)
	// floor(4.4) = 4 (floor)
	// Round(4.4) = 4 (round up or down)
	// Log2(8) = 3 he math.Log2 function takes a single argument of type float64 and returns the base-2 logarithm as a float64 value. In this case, the result would be 3 because 2 raised to the power of 3 equals 8.
	// Log10(100) = 2 unction takes a single argument of type float64 and returns the base-10 logarithm as a float64 value. In this case, the result would be 2 because 10 raised to the power of 2 equals 100.
	// Log(7.389) = 2 The math.Log function takes a single argument of type float64 and returns the natural logarithm as a float64 value. In this case, the result would be approximately 2 because e raised to the power of 2 equals 7.389, where e is Euler's number approximately equal to 2.71828.
	// Max(5,4) = 5
	// Min(5,4) = 4

	

	// // float precision
	// pl("Float Precision = ", 0.11111111111 + 0.11111111111) // important to understand percision when dealing with floats. when doing the calculation it will hold all the decimal places

	// ---------------random numbers

	// seedSecs := time.Now().Unix() //returns seconds since 1/1/70 as an int
	// random.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1 // generates a random number exclusive of 50 (0 -49). So we add the +1 to generate it to 0 - 50
	// pl("Random : ", randNum)

	//  -------- formatted print

	//  %d : Integer
	//  %c : Character
	//  %f : Float
	//  %t : Boolean
	//  %s : String
	//  %o : Base 8
	//  %x : Base 16
	//  %v : Guesses based on data type
	//  %T : Type of supplied value

	// fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	// fmt.Printf("%9f\n", 3.14)
	// fmt.Printf("%.2f\n", 3.141592)
	// fmt.Printf("%9.f\n", 3.141592)

	// sp1 := fmt.Sprintf("%9.f\n",3.141592)
	// pl(sp1)

	// ------------- conditionals

	//  for inititiazation; condition; poststatment {BODY}
	// for x := 1; x <= 5; x++ {
	// 	fmt.Println(x)
	// }

	// fX := 0
	// for fX < 5 { // while loop
	// 	pl(fX)
	// 	fX++
	// }

	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) +1
	// for true {
	// 	fmt.Println("Guess a number between 0 and 50")
	// 	fmt.Println("random number is %d", randNum)
	// 	reader := bufio.NewReader(os.Stdin)
	// 	guess, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	guess = strings.TrimSpace(guess)
	// 	iGuess, err := strconv.Atoi(guess)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if iGuess > randNum {
	// 		fmt.Println("enter a lower number")
	// 	} else if iGuess < randNum {
	// 		fmt.Println("guess a higher number")
	// 	} else {
	// 		fmt.Println("Your Right")
	// 		break
	// 	}
		
	// }

	//  ---- loop through an array
	// aNums := []int{1,2,3}
	// for _, num := range aNums {
	// 	pl(num)
	// }

	// ----------------------Arrays

	// var arr1 [5]int //an array that has a declared size of 5
	// arr1[0] = 1
	// arr2 := [5]int{1,2,3,4,5} // declaring an array of size 5 with 5 values in the array
	// pl("Index 0:", arr2[0]) //prints whats at arr1 index 0. inthis case the value is 1
	// pl("Arr Length :", len(arr2))
	// for i := 0 ; i < len(arr2) ; i++ {
	// 	pl(arr2[i])
	// }
	// for i, v := range arr2 {
	// 	fmt.Printf("%d : %d \n", i, v)
	// }
	// arr3 := [2][2]int{
	// 	{1,2},
	// 	{3,4},
	// }
	// for i := 0 ; i < 2 ; i++ {
	// 	for j := 0 ; j < 2 ; j++ {
	// 		pl(arr3[i][j])
	// 	}
	// }

	// aStr1 := "abcde"
	// rArr := []rune(aStr1)
	// for _, v := range rArr {
	// 	fmt.Printf("Rune Array : %d\n", v)
	// }
	// byteArr := []byte{'a', 'b', 'c'}
	// pl(byteArr)
	// bStr := string(byteArr[:]) //the [:] is referencing a slice of the slice for instance you can do [1:3] so it represents a slice from index 1 and ending in index 2 not inclusive of 3
	// pl("I'm a string :", bStr)

	// --------slices
	// slices are like arrays but can grow in size

	// sl1 := make([]string, 6)
	// sl1[0] = "Society"
	// sl1[1] = "of"
	// sl1[2] = "the"
	// sl1[3] = "Simulated"
	// sl1[4] = "Universe"
	// pl("Slice Size :", len(sl1))
	// for i := 0; i < len(sl1); i++ {
	// 	pl(sl1[i])
	// }
	// for i, x := range sl1 {
	// 	pl(x, i)
	// }

	// sArr := [5]int{1,2,3,4,5}
	// sl3 := sArr[0:2]
	// pl(sl3)
	// pl("1st 3 numbers of the array: ", sArr[:3])
	// pl("slice from index 2 to the end: ", sArr[2:])
	// sArr[0] = 10 //changing the array changes the slice that referenced it
	// pl("sl3 : ", sl3)
	// sl3[0] = 1 //changing the slice changes the array
	// pl("sArr :", sArr)

	// sl3 = append(sl3, 12) //appending to a slice will append will modify the array that was referenced, in this example 12 is appeneded to the end of the slice, but it changes index 2 of the referenced array
	// pl("sl3 :", sl3)
	// pl("sArr :", sArr)

	// sl4 := make([]string, 6) //empty slices will out put an empty array
	// pl("sl4 :", sl4) //this will output an empty array
	// pl("sl4[0] :", sl4[0]) // this will output nil sl4[0] :

	//  ------------------functions
	//func funcName(parameters) returnType {BODY}
	// ********************** Go does not support nested functions. declaring a function in here is goin to throw an error

	// func sayHello(){
	// 	pl("Hello")
	// }

	// sayHello() // to call the function
	// getSum(5, 6)

	// ------ Pointers

	// func getArraySum(arr []int) int { // using an array as a parameter. note pass by value will not change the contents of the array outside of the function when used inside of a function
	// 	sum := 0
	// 	for _, val := range arr {
	// 		sum += val
	// 	}
	// 	return sum
	// }

	f4 := 10
	var f3 int
	var f4PTR * int = &f4
	pl("f4 addess :", f4PTR)
	pl("f4 Value :", *f4PTR)
	*f4PTR = 11 //using pointer to directly change a value
	pl("f4 Value :", *f4PTR)

	pl("f3 before func :", f3)
	changeVal(&f3) //changes the value as the argument is passed in
	pl("f3 after func :", f3)

// 	func changeVal(myPtr *int) { //to change the value of a passed by value you need to use a pointer  '*'  and the address '&' changeVal(&f3). to store a pbv 'var variable *int = some_value'
// 	*myPtr = 12
// }

	// func dblArrVals(arr *[4]int){ //passing an array through a function with a pointer to change the value of the array with the pointer and the address
	// 	for x := 0 ; x < 4; x++ {
	// 		arr[x] *= 2
	// 	}
	// }

	// pArr := [4]int{1,2,3,4}
	// dblArrVals(&pArr)
	// pl(pArr)

	// func getAverages(nums ...float64) float64{
	// 	var sum float64 = 0.0
	// 	var numSize float64 = float64(len(nums))
	// 	for _, val := range nums {
	// 		sum += val
	// 	}
	// 	return(sum/numSize)
	// }

	// iSlice := []float64{11,13,17}
	// fmt.Printf("Averages : %.3f\n", getAverages(iSlice...)) //the .3f limits the return to 3 decimals points
}