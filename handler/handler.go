package handler

import (
	"RealImageSolution/models"
	"RealImageSolution/utils"
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

type DistributorInterface interface {
	LoadCitiesFromCSV(filename string) (bool, error)
	AddDistributor(id *int)
	ListDistributors()
	CheckPermission()
	SetPermission(sufix string, id int, permission string) (bool, error)
	VerifyQuery(query string) bool
	CreateSubDistributorNetwork(id *int)
}

type DistributorsModel struct {
	CountryStateMap    map[string][]string
	StateCityMap       map[string][]string
	CurrentDistributor models.Distributor
	Distributors       []models.Distributor
}

// LoadCitiesFromCSV loads the cities from the csv file
func (d *DistributorsModel) LoadCitiesFromCSV(filename string) (bool, error) {
	// Read the CSV file
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("Error while reading the csv file")
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	// skip the header row
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	d.CountryStateMap = make(map[string][]string)
	d.StateCityMap = make(map[string][]string)

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		country := strings.ToLower(utils.RemoveSpace(row[5]))
		province := strings.ToLower(utils.RemoveSpace(row[4]))
		city := strings.ToLower(utils.RemoveSpace(row[3]))

		if _, ok := d.CountryStateMap[country]; !ok {
			d.CountryStateMap[country] = make([]string, 0)
		}

		if !utils.Contains(d.CountryStateMap[country], province) {
			d.CountryStateMap[country] = append(d.CountryStateMap[country], province)
		}

		if _, ok := d.StateCityMap[province]; !ok {
			d.StateCityMap[province] = make([]string, 0)
		}

		if !utils.Contains(d.StateCityMap[province], city) {
			d.StateCityMap[province] = append(d.StateCityMap[province], city)
		}
	}

	return true, nil
}

