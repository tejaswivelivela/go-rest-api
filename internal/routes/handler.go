package routes

import (
	"fmt"
	"go-rest-api/internal/model"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func GetSchema(w http.ResponseWriter, r *http.Request) {
	// localhost:8000/getschema?name=teju&id=123
	//fmt.Printf("url query:%v\n", r.URL.Query)
	//fmt.Printf("url host:%v\n", r.URL.Hostname)
	var err error
	var req model.Request
	err = schema.NewDecoder().Decode(&req, r.URL.Query())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("destination request:", req)
	//res := fmt.Sprintf("hi,%v,age is:%d", req.Name, req.Age)
	//byteArray := []byte(res)
	//w.Write(byteArray)
	res := model.Response{
		Name: req.Name,
		Age : req.Age,
		Role :"golang developer"
	}
	json.NewDecoder(w).Decode(res)
}
