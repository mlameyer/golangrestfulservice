package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var function string
var user string
var password string
var id int
var carriername string
var address string
var active string
var token string

type Auth struct {
	Token string `json:"token"`
}

type Carrier struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewCarrier struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Active  bool   `json:"active"`
}

type UpdateAddress struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}

type UpdateActiveStatus struct {
	Id     int  `json:"id"`
	Active bool `json:"active"`
}

type Response struct {
	Carrier Carrier `json:"carrier"`
	Message string  `json:"message"`
	Success bool    `json:"success"`
}

type Message struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}

type ErrorResponse struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"status_code"`
	Message    Message `json:"message"`
}

func init() {
	flag.StringVar(&function, "f", "", "function")
	flag.StringVar(&user, "u", "", "user")
	flag.StringVar(&password, "p", "", "password")
	flag.IntVar(&id, "i", 0, "id")
	flag.StringVar(&carriername, "cn", "", "carriername")
	flag.StringVar(&address, "add", "", "address")
	flag.StringVar(&active, "act", "", "active")
	flag.StringVar(&token, "t", "", "token")

	flag.Parse()
}

func main() {
	deleteCarrier(6, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQ2Mjk4NTcsIm5hbWUiOiJKYWNrIEJ1cnRvbiJ9.cRFZFvGn4ezhTvsPJfjR782Ox99oQT60NnN818e02OA")
	switch functionToExecute := function; functionToExecute {
	case "authenticate":
		if user == "" || password == "" {
			panic(errors.New("Must provide user name and password... hint: user=jack password=burton"))
		}
		getToken(user, password)
	case "get":
		if token == "" {
			panic(errors.New("Must provide token... hint: did you authenticate?"))
		}
		getCarrier(id, token)
	case "create":
		if token == "" {
			panic(errors.New("Must provide token... hint: did you authenticate?"))
		}

		var activeStatus bool
		var err error
		if active != "" {
			activeStatus, err = strconv.ParseBool(active)
			if err != nil {
				panic(err)
			}
		}
		createCarrier(carriername, address, activeStatus, token)
	case "update":
		if token == "" {
			panic(errors.New("Must provide token... hint: did you authenticate?"))
		}
		if address == "" && active == "" {
			panic(errors.New("Must provide at least one for address or active status for update"))
		}
		if len(address) > 0 && len(active) > 0 {
			panic(errors.New("Must only provide address or active status for update. Not both"))
		}

		var activeStatus bool
		var err error
		if active != "" {
			activeStatus, err = strconv.ParseBool(active)
			if err != nil {
				panic(err)
			}
		}
		updateCarrier(id, address, activeStatus, token)
	case "delete":
		if token == "" {
			panic(errors.New("Must provide token... hint: did you authenticate?"))
		}

		deleteCarrier(id, token)
	default:
		panic(errors.New("Must provide a function to execute. choices: authenticate, get, create, update, delete"))
	}
}

func deleteCarrier(i int, t string) {
	url := "http://localhost:8080/api/carriers/" + strconv.Itoa(i)
	bearer := "Bearer " + t

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {

		var message Message
		err = json.Unmarshal(responseBody, &message)
		var e = ErrorResponse{
			Message:    message,
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}

		fmt.Printf("ErrorResponse\nMessage{ error: %v, msg: %s},\nStatus: %s,\nStatusCode: %d", e.Message.Error, e.Message.Msg, e.Status, e.StatusCode)
	} else {
		var response Response
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Message: %s\nSuccess: %v\n", response.Message, response.Success)
	}
}

func updateCarrier(id int, address string, active bool, token string) {
	url := "http://localhost:8080/api/carriers/" + strconv.Itoa(id)
	bearer := "Bearer " + token

	var body []byte
	var err error
	if len(address) > 0 {
		url += "/address"
		body, err = json.Marshal(UpdateAddress{
			Id:      id,
			Address: address,
		})
		if err != nil {
			panic(err)
		}
	} else {
		url += "/active"
		body, err = json.Marshal(UpdateActiveStatus{
			Id:     id,
			Active: active,
		})
		if err != nil {
			panic(err)
		}
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {

		var message Message
		err = json.Unmarshal(responseBody, &message)
		var e = ErrorResponse{
			Message:    message,
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}

		fmt.Printf("ErrorResponse\nMessage{ error: %v, msg: %s},\nStatus: %s,\nStatusCode: %d", e.Message.Error, e.Message.Msg, e.Status, e.StatusCode)
	} else {
		var response Response
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Message: %s\nSuccess: %v\n", response.Message, response.Success)
	}
}

func createCarrier(name string, address string, active bool, token string) {
	url := "http://localhost:8080/api/carriers"
	bearer := "Bearer " + token

	carrier := NewCarrier{Name: name, Address: address, Active: active}
	body, err := json.Marshal(carrier)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 201 {

		var message Message
		err = json.Unmarshal(responseBody, &message)
		var e = ErrorResponse{
			Message:    message,
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}

		fmt.Printf("ErrorResponse\nMessage{ error: %v, msg: %s},\nStatus: %s,\nStatusCode: %d", e.Message.Error, e.Message.Msg, e.Status, e.StatusCode)
	} else {
		var response Response
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Message: %s\nSuccess: %v\n", response.Message, response.Success)
	}
}

func getCarrier(i int, t string) {
	url := "http://localhost:8080/api/carriers/" + strconv.Itoa(i)
	bearer := "Bearer " + t

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var message Message
		err = json.Unmarshal(body, &message)
		var e = ErrorResponse{
			Message:    message,
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}

		fmt.Printf("ErrorResponse\nMessage{ error: %v, msg: %s},\nStatus: %s,\nStatusCode: %d", e.Message.Error, e.Message.Msg, e.Status, e.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	if response.Success != true {
		fmt.Printf("Error: %s\n", response.Message)
	} else {
		fmt.Printf("Carrier name: %s\nAddress: %s\nActive: %v\nCreated: %v\nUpdated: %v\n", response.Carrier.Name, response.Carrier.Address, response.Carrier.Active, response.Carrier.CreatedAt, response.Carrier.UpdatedAt)
	}
}

func getToken(user string, password string) {
	res, err := http.Post("http://localhost:8080/api/authenticate/", "application/json", bytes.NewBufferString(`{"User":"`+user+`","Password":"`+password+`"}`))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(errors.New(res.Status + " Server unavailable"))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var auth Auth
	err = json.Unmarshal(body, &auth)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bearer token: %s\n", auth.Token)
}