// to add a distributor with permission
func (d *DistributorsModel) AddDistributor(id *int) {
	d.CurrentDistributor = models.Distributor{}
	*id++
	var name string
	fmt.Println("")
	fmt.Println("->Enter Distributor Name: ")
	fmt.Scanln(&name)
	d.CurrentDistributor.ID = *id
	d.CurrentDistributor.Name = name
	d.Distributors = append(d.Distributors, d.CurrentDistributor)
	fmt.Println("->Now Add Permissions for ", d.CurrentDistributor.Name)
	for {
		var permission string
		fmt.Println("Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		permission = scanner.Text()
		if permission == "4" {
			break
		}

		data := strings.Split(permission, ":")
		prefix := strings.TrimSpace(data[0])
		sufix := strings.TrimSpace(strings.ToLower(data[1]))

		if prefix == "INCLUDE" || prefix == "EXCLUDE" {
			ok, err := d.SetPermission(sufix, *id-1, prefix)
			if !ok {
				fmt.Printf("ERROR: %s", err.Error())
				println("")
			}
		} else {
			fmt.Println("Invalid Input")
		}
	}
}

// to list distributors
func (d *DistributorsModel) ListDistributors() {
	fmt.Printf("->Number of Distributor: %d", len(d.Distributors))
	fmt.Println("")
	fmt.Println("->Distributor List: ")
	for _, distributor := range d.Distributors {
		fmt.Println("Distributor ID :", distributor.ID, " Distributor Name : ", distributor.Name)
		fmt.Println("Permitted Places: ", strings.Join(distributor.Include, " "))
		fmt.Println("not Permitted Places: ", strings.Join(distributor.Exclude, " "))
		if distributor.SubDistributor {
			fmt.Println("Sub Distributor: YES")
		} else {
			fmt.Println("Sub Distributor: NO")
		}

		if distributor.Parent != "" {
			fmt.Println("Parent: " + distributor.Parent)
		} else {
			fmt.Println("Parent: NONE")
		}
		fmt.Println("")
	}
}

// to check if distributor have permission
func (d *DistributorsModel) CheckPermission() {
	for {
		fmt.Println("->Enter Distributor Name: or press 4 for Main menu")
		var name string
		fmt.Scanln(&name)

		if name == "4" {
			break
		}
		if d.CurrentDistributor.Name == "" {
			fmt.Println("Distributor is empty")
			return
		}
		dist := utils.GetDistributorByName(name, d.Distributors)
		if dist == nil {
			fmt.Println("Distributor not found")
			return
		}

		fmt.Println("->Enter your query to check permission: ")
		var query string
		fmt.Scanln(&query)

		ans := d.VerifyQuery(strings.ToLower(utils.RemoveSpace(query)))
		if ans {
			fmt.Println("")
			fmt.Println("YES")
			fmt.Println("")
		} else {
			fmt.Println("")
			fmt.Println("NO")
			fmt.Println("")
		}
	}
}

// helper function to verify if distributor have permission for given region
func (d *DistributorsModel) VerifyQuery(query string) bool {
	querySlice := strings.Split(query, "-")
	switch len(querySlice) {
	case 1:
		fmt.Println("Checking for CountryName")
		querySlice[0] = utils.RemoveSpace(querySlice[0])
		country := querySlice[0]
		for _, exclude := range d.CurrentDistributor.Exclude {
			exclude = utils.RemoveSpace(exclude)
			if country == exclude {
				return false
			}
		}
		for _, include := range d.CurrentDistributor.Include {
			include = utils.RemoveSpace(include)
			if country == include {
				return true
			}
		}

	case 2:
		fmt.Println("Checking for ProvinceName & CountryName")
		querySlice[0] = utils.RemoveSpace(querySlice[0])
		querySlice[1] = utils.RemoveSpace(querySlice[1])
		country := querySlice[1]
		stateCountry := querySlice[0] + "-" + querySlice[1]
		for _, exclude := range d.CurrentDistributor.Exclude {
			exclude = utils.RemoveSpace(exclude)
			if country == exclude || stateCountry == exclude {
				return false
			}
		}
		for _, include := range d.CurrentDistributor.Include {
			include = utils.RemoveSpace(include)
			if country == include || stateCountry == include {
				return true
			}
		}
	case 3:
		fmt.Println("Checking for CityName, ProvinceName & CountryName")
		querySlice[0] = utils.RemoveSpace(querySlice[0])
		querySlice[1] = utils.RemoveSpace(querySlice[1])
		querySlice[2] = utils.RemoveSpace(querySlice[2])
		country := querySlice[2]
		stateCountry := querySlice[1] + "-" + querySlice[2]
		cityStateCountry := querySlice[0] + "-" + querySlice[1] + "-" + querySlice[2]
		for _, exclude := range d.CurrentDistributor.Exclude {
			exclude = utils.RemoveSpace(exclude)
			if country == exclude || stateCountry == exclude || cityStateCountry == exclude {
				return false
			}
		}
		for _, include := range d.CurrentDistributor.Include {
			include = utils.RemoveSpace(include)
			if country == include || stateCountry == include || cityStateCountry == include {
				return true
			}
		}
	}

	return false
}

// for setting include and exclude permission to a distributor
func (d *DistributorsModel) SetPermission(region string, id int, permission string) (bool, error) {
	regionSlice := strings.Split(region, "-")
	for _, val := range regionSlice {
		val = utils.RemoveSpace(val)
	}

	switch len(regionSlice) {
	case 1:
		{
			if _, ok := d.CountryStateMap[regionSlice[0]]; !ok {
				return false, errors.New("region not in csv")
			}
		}
	case 2:
		{
			if _, ok := d.CountryStateMap[regionSlice[1]]; !ok {
				return false, errors.New("region not in csv")
			}

			if _, ok := d.StateCityMap[regionSlice[0]]; !ok {
				return false, errors.New("region not in csv")
			}

		}
	case 3:
		{
			if _, ok := d.CountryStateMap[regionSlice[2]]; !ok {
				return false, errors.New("region not in csv")
			}

			if cities, ok := d.StateCityMap[regionSlice[1]]; !ok {
				return false, errors.New("region not in csv")
			} else {
				if !strings.Contains(strings.Join(cities, "-"), regionSlice[0]) {
					return false, errors.New("region not in csv")
				}
			}
		}
	default:
		fmt.Println("Invalid Input, Please try again")
	}

	if permission == "INCLUDE" {
		d.CurrentDistributor.Include = append(d.CurrentDistributor.Include, region)
	}
	if permission == "EXCLUDE" {
		d.CurrentDistributor.Exclude = append(d.CurrentDistributor.Exclude, region)
	}

	d.Distributors[id] = d.CurrentDistributor
	return true, nil
}

// for creating a sub distributor
func (d *DistributorsModel) CreateSubDistributorNetwork(id *int) {
	d.CurrentDistributor = models.Distributor{}
	*id++
	var name string
	var parentDistributor *models.Distributor
	fmt.Println("")
	fmt.Println("->Enter Distributor Name: ")
	fmt.Scanln(&name)
	d.CurrentDistributor.ID = *id
	d.CurrentDistributor.Name = name
	d.CurrentDistributor.SubDistributor = true
	parentDistributor = utils.GetParentDistributor(d.Distributors)
	if parentDistributor != nil {
		d.CurrentDistributor.Parent = parentDistributor.Name
	}
	d.Distributors = append(d.Distributors, d.CurrentDistributor)
	fmt.Println("->Now Add Permissions for ", d.CurrentDistributor.Name)
	for {
		var permission string
		fmt.Println("Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		permission = scanner.Text()
		if permission == "4" {
			break
		}

		data := strings.Split(permission, ":")
		prefix := strings.TrimSpace(data[0])
		sufix := strings.TrimSpace(strings.ToLower(data[1]))

		switch prefix {
		case "INCLUDE":
			child := d.CurrentDistributor
			d.CurrentDistributor = *parentDistributor
			havePermission := d.VerifyQuery(sufix)
			d.CurrentDistributor = child
			if havePermission {
				ok, err := d.SetPermission(sufix, *id-1, prefix)
				if !ok {
					fmt.Printf("ERROR: %s", err.Error())
					println("")
				} else {
					// adding excludes related to the included region from parent
					d.CurrentDistributor.Exclude = append(d.CurrentDistributor.Exclude, utils.GetExcludesRelatedToTheRegion(parentDistributor.Exclude, sufix)...)
					d.Distributors[*id-1] = d.CurrentDistributor
				}
			} else {
				fmt.Println("Parent distributor does to have rights for this region. please try again")
				continue
			}
		case "EXCLUDE":
			ok, err := d.SetPermission(sufix, *id-1, prefix)
			if !ok {
				fmt.Printf("ERROR: %s", err.Error())
				println("")
			}
		default:
			fmt.Println("Invalid Choice, Try Again!")
		}
	}
}
