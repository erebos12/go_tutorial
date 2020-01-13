# Go - Idioms & Best Practices

Not an OO language !!!

## From OO Aproach to Go Approach

type <name_of_type> <datatype>

## Receiver Functions

## Unused variables with "_"

## Slices

sliceName[startIndexIncluding:upToNotIncludingIndex]

i.e. slice[0:2] - give me index[0] and index[1] BUT NOT index[2]
=> 0 can be left out since Go would automatically assume the 0 index
i.e. slice[:2]
Same goes for end index!

## Multiple Return Values