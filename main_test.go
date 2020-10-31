package main

import (
	"fmt"
	"testing"
)

//TestFindAncestors test
func TestFindAncestors(t *testing.T) {
	// <setup code>
	testCases := []struct {
		code      string
		name      string
		relations map[string][]string
		childA    string
		childB    string
		ancestor  string
	}{
		//SETTING UP THE TEST FAMILY TREES
		//SIMPLE FAMILY WITH 3 PEOPLE, A = PARENT OF B,C
		{"IMPLE FAMILY WITH 3 PEOPLE, A = PARENT OF B,C",
			"Test1", map[string][]string{
				"A": []string{"B", "C"},
				"B": []string{},
				"C": []string{},
			}, "B", "C", "A"},
		//LARGER FAMILY
		{"LARGER FAMILY", "Test2", map[string][]string{
			"A": []string{"B", "C"},
			"H": []string{},
			"C": []string{"E", "F"},
			"E": []string{"G"},
			"G": []string{},
			"F": []string{},
		}, "G", "F", "C"},
		//3 PEOPLE WITHOUT RELATION
		{"3 PEOPLE WITHOUT RELATION", "Test3", map[string][]string{
			"A": []string{},
			"B": []string{},
			"C": []string{},
		}, "B", "C", ""},
		//VERY LOB SIDED TREE
		{"VERY LOB SIDED TREE", "Test4", map[string][]string{
			"A": []string{"B", "C"},
			"B": []string{},
			"C": []string{"D"},
			"D": []string{"E"},
			"E": []string{"F"},
			"F": []string{"G"},
			"G": []string{"H"},
			"H": []string{"I"},
			"I": []string{},
		}, "I", "B", "A"},
	}

	for i, tc := range testCases {
		fmt.Printf("Test %d \nTest is : %s Children : %s %s | Parent : %s ", i, tc.code, tc.childA, tc.childB, tc.ancestor)
		t.Run(tc.name, func(t *testing.T) {
			parent := FindAncestors(tc.relations, tc.childA, tc.childB)
			if parent != tc.ancestor {
				t.Fatal("incorrect ancestor found")
			}
		})
	}
}
