package main
import "fmt";

// Defining a sample type in go lang  
type twoInt struct{
	start int64;
	end int8; 
}

// foo funtion with inputs and return type defined in first line;
func foo(s1 string, s2 string)string{
	return s1 + ", " + s2;
}

// can also return multiple things in tuple  
func updateEndian(index1 int64, index2 int8,endian twoInt)(newendian twoInt, index3 int){
	newendian.start = index1 + endian.start; 
	newendian.end = index2 + endian.end; 
	index3 = int(index1) + int(index2);
	return; 
	 
}

// variadic functions with n number of arguments; 
func printarrayofint(array ...int){
	for index, value:= range(array){
		fmt.Printf("%d at %d\n",value,index);
	}
}

func main() {
	
	// Initializing a constructor kind of initialization for twoInt
	// kind of dictionary based initialization  
	instance1 := twoInt{end:67,start:45};
	instance2 := twoInt{67,45};
	instance3 := twoInt{};

	fmt.Printf("instance1 "instance1,instance2,instance3); 

	fmt.Println("Foo says:",foo("Hello Monty Python", "in the world of golang!"));

	// yet another function 
	instance3, offset := updateEndian(12,34,instance2); 
	fmt.Println("updateEndian from",instance2,instance3,offset);

	// variadic functions 
	array := []int{5,6,7,8};
	printarrayofint(array...);
	println("\n");
	printarrayofint(1,7,8,6,4,33,2);
}

