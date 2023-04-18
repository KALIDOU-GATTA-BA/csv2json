package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	EmpId string
	NamePrefix  string
	FirstName       string
	MiddleInitial  string
	LastName string
	Gender string
	Email string
	FatherName string
	MotherName string
	MotherMaidenName string
	DateOfBirth string
	TimeOfBirth string
	WeightInKg string
	DateofJoining string
	QuarterOfJoining string
	HalfOfJoining string
	YearOfJoining string
	MonthOfJoining string
	MonthNameOfJoining string
	ShortMonth string
	DayOfJoining string
	DOWOfJoining string
	ShortDOW string
	AgeInCompany string
	Salary string
	LastPercentHike string
	SSN string
	PhoneNumber string
	Country string
	City string
	State string
	Zip string
	Region string
	UserName string
	Password string
}

func main() {
	csvFile, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var rec Record
	var records [] Record

	for _, each := range csvData {
		rec.EmpId = each[0]
		rec.NamePrefix = each[1]
		rec.FirstName = each[2]
		rec.MiddleInitial  = each[3]
		rec.LastName  = each[4]
		rec.Gender  = each[5]
		rec.Email  = each[5]
		rec.FatherName  = each[6]
		rec.MotherName  = each[7]
		rec.MotherMaidenName  = each[8]
		rec.DateOfBirth  = each[9]
		rec.TimeOfBirth  = each[10]
		rec.WeightInKg  = each[11]
		rec.DateofJoining  = each[12]
		rec.QuarterOfJoining  = each[13]
		rec.HalfOfJoining  = each[14]
		rec.YearOfJoining  = each[15]
		rec.MonthOfJoining  = each[16]
		rec.MonthNameOfJoining  = each[17]
		rec.ShortMonth  = each[18]
		rec.DayOfJoining  = each[19]
		rec.DOWOfJoining  = each[20]
		rec.ShortDOW  = each[21]
		rec.AgeInCompany  = each[22]
		rec.Salary  = each[23]
		rec.LastPercentHike  = each[24]
		rec.SSN  = each[25]
		rec.PhoneNumber  = each[26]
		rec.Country  = each[27]
		rec.City  = each[28]
		rec.State  = each[29]
		rec.Zip  = each[30]
		rec.Region  = each[31]
		rec.UserName  = each[32]
		rec.Password  = each[33]
		records = append(records, rec)
	}

	jsonData, err := json.MarshalIndent(records, "", "    ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("Conversion complete")

}
