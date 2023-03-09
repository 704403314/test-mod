package main

import (
	"encoding/json"
	"fmt"
	"github.com/704403314/test-mod/function"
)

type User struct {
	Name   string `json:"name" validate:"required"`
	KycEmi string `json:"kycEmi"`
	Slice  []int64
}

func main() {
	strLen := function.GetStringLen("1234567")
	fmt.Println(strLen)
	var applicant User

	userJson := "{\n    \"name\":\"1239999\",\n    \"kycEmi\":\"2023-02-16 09:31:07\",\n    \"key\":\"VMMMJYVADBHYYT\"}"

	if err := json.Unmarshal([]byte(userJson), &applicant); err != nil {
		fmt.Println("err:", err.Error())
	}

	fmt.Println("-----applicant--------", applicant)
}
