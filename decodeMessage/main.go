/*
Message has been encoded in a text stream read char by char
message is a comma-delimited int
  - each posative num can be represented by a c++ int
  - depends on encoding
    - uppercase, lowercase, punctuation
  - uppercase mode:
    int % 27 == number where 1 = A
  - lowercase mode:
    int % 27 == number where 1 = a
  - punctuation mode:
    int % 9 = punc where 1 = '!'

at the beginning of each message the mode is Uppercase
each time number % == 0 the encoding changes
  - if mode is uppercase next is lower
  - if mode is lowercase next is punctuation
  - if mode is punctuation next is uppercase


Plan:
 - handle as an array, loop over
   check modulo
      ? how to map char to modulo
   if modulo == 0 change state

punctuation table:
1: !
2: ?
3: ,
4: .
5: space
6: ;
7: "
8: '

sample input:
in: 18,12312,171,763,98423,1208,216,11,500,18,241,0,32,20620,27,10
out: Right? Yes!
*/

package main

import (
	"fmt"
	//"log"
	//"errors"
)

var punctuation []string = []string{"!", "?", ",", ".", " ", ";", "\"", "`"}

func changeEncoding(current string) string {
	if current == "upper" {
		return "lower"
	} else if current == "lower" {
		return "punctuation"
	} else if current == "punctuation" {
		return "upper"
	} else {
		return "NOPE"
	}

}

func upperChar(i int, encoding string) rune {
	if encoding == "upper" {
		return rune('A' - 1 + i)
	} else if encoding == "lower" {
		return rune('a' - 1 + i)
	} else {
		return rune('!' - 1 + i)
	}

}

func main() {
	message := []int{18, 12312, 171, 763, 98423, 1208, 216, 11, 500, 18, 241, 0, 32, 20620, 27, 10}
  decodedMessage := ""
	encoding := "upper"
	for _, i := range message {
		if encoding == "upper" || encoding == "lower" {
			if i%27 == 0 {
				encoding = changeEncoding(encoding)
			} else {
        decodedMessage += string(upperChar(i%27, encoding))
				//fmt.Println(string(upperChar(i%27, encoding)))
			}
		} else {
			if i%9 == 0 {
				encoding = changeEncoding(encoding)
			} else {
				//fmt.Println(i)
				//fmt.Println(string(upperChar(i%9, encoding)))
        decodedMessage += punctuation[i%9-1]
			}
		}
	}
  fmt.Println(decodedMessage)
}
