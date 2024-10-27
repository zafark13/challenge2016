# REAL IMAGE SOLUTION Golang | Zafar Khan  (CLI Application)

#### Run the following command to execute the CLI application
`go run main.go`

#### Just Start the application and follow the instructions, thats all you need to do.

#### For more refer the output.txt file for the output of the CLI application

### NOTE:
1. If input entered like `xyz` program will consider `xyz` as country
2. If input entered like `uvw-xyz` program will consider `xyz` as country and `uvw` as state
3. If input entered like `abc-uvw-xyz` program will consider `xyz` as country, `uvw` as state annd `abc` as city

#### Sample Input and Output
```md
#############################|Real Image Challenge CLI TOOL|#############################
#############################|Author: Zafar Khan|#############################

######## MAIN MENU ########

1. Add Distributor with Permission
2. List all Distributors
3. Check Permission for a Distributor
4. Add a Sub-Distributor with Permission
5. Back to the Main Menu
6. Exit

1

#### ADDING A DISTRIBUTOR WITH PERMISSIONS ####

->Enter Distributor Name:
DIST1
->Now Add Permissions for  DIST1
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
INCLUDE: INDIA
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
INCLUDE: UNITEDSTATES
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
EXCLUDE: KARNATAKA-INDIA
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
EXCLUDE: CHENNAI-TAMILNADU-INDIA
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
4
######## MAIN MENU ########

1. Add Distributor with Permission
2. List all Distributors
3. Check Permission for a Distributor
4. Add a Sub-Distributor with Permission
5. Back to the Main Menu
6. Exit

3

#### CHECKING THE DISTRIBUTOR PERMISSIONS ####
->Enter Distributor Name: or press 4 for Main menu
DIST1
->Enter your query to check permission: 
CHICAGO-ILLINOIS-UNITEDSTATES
Checking for CityName, ProvinceName & CountryName

YES

->Enter Distributor Name: or press 4 for Main menu
DIST1
->Enter your query to check permission: 
CHENNAI-TAMILNADU-INDIA
Checking for CityName, ProvinceName & CountryName

NO

->Enter Distributor Name: or press 4 for Main menu
DIST1
->Enter your query to check permission: 
BANGALORE-KARNATAKA-INDIA
Checking for CityName, ProvinceName & CountryName

NO

->Enter Distributor Name: or press 4 for Main menu
4
######## MAIN MENU ########

1. Add Distributor with Permission
2. List all Distributors
3. Check Permission for a Distributor
4. Add a Sub-Distributor with Permission
5. Back to the Main Menu
6. Exit

4

#### Adding Sub-Distributor ####

->Enter Distributor Name:
DIST3
->Enter Parent Distributor Name: 
DIST 
No distributor found with name DIST. try again
->Enter Parent Distributor Name:
DIST1
->Now Add Permissions for  DIST3
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
INCLUDE: INDIA
Checking for CountryName
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
EXCLUDE: TAMILNADU-INDIA
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
4
######## MAIN MENU ########

1. Add Distributor with Permission
2. List all Distributors
3. Check Permission for a Distributor
4. Add a Sub-Distributor with Permission
5. Back to the Main Menu
6. Exit

2

#### PRINTING THE DISTRIBUTOR LIST ####
->Number of Distributor: 2
->Distributor List:
Distributor ID : 1  Distributor Name :  DIST1
Permitted Places:  india unitedstates
not Permitted Places:  karnataka-india chennai-tamilnadu-india    
Sub Distributor: NO
Parent: NONE

Distributor ID : 2  Distributor Name :  DIST3
Permitted Places:  india
not Permitted Places:  karnataka-india chennai-tamilnadu-india tamilnadu-india
Sub Distributor: YES
Parent: NONE

######## MAIN MENU ########

1. Add Distributor with Permission
2. List all Distributors
3. Check Permission for a Distributor
4. Add a Sub-Distributor with Permission
5. Back to the Main Menu
6. Exit

3

#### CHECKING THE DISTRIBUTOR PERMISSIONS ####
->Enter Distributor Name: or press 4 for Main menu
DIST3
->Enter your query to check permission: 
HARYANA-INDIA
Checking for ProvinceName & CountryName

YES

->Enter Distributor Name: or press 4 for Main menu
DIST3
->Enter your query to check permission: 
CHINA
Checking for CountryName

NO

->Enter Distributor Name: or press 4 for Main menu
4
######## MAIN MENU ########

1. Add Distributor with Permission
2. List all Distributors
3. Check Permission for a Distributor
4. Add a Sub-Distributor with Permission
5. Back to the Main Menu
6. Exit

4

#### Adding Sub-Distributor ####

->Enter Distributor Name:
DIST5
->Enter Parent Distributor Name:
DIST3
->Now Add Permissions for  DIST5
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA
INCLUDE: HUBLI-KARNATAKA-INDIA
Checking for CityName, ProvinceName & CountryName
Parent distributor does to have rights for this region. please try again
Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu | Ex: INCLUDE: INDIA or EXCLUDE: KARNATAKA-INDIA


```