package values

import "fmt"

func NumCompare(){
	fmt.Println("Works")
	// var(
	// 	a string
	// 	first_letter string
	// )

	// fmt.Print("Input a number to compare: ")
	// fmt.Scanln(&a)

	// if len(a) % 2 == 0{
	// 	first_letter = string(a[0])
	// 	last_letter = string(a)
	// }

}

// func Palindrome(){
// 	var word string
// 	var reverseWord string

// 	fmt.Print("Type a word: ")
// 	fmt.Scanln(&word)

// 	for i := len(word) - 1; i >= 0; i-- {
// 		reverseWord += string(word[i])
// 	}
// 	if word == reverseWord{
// 		fmt.Println("Your given word is palindrome")
// 	} else {
// 		fmt.Println("Your given word is not palindrome")
// 	}
// }

// func Square(){
// 	var (
// 		a, P, S int
// 	)

// 	fmt.Print("Write a value for a side of a square = ")
// 	fmt.Scanln(&a)

// 	P = 4 * a
// 	S = a * a

// 	fmt.Println("The Perimeter is = ", P)
// 	fmt.Println("The Area is = ", S)
// }

// func Beauty(){
// 	var (
// 		date string
// 		count int
// 	)

// 	fmt.Print("Input the date: ")
// 	fmt.Scanln(&date)

// 	for i := 0; i < len(date); i++ {
// 		for j := i + 1; j < len(date); j++ {
// 			// fmt.Println("-",string(date[i]))
// 			// fmt.Println(string(date[j]), "-")
// 			// fmt.Println(string(date[i]),string(date[j]))
// 			if string(date[i]) == string(date[j]){
// 				count += 1
// 			}
// 		}
// 	}

// 	if count > 0 {
// 		fmt.Println("This number is beautiful!")
// 	} else {
// 		fmt.Println("This number is not beautiful!")
// 	}
// }

// func Values() {
// 	var name string = "nimadur"
// 	var age int32 = 123
// 	var boolean bool = true

// 	fmt.Printf("Type: %T Value: %v\n", name, name)
// 	fmt.Printf("Type: %T Value: %v\n", age, age)
// 	fmt.Printf("Type: %T Value: %v\n", boolean, boolean)
// }
