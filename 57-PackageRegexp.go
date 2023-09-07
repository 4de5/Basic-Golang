/*
Regexp = regular Expresion

dibuat dengan bahasa C oleh Google dan diberi nama RE2

https:/github.com/google/re2/wiki/Syntax
https:/golang.org/pkg/regexp
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z]o)")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("elo"))
	fmt.Println(regex.MatchString("eDo")) // Huruf besar

	cari := regex.FindAllString("eko edo eli eka eke budi iko ewo elo", -1)
	fmt.Println(cari)

}
