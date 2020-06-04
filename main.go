package main

import "fmt"

type Company struct {
	Name string
	Id int
	Location string
	Domain string
}
func main() {
	c1:= Company{
		Name: "Kloudone",
		Id: 01,
		Location: "Chennai",
		Domain: "Cloud Services",
	}
	c2:= Company{
		Name: "TCS",
		Id: 02,
		Location: "Kolkata",
		Domain: "IT",
	}
	companyList:= []Company{c1,c2}
	fmt.Println(companyList)
	companyList = append(companyList, Company{"Wipro", 03, "Delhi", "Software"})
	fmt.Println(companyList)
	for i, v := range companyList{
		if i==1 {
			v.Name= "cognizant"

		}
		fmt.Println(i, v,)
	}

		companyList = append(companyList[:2], companyList[3:]...)
		fmt.Println(companyList)
}
