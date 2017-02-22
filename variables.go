package main
import(
	"fmt"
	"strconv"
	)

/* Declaring variables */  
var fooString1 string ="ho ho";  
var fooString2, fooString3 = "ram" , "bo";
const (
	one = "yo";
	two = "mama"
	three = 3;
)

func main() {
	/* different ways to declare variables */
	fooString4 := "12"; 
	const sampleInputs = 12

	fmt.Printf("%s %s %s %s %s \n",fooString4, fooString2,fooString3,one,two);

	/* returns a tuple, _ is for error message */ 
	integer, _ := strconv.Atoi(fooString4)
	fmt.Println( integer - sampleInputs + 3 );
}

