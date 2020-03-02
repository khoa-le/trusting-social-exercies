package main

import (
	"testing"
)

func TestGetRealActivationDate(t *testing.T) {
	phoneData := make(map[string]Phone)
	activationSession := ActivationSession{ActivationDate: "2016-03-01", DeActivationDate: "2016-05-01"}
	phoneObject := Phone{PhoneNumber: "0978181187", ActivationSession: activationSession}
	phoneData[phoneObject.PhoneNumber] = phoneObject

	testingPhone := "00000000"
	realActivationNotExist := getRealActivationDate(phoneData, testingPhone)
	if realActivationNotExist != "" {
		t.Errorf("The Phone did not exist expect emtpty string, but result is %s ", realActivationNotExist)
	}

	realActivationExist := getRealActivationDate(phoneData, "0978181187")
	expectRealActivationDate := "2016-03-01"
	if realActivationExist != expectRealActivationDate {
		t.Errorf("The Phone did expect exist expect %s, but result is %s ", expectRealActivationDate, realActivationExist)
	}
}

func TestMergeActivationSessionToPhone(t *testing.T) {
	activationSession := ActivationSession{ActivationDate: "2016-03-01", DeActivationDate: "2016-05-01"}
	phone := Phone{PhoneNumber: "0978181187", ActivationSession: activationSession}

	newPhone := mergeActivationSessionToPhone(phone, "0978181187", "2016-01-01", "2016-03-01")
	if newPhone.PhoneNumber != phone.PhoneNumber {
		t.Error("After merge phone number should be the same")
	}

	if newPhone.ActivationSession.ActivationDate != "2016-01-01" {
		t.Errorf("Expect ActivationDate is 2016-01-01 but result is %v", newPhone.ActivationSession.ActivationDate)
	}

	newPhone = mergeActivationSessionToPhone(phone, "0978181187", "2016-12-01", "")
	if newPhone.ActivationSession.ActivationDate != "2016-12-01" {
		t.Errorf("Expect ActivationDate is 2016-12-01 but result is %v", newPhone.ActivationSession.ActivationDate)
	}

	newPhone1 := mergeActivationSessionToPhone(phone, "0978181187", "2016-07-01", "2016-09-01")
	if newPhone1.ActivationSession.ActivationDate != "2016-07-01" {
		t.Errorf("Expect ActivationDate is 2016-07-01 but result is %v", newPhone1.ActivationSession.ActivationDate)
	}

	activationSession1 := ActivationSession{ActivationDate: "2016-02-01", DeActivationDate: "2016-03-01"}
	phone1 := Phone{PhoneNumber: "0978181187", ActivationSession: activationSession1}

	newPhone3 := mergeActivationSessionToPhone(phone1, "0978181187", "2016-03-01", "2016-05-01")
	if newPhone3.ActivationSession.ActivationDate != "2016-02-01" {
		t.Errorf("Expect ActivationDate is 2016-02-01 but result is %v", newPhone3.ActivationSession.ActivationDate)
	}
	newPhone4 := mergeActivationSessionToPhone(newPhone3, "0978181187", "2016-05-01", "")
	if newPhone4.ActivationSession.ActivationDate != "2016-02-01" {
		t.Errorf("Expect ActivationDate is 2016-02-01 but result is %v", newPhone4.ActivationSession.ActivationDate)
	}
}
