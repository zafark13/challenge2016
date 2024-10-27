package main

import (
	"RealImageSolution/handler"
	"RealImageSolution/utils"
	"fmt"
	"os"
)

const (
	AddDistributor    = 1
	ListDistributors  = 2
	CheckPermission   = 3
	AddSubDistributor = 4
	BackToMainMenu    = 5
	ExitCLI           = 6
)

var (
	Distributor handler.DistributorInterface = &handler.DistributorsModel{}
)

func init() {
	fmt.Println("Loading cities from CSV file...")
	cityLoadStatus, err := Distributor.LoadCitiesFromCSV("./data/cities.csv")
	if err != nil {
		fmt.Println("Error in loading cities from CSV file")
		os.Exit(1)
	}

	if cityLoadStatus {
		fmt.Println("Cities loaded successfully")
	} else {
		fmt.Println("Error in loading cities from CSV file")
		os.Exit(1)
	}
}

func main() {

	fmt.Println("#############################|Real Image Challenge CLI TOOL|#############################")
	fmt.Println("#############################|Author: Zafar Khan|#############################")
	fmt.Println(" 						   ")
	var id int = 0

	for {

		fmt.Println("######## MAIN MENU ########")
		utils.GetMainMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case AddDistributor:
			// Adding a distributor to the list
			fmt.Println("")
			fmt.Println("#### ADDING A DISTRIBUTOR WITH PERMISSIONS ####")
			Distributor.AddDistributor(&id)
		case ListDistributors:
			// Printing the distributor from the list
			fmt.Println("")
			fmt.Println("#### PRINTING THE DISTRIBUTOR LIST ####")
			Distributor.ListDistributors()
		case CheckPermission:
			// Checking the distributor permissions
			fmt.Println("")
			fmt.Println("#### CHECKING THE DISTRIBUTOR PERMISSIONS ####")
			Distributor.CheckPermission()
		case AddSubDistributor:
			// Creating the network of distributors
			fmt.Println("")
			fmt.Println("#### Adding Sub-Distributor ####")
			Distributor.CreateSubDistributorNetwork(&id)
		case BackToMainMenu:
			// Get back to the main menu
			fmt.Println("")
			fmt.Println("#### GETTING BACK TO THE MAIN MENU ####")
			utils.GetMainMenu()
		case ExitCLI:
			fmt.Println("Exiting from the CLI Tool")
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
