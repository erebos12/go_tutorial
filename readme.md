# Go - Idioms & Best Practices

* Not an OO language !!!

* Statically typed language

* Paas-by-Value language

Package Documentation https://golang.org/pkg/

## From OO Aproach to Go Approach

type <name_of_type> <datatype>

## Receiver Functions

Receiver vs. Argument !

## Unused variables with "_"

## Slices

sliceName[startIndexIncluding:upToNotIncludingIndex]

i.e. slice[0:2] - give me index[0] and index[1] BUT NOT index[2]
=> 0 can be left out since Go would automatically assume the 0 index
i.e. slice[:2]
Same goes for end index!

## Multiple Return Values

## Public / Private Functions

## Type Conversion
Many interfaces in Go work with slice-of-bytes ([]byte). So we need to convert
our specific values:

```
greeting := "Hello World!"
[]byte(greeting)
```

## Error Handling

Common Go pattern is:

```
result, error := someFunctionCanThrowAnError()
if error != nil {
// logit or panic or exit or whatever
}
```

## Pointers and * & Operators

A pointer holds the memory address of a value.
In Go a pointer is represented using the `*` (asterisk) character followed by the type of the stored value.
`*` is also used to "dereference" pointer variables. Dereferencing a pointer gives us access to the value the pointer points to.

The `&` operator generates a pointer to its operand.

1. To turn an address into a value then use `*address`
2. To turn a value into an address then use `&value`

### Question / Answer about Pointers

**Question**: Whenever you have a pointer in the receiver of a function (see snippet), what does that say to us?
```
func (p *person) updateName(newFirstName string) {
```
> **Answer**: It specifies the type of the receiver that the function expetcs. It is NOT turning the pointer address into a value (no dereferencing!!!).


**Question**: Whenever you pass an integer, float, string, or struct into a function, what does Go do with that argument?
> **Answer**: It creates a copy of each argument, and these copies are used inside of the function.

**Question**: What is the &  operator used for?
> **Answer**: Turning a value into a pointer.

**Question**: When you see a `*` operator in front of a pointer, what will it turn the pointer into?
> **Answer**: A value.

**Question**: Take a look at the following program.  The changeLatitude function expects a receiver of type pointer to a location struct , but in the main function the receiver is a value type of a struct.  What will happen when this code is executed?

```
package main
 
import "fmt"
 
type location struct {
 longitude float64
 latitude float64
}
 
func main() {
 newYork := location{
   latitude: 40.73,
   longitude: -73.93,
 }
 
 newYork.changeLatitude()
 
 fmt.Println(newYork)
}
 
func (lo *location) changeLatitude() {
 (*lo).latitude = 41.0
}

```
> **Answer**: This program uses a shortcut, where Go will automatically assume even though ``` newYork.changeLatitude() ``` is using a value type we probably meant to pass in a pointer to the ```newYork``` struct.


**Question**: Here's a tricky one!  What will the following program print out?

```
package main
 
import "fmt"
 
func main() {
    name := "Bill"
    fmt.Println(*&name)
}

```
> **Answer**: The string "Bill".

**Question** What will the main() function will print out?
```
package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Hi", "There"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string){
    s[0] = "Bye"
}
```

> **Answer**: "Bye There".

WHY IS THE SLICE NOT COPIED HERE?

It's not like passing a `struct` which would be copied. slices (and other types) are special kind of types and are handled differently in memory by Go.

A ```slice``` in Go consists of a pointer to the head of an array, the capacity of the array and the length of the array. These metadata of the slice are stored in a separate register address. So when we call a function with a slice then Go creates a copy (call-by-value) of the slice metadata in memory, not a copy of the referenced array. See image below:

<table><tr><td>
<img align="center" src="./pics/slice_mem.png" title="Passing a slice to a function" width="500">
</td></tr></table>

In Go we refer to these kind of types as "reference types" which all have this behaviour.

<table><tr><td>
<img align="center" src="./pics/reference_types.png" title="Passing a slice to a function" width="500">
</td></tr></table>