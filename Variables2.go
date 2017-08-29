package main

import "fmt"

func main() {
    a := 32       // initialize a variable
    fmt.Printf("This line is using Printf: %d is of type %T", a,a) // This doesn't create a new line once the program is completed. We need to explicitly define the \n which is for creating the new line. Please refer to 3rd print statement for more clarity.

    fmt.Println("This line is using Println: %d is of type %T", a,a) // This by default goes to the new line once the program is completed.

    fmt.Printf("This line is using Printf: %d is of type %T \n", a,a)
}
