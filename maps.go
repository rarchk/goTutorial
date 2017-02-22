package main
import(
	"fmt"
	)

func main() {
	
	/* Creating a map with string key and int value */
	fooMap := make(map[string]int);
	fooMap["christmas"] = 25;
	fooMap["Halloween"] = 31;
	fmt.Println(fooMap);
	delete(fooMap,"Halloween");
	}

