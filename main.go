package main

import (
	"fmt"
	"strings"
)

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func FindAncestors(relations map[string][]string, child2 string, child1 string) string {
	fmt.Println("Starting....")
	found := true
	returned, child1Ancestory, child2Ancestory := "", "", ""
	for found == true {
		found = false
		for key, value := range relations {
			if contains(value, child1) == true && contains(value, child1) == true {

				return key
			} else if contains(value, child1) == true {
				if strings.Contains(child2Ancestory, key) {
					return key
				}
				child1Ancestory += key
				child1 = key
				found = true
			} else if contains(value, child2) == true {
				if strings.Contains(child1Ancestory, key) {
					return key
				}
				child2Ancestory += key
				child2 = key
				found = true
			}
		}
	}
	return returned
}

func main() {
}
