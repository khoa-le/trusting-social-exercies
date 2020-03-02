package main

import (
	"fmt"
	"github.com/khoa-le/trusting-social-exercies/file"
	"log"
	"time"
)

type ActivationSession struct {
	ActivationDate   string
	DeActivationDate string
}
type Phone struct {
	PhoneNumber       string
	ActivationSession ActivationSession
}

func main() {
	inputFileName := "data.csv"
	csvData := file.OpenAndReadFile(inputFileName)
	phoneData := buildMappingData(csvData)

	csvRows := buildDataToWriteFile(phoneData)
	outputFileName := "real_activation_date.csv"
	err := file.WriteFile(outputFileName, csvRows)
	if err != nil {
		fmt.Printf("Can not write file with error %v", err)
	} else {
		fmt.Printf("Data write success to file: %v", outputFileName)
	}

}

func buildMappingData(csvFileData [][]string) map[string]Phone {
	var result = make(map[string]Phone)
	for i, row := range csvFileData {
		if i == 0 {
			continue
		}
		phoneNumber := row[0]
		if result[phoneNumber].PhoneNumber == "" {
			activationSession := ActivationSession{ActivationDate: row[1], DeActivationDate: row[2]}
			phone := Phone{PhoneNumber: phoneNumber, ActivationSession: activationSession}
			result[phoneNumber] = phone
		} else {
			result[phoneNumber] = mergeActivationSessionToPhone(result[phoneNumber], phoneNumber, row[1], row[2])
		}
	}
	return result
}

func mergeActivationSessionToPhone(phone Phone, phoneNumber string, activationDate string, deActivationDate string) Phone {
	if phone.PhoneNumber != phoneNumber {
		log.Fatal("Phone Number was difference, can not merge")
	}
	phone.ActivationSession = calculateActivationDateSession(phone.ActivationSession, activationDate, deActivationDate)
	return phone
}

func calculateActivationDateSession(currentActivationDate ActivationSession, dateStart string, dateEnd string) ActivationSession {
	currentActivation, _ := time.Parse("2006-01-02", currentActivationDate.ActivationDate)
	currentDeActivation, _ := time.Parse("2006-01-02", currentActivationDate.DeActivationDate)

	newActivationDate, _ := time.Parse("2006-01-02", dateStart)
	newDeActivationDate, _ := time.Parse("2006-01-02", dateEnd)

	if currentActivation.Before(newActivationDate) {
		if currentDeActivation.Equal(newActivationDate) {
			return ActivationSession{ActivationDate: currentActivationDate.ActivationDate, DeActivationDate: dateEnd}
		} else {
			currentActivationDate = ActivationSession{ActivationDate: dateStart, DeActivationDate: dateEnd}
		}
	} else {
		if currentActivation.Equal(newDeActivationDate) {
			currentActivationDate = ActivationSession{ActivationDate: dateStart, DeActivationDate: currentActivationDate.DeActivationDate}
		}
	}
	return currentActivationDate
}

func buildDataToWriteFile(phoneData map[string]Phone) [][]string {
	var result [][]string
	//Create header
	result = append(result, []string{"PHONE_NUMBER", "REAL_ACTIVATION_DATE"})
	for _, phone := range phoneData {
		item := []string{phone.PhoneNumber, phone.ActivationSession.ActivationDate}
		result = append(result, item)
	}
	return result
}
