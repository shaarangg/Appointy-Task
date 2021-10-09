package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUserPost(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Username": "Shaarang Singh",
		"Password": "password",
		"Email":    "shaarangg.singh@gmail.com",
	})
	reqBody := bytes.NewBuffer(postBody)
	res, err := http.Post("http://localhost:8081/users", "application/json", reqBody)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
func TestGetUserID(t *testing.T) {
	res, err := http.Get("http://localhost:8081/users/6161f2e3e6cc8dbf2a559459")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v", res.StatusCode, http.StatusOK)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	user := string(body)
	expected := `{"ID":"6161f2e3e6cc8dbf2a559459","Name":"Shaarang Singh","Email":"shaarangg.singh@gmail.com"}`
	if user != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", user, expected)
	}
}

func TestPostPost(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Caption":  "My fourth Post",
		"ImageURL": "github.com/images",
		"UserID":   "6161f2e3e6cc8dbf2a559459",
	})
	reqBody := bytes.NewBuffer(postBody)
	res, err := http.Post("http://localhost:8081/posts", "application/json", reqBody)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v", res.StatusCode, http.StatusOK)
	}
}

func TestGetPostId(t *testing.T) {
	res, err := http.Get("http://localhost:8081/posts/6161f5dc37246212fa47a41a")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v", res.StatusCode, http.StatusOK)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	post := string(body)
	expected := `{"ID":"6161f5dc37246212fa47a41a","Uid":"6161f5041c894b4f8a83f582","Caption":"My first Post","Image_url":"google.com/images","Timestamp":"2021-10-10 01:34:44.49766 +0530 IST m=+18.081275601"} `
	c := strings.Compare(post, expected)
	if c != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", post, expected)
		fmt.Println(c)
	}
}
func TestGetPostByUserId(t *testing.T) {
	res, err := http.Get("http://localhost:8081/posts/users/6161f2e3e6cc8dbf2a559459")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
