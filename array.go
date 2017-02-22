package main
import(
	"fmt"
	)

func main() {
	
	/* Playing with array */
	var fooArray [3]int;
	fooArray[0] = 2; 
	fooArray[1] = 5;
	fooArray[2] = 7;

	fooArray1 := [3]string{"Amar","Akbar","Anthony"};

	fmt.Println(fooArray, "has length of", len(fooArray));
	fmt.Println(fooArray1, "has length of", len(fooArray1));

	/* Introducing array slice */
	fooSlice := []int{1, 2, 3};
 	fooSlice = append(fooSlice, 5);

 	fmt.Println(fooSlice,fooSlice[1:3]);

 	fooSlice1 := make([]string,3);
 	fooSlice1[0] = "A";
 	fooSlice1[1] = "AA";
 	fooSlice1[2] = "AAA";

 	fmt.Println(fooSlice1);

 	/* Looping  */
 	for i:=0; i < len(fooArray);i++{
 		fmt.Println(fooArray[i]);
 	}

 	for i,value := range(fooArray1){
 		fmt.Println(i,value);
 	} 

}

