package main

import (  
    "fmt"
	"net/http"
)

func getStatus(res http.ResponseWriter) {
	fmt.Fprint(res, "{'status':'success','message':'Online'}")
}