package main

import (
	"fmt"
	"log"

	"github.com/mytkdals93/headfirst_go/map/datafile"
)

func count(votes []string) {
	var names []string
	var countes []int

	for _, vote := range votes {
		matched := false
		for i, name := range names {
			if name == vote {
				countes[i]++
				matched = true
			}
		}
		if !matched {
			names = append(names, vote)
			countes = append(countes, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d \n", name, countes[i])
	}
}
func countWithMap(votes []string) {
	var names map[string]int
	names = make(map[string]int)

	for _, vote := range votes {
		names[vote]++
	}
	for name, count := range names {
		fmt.Printf("%s: %d \n", name, count)
	}

}

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	countWithMap(lines)
}
