package main

//CRUD для структуры “пользователь”
//Нужно реализовать какую-то структуру User. И http server, который умеет отвечать на 4 запроса /{id} GET, POST, /{id} PUT, /{id} DELETE
//Данные передаются в JSON’е.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"full-name"`
	City string `json:"city"`
}

var users = make(map[int]User)

func main() {
	http.HandleFunc("/users", handler)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		createUser(w, r)
	case "GET":
		getUser(w, r)
	case "PUT":
		updateUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	index, _ := strconv.Atoi(id)
	delete(users, index)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var u User
	err = json.Unmarshal(body, &u)
	users[u.Id] = u
	fmt.Fprint(w, "Successfully updated")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	index, _ := strconv.Atoi(id)
	user := users[index]
	fmt.Fprint(w, unmarshalToString(user))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var u User
	err = json.Unmarshal(body, &u)
	u.Id = len(users)
	users[len(users)] = u
	fmt.Fprint(w, unmarshalToString(u))
}

func unmarshalToString(user User) string {
	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	s := string(bytes)
	return s
}
