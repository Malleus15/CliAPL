package api

import (
	"awesomeProject2/models"
	"awesomeProject2/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Communication struct{
	Client http.Client
	//username of the current user
	CurrentUser models.User
	token string
}
//constructor
func NewCommunication(Client http.Client) Communication {
	currentUser:=models.User{}
	return Communication{Client, currentUser, ""}
}

//registration
func (comm Communication) CreateUser(userCreateReq models.UserReq) map[string]interface{}{
	json_data, err := json.Marshal(userCreateReq)

	resp, err := comm.Client.Post("http://127.0.0.1:5000/users/", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	return m
}

// Login to get the token
func (comm Communication) Login(login models.LoginReq) string{
	jsonData, err := json.Marshal(login)

	resp, err := comm.Client.Post("http://127.0.0.1:5000/token/", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("Error %s", err)
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]string)
	err = json.Unmarshal(body, &m)
	tmp:= m["access_token"]
	return tmp
}

// DeleteUser to get the token
func (comm Communication) DeleteUser(login models.LoginReq) string{
	jsonData, err := json.Marshal(login)
	resp, err := comm.Client.Post("http://127.0.0.1:5000/users/delete/"+comm.token, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("Error %s", err)
		return "Error!"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]string)
	err = json.Unmarshal(body, &m)
	tmp := m["detail"]
	return tmp
}

// GetReqUser to get a user from username
func  (comm Communication) GetReqUser(username string) models.User{
	resp, err := comm.Client.Get("http://127.0.0.1:5000/users/"+username)
	if err != nil {
	fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := models.User{}
	err = json.Unmarshal(body, &m)
	user := models.User{UserReq: models.UserReq{FirstName: m.FirstName, Name: m.Name, Username: m.Username, Password: m.Password}, ID: m.ID, IsActive: m.IsActive}
	return user
}

// CreateParameters creation and parameters verification of values
func (comm Communication) CreateParameters (params models.ParametersReq) map[string]interface{}{

	jsonData, err := json.Marshal(params)

	resp, err := comm.Client.Post("http://127.0.0.1:5000/users/parameters/"+comm.token, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	return m

}

// ReadParamsUser read parameters of a user
func (comm Communication) ReadParamsUser() []map[string] interface{}{
	s := strconv.Itoa(comm.GetCurrentUser().ID)
	resp, err := comm.Client.Get("http://127.0.0.1:5000/parameters/" + s + "/" + comm.token)

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res, err := utils.ParseJsonToMap(body)
	return res
}

func (comm Communication) DeleteParameters(paramsId int) string{
	s := strconv.Itoa(paramsId)
	resp, err := comm.Client.Get("http://127.0.0.1:5000/parameters/delete/" + s + "/" + comm.token)

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]string)
	err = json.Unmarshal(body, &m)
	tmp := m["detail"]
	return tmp
}

func (comm Communication) CreateInvest(req models.InvestmentReq) map[string]interface{}{

	jsonData, err := json.Marshal(req)

	resp, err := comm.Client.Post("http://127.0.0.1:5000/users/investment/"+comm.token, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	return m
}
//get all investments for a user

// ReadInvestUser read parameters of a user
func (comm Communication) ReadInvestUser() []map[string] interface{}{
	s := strconv.Itoa(comm.GetCurrentUser().ID)
	resp, err := comm.Client.Get("http://127.0.0.1:5000/investments/" + s + "/" + comm.token)

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res, err := utils.ParseJsonToMap(body)
	return res
}

func (comm Communication) DeleteInvest(invId int) string{
	s := strconv.Itoa(invId)
	resp, err := comm.Client.Get("http://127.0.0.1:5000/investments/delete/" + s + "/" + comm.token)

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]string)
	err = json.Unmarshal(body, &m)
	tmp := m["detail"]
	return tmp
}

func (comm *Communication) SetToken(token string) {
	comm.token = token
}

func (comm Communication) GetToken() string {
	return comm.token
}

func (comm *Communication) SetCurrentUser(CurrentUser models.User) {
	comm.CurrentUser = CurrentUser
}

func (comm Communication) GetCurrentUser() models.User {
	return comm.CurrentUser
}
