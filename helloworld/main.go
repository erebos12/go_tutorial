/*
1. What is a package: project name == workspace

2. Two types of packages:
 - Executable package which you run like 'go run <package-binary>'
   (main package must have main function)
 - Reusable package which you use as dependency or library
*/
package main

/*
  - import another package (reusable package)
*/
import "fmt"

func main() {
	fmt.Println("Hello World")
}
