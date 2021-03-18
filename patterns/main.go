/*
Print:

*    *
**  **
******
**  **
*    *

constraints:
1) only use 1 print statemet. 
2) arbitrary length

step:
print first + last positon
print first + second until sides meet then reverse

plan:
get lenght
make an empty array of length length
print first + last index of "*"
*/
package main

import (
        "fmt"
)

func pattern(length int) {
        arr := make([]string, length)
        for i := 0; i < length / 2; i++ {
                arr[i] = "*"
                arr[length - (1 + i)] = "*"
                fmt.Println(arr)
                if length / 2 == i + 1 {
                        for j := i; j > 0; j-- {
                                arr[j] = ""
                                arr[length - j - 1] = ""
                                fmt.Println(arr)
                        }
                }
        }
}

func main() {
        // even
        pattern(6)
        // odd
        pattern(7)
}
