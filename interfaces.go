package main
import(
	"fmt"
	)
func main() {
	
	instance1 := twoInt{end:67,little:45};
	instance2 := new(twoInt);
	fmt.Println(instance2.add());
}

// Defining Types in go 
type twoInt struct{
	little int64;
	end int8; 
}

type conversions interface{
	add() int
	mult() int
}

func (twoInt*instance2)add() int {
	return 0; 
}

func (twoInt*instance2)mult() int {
	return 0; 
}


