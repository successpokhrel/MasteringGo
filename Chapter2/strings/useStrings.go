package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	// comparing without considering their case
	f("EqualFold: %v\n", s.EqualFold("Success", "SUCCess"))
	f("EqualFold: %v\n", s.EqualFold("Success", "SUCCEs"))

	// check if string of second parameter is in string given at first parameter.
	// strings.Index() function if successful: returns index unsuccessful: -1
	f("Index: %v\v", s.Index("Success", "ce"))
	f("Index: %v\n", s.Index("Success", "Ce"))

	// HasPrefix() checks if second parameter is what first parameter starts with
	f("Prefix: %v\n", s.HasPrefix("Success", "Su"))
	f("Prefix: %v\n", s.HasPrefix("Success", "su"))
	f("Suffix: %v\n", s.HasSuffix("Success", "ss"))
	f("Suffix: %v\n", s.HasSuffix("Success", "SS"))

	// to split using white space characters
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t_new := s.Fields("ThisIs a string!")
	f("Fields: %v\n", len(t_new))

	// split and replace
	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	//split after keep the seperator too
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
