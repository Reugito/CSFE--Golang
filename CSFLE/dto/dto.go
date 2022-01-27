package dto

import (
	"encoding/json"
	"log"
	"net/http"
)

type EmployeeDetailsDTO struct {
	EmpId     string      `json:"emp_id" bson:"emp_id"`
	FirstName string      `json:"first_name" bson:"first_name"`
	LastName  string      `json:"last_name" bson:"last_name"`
	EmailId   string      `json:"email" bson:"email"`
	Password  string      `json:"password" bson:"password"`
	StartDate interface{} `json:"start_date" bson:"start_date"`
	EndDate   interface{} `json:"end_date" bson:"end_date"`
}

type ResponseDTO struct {
	ResponseCode int         `json:"response_code"`
	Status       string      `json:"status"`
	Response     interface{} `json:"response"`
}

type LoginResponseDTO struct {
	ResponseCode        int         `json:"response_code"`
	Status              string      `json:"status"`
	UserId              interface{} `json:"user_id"`
	Token               string      `json:"token"`
	SessionTimeINMinute interface{} `json:"sessionTimeINMinute"`
}

type TokenResponseDTO struct {
	Username       string      `json:"username"`
	Token          string      `json:"token"`
	ExpirationTime interface{} `json:"expirationTime"`
}

type SearchDTO struct {
	EmpId     string      `json:"emp_id" `
	FirstName string      `json:"first_name" `
	LastName  string      `json:"last_name" `
	EmailId   string      `json:"email_id" `
	Password  string      `json:"password"`
	StartDate string      `json:"start_date"`
	EndDate   string      `json:"end_date" `
	Page      interface{} `json:"page"`
}

type SortDTO struct {
	SortBy string      `json:"sort_by"`
	Order  string      `json:"order"`
	Page   interface{} `json:"page"`
}

type PageDetails struct {
	Data     interface{} `json:"data"`
	Page     interface{} `json:"current_page"`
	LastPage interface{} `json:"last_Page"`
	Total    interface{} `json:"total"`
}

func ResponseEntity(responseDT0 ResponseDTO, w http.ResponseWriter) {

	log.Println("/dto.ResponseEntity()")

	jsonOBJ, err := json.Marshal(responseDT0)
	if err != nil {

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Failed to convert the ResponseDTO"))
	} else {

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonOBJ)
	}

}

func LogInResponse(responseDT0 LoginResponseDTO, w http.ResponseWriter) {

	log.Println("/dto.Login Response()")

	jsonOBJ, err := json.Marshal(responseDT0)
	if err != nil {

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Failed to convert the ResponseDTO"))
	} else {

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonOBJ)
	}

}
