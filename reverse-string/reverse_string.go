
//Define the package name
package main

//Import the required modules
import (
	"bufio"
	"fmt"
	"os"
)

//Main function
func main(){
	//Create hte reader variable that will allow us to read from standard input
	reader := bufio.NewReader(os.Stdin)

	//Prompt the user for their input
	fmt.Print("Enter a string to reverse: \n")

	// read the standard input from the user
	text, err := reader.ReadString('\n')

	// Check to see if there was any errors
	if err == nil{

		// Split the text into an array
		runes := []rune(text)

		// get the length of the users input
		length := len(text) - 1

		//Loop though the users input in reverse
		for i := length; i >=0; i -- {
			//Print the letter at index i
			fmt.Printf(string(runes[i]))
		}
		fmt.Println()
	} else{
		// Else if there was an error reading the users input  then output the error message
		fmt.Println("An error has occurred :(")
	}

}

