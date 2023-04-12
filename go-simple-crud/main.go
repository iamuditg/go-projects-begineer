package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Email    interface{} `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func main() {

	getDataFromURL()

}

func getDataFromURL() {
	url := "https://jsonplaceholder.typicode.com/users/1"
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return
	}
	fmt.Println(user)
}
