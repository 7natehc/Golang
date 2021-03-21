package main

import (
	"applib"
	"applib/models"
	"fmt"
	"strconv"
)

func main() {

	//task3:
	repoMix, _ := CustomerInitAllLevels()
	repoMix.StoreTofile("data.json")					//uncomment necessary part in applib

	//task4:
	// repoSilver, _ := CustomerInitSilver()
	// repoSilver.StoreTofile("data.json")
	// LoadedRepo, _ := applib.LoadFromFile("data.json")
	// CustomerUpgrade(LoadedRepo)
	// LoadedRepo1, _ := applib.LoadFromFile("data.json")
	// fmt.Println("data===>",*LoadedRepo1)
}

//to upgrade customers to Gold and Platinum levels
func CustomerUpgrade(repo *applib.CustomerRepo) {
	for i := range repo.Customers {
		if (500 < repo.Customers[i].Points) && (repo.Customers[i].Points < 900) {
			repo.Customers[i].Level = 2
		}
		if repo.Customers[i].Points > 900 {
			repo.Customers[i].Level = 3
		}
	}
	repo.StoreTofile("data.json")
}

//to initialize silver level customers only
func CustomerInitSilver() (applib.CustomerRepo, error) {

	repo := applib.NewRepo()
	for i := 1; i < 16; i++ {
		cust, err := models.NewCustomer(("Cust" + strconv.Itoa(i)), (1500 + i*1), 70*i, 1)
		if err != nil {
			fmt.Println(err)
			return repo, err
		}
		repo.Add(*cust)
	}
	return repo, nil
}

//to initialize all level customers
func CustomerInitAllLevels() (applib.CustomerRepo, error) {

	repo := applib.NewRepo()
	for i := 1; i < 6; i++ {
		cust, err := models.NewCustomer(("Cust" + strconv.Itoa(i)), (1500 + i*1), (120 * i), models.Silver)
		if err != nil {
			fmt.Println(err)
			return repo, err
		}
		repo.Add(*cust)
	}
	for i := 1; i < 6; i++ {
		cust, err := models.NewCustomer(("Cust" + strconv.Itoa(i)), (1600 + i*1), (150 * i), models.Gold)
		if err != nil {
			fmt.Println(err)
			return repo, err
		}
		repo.Add(*cust)
	}
	for i := 1; i < 6; i++ {
		cust, err := models.NewCustomer(("Cust" + strconv.Itoa(i)), (1700 + i*1), (100 * i), models.Platinum)
		if err != nil {
			fmt.Println(err)
			return repo, err
		}
		repo.Add(*cust)
	}
	return repo, nil
}
