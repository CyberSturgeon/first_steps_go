package main

import (
	"first_steps_go/internal/homework"
	"fmt"
)

var wordCountTestCases = []string{
	"ya ebu ebu sobak",
	"  a a a  ",
	"",
	"  ",
}

var areAnagramsTestCases = [][2]string{
	{"aboba", "aboba"},
	{"y a g", "a g y"},
	{"", "a"},
	{"a", "  "},
}

var firstUniqueTestCases = []string{
	"a b c a b c",
	"a b b",
	"",
	"  ",
}

func main() {
	fmt.Println(1)
	for _, tc := range wordCountTestCases {
		fmt.Printf("input %v\n", tc)
		res, err := homework.WordCount(tc)
		fmt.Printf("result=%v error=%v\n", res, err)
	}
	fmt.Println(2)
	for _, tc := range areAnagramsTestCases {
		fmt.Printf("input %v | %v\n", tc[0], tc[1])
		res, err := homework.AreAnagrams(tc[0], tc[1])
		fmt.Printf("result=%v error=%v\n", res, err)
	}
	fmt.Println(3)
	for _, tc := range firstUniqueTestCases {
		fmt.Printf("input %v\n", tc)
		res, err := homework.FirstUnique(tc)
		fmt.Printf("result=%v error=%v\n", res, err)
	}
}
