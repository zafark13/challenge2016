package utils

import (
	"RealImageSolution/models"
	"fmt"
	"strings"
	"unicode"
)

func GetMainMenu() {
	fmt.Println("                                      ")
	fmt.Println("1. Add Distributor with Permission")
	fmt.Println("2. List all Distributors")
	fmt.Println("3. Check Permission for a Distributor")
	fmt.Println("4. Add a Sub-Distributor with Permission")
	fmt.Println("5. Back to the Main Menu")
	fmt.Println("6. Exit")
	fmt.Println("                                      ")
}

// RemoveSpace removes the spaces from the string and returns the string in lower case
func RemoveSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}

	return strings.ToLower(string(rr))
}

// Contains checks if the given string is present in the given slice
func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}

	return false
}

func GetExcludesRelatedToTheRegion(Exclude []string, relatedTo string) []string {
	var list []string
	for _, val := range Exclude {
		if strings.Contains(val, relatedTo) {
			list = append(list, val)
		}
	}
	return list
}

func GetParentDistributor(distributors []models.Distributor) *models.Distributor {
	var parentName string
	for {
		fmt.Println("->Enter Parent Distributor Name: ")
		fmt.Scanln(&parentName)

		// Find the parent and child distributors in the list
		for _, dis := range distributors {
			if dis.Name == parentName {
				return &dis
			}
		}
		fmt.Printf("No distributor found with name %s. try again", parentName)
		fmt.Println("")
	}
}

func GetDistributorByName(name string, distributors []models.Distributor) *models.Distributor {
	for _, dist := range distributors {
		if dist.Name == name {
			return &dist
		}
	}
	return nil
}
