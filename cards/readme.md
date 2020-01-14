# Go - Idioms & Best Practices

* Not an OO language !!!

* Statically typed language

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