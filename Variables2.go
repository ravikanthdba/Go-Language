package main

import "fmt"

func main() {
    /* This method is called initialization */

     a := 32       // initialize a variable
     fmt.Println("This line is using Println: %d is of type %T", a,a) // This by default goes to the new line once the program is completed.
     fmt.Printf("This line is using Printf: %d is of type %T\n", a,a) // This doesn't create a new line once the program is completed. We need to explicitly define the \n which is for creating the new line. Please refer to 3rd print statement for more clarity.

     // This method is called declaration

     var str string   
     str = "This program describes the different types of Variable declaration"
     fmt.Println("This line is using Println command. %s is of type %T", str,str )
     fmt.Printf("This line is using Printf command. %s is of type %T\n", str,str)

     var decimal float32 = 1.23456789
     fmt.Println("The decimal value %f is of type %T", decimal,decimal)
     fmt.Printf("The decimal value %f is of type %T\n", decimal,decimal)


     /* Notes : When we print the statement through "Println" instead of "Printf", %d, %T, or any formatting values will not be taken into consideration. Only with "Printf", we see the %f, %T into consideration. See the execution of this program for more details. */


}
