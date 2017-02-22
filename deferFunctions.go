package main
import(
	"fmt"
	"os"
 )

func createFile(filename string)*os.File{
	fmt.Println("Creating a file");
	file, err := os.Create(filename);
	if (err != nil){
		panic(err);
	}
	return file;
}

func writeFile(filename *os.File, input string){
	fmt.Println("Writing to the file");
	fmt.Fprintln(filename,input);
}

func closeFile(filename *os.File){
	fmt.Println("Closing the file")
	filename.Close();
}

func fileHandler(filename string, input string){
	file := createFile(filename);
	defer closeFile(file); // in case anything wrong happens, defer closes file 
	writeFile(file, input);
}

func main() {
	fileHandler("/tmp.txt","This is real world");
	}

