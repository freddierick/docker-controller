package main

import (  
    "fmt"
	"net/http"
	"strings"
)

func startAPI(cfg Config) {
	logMsg("Starting HTTP API server...", "startup")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		logMsg("API REQUEST "+req.Method+" "+req.URL.Path, "debug" )


		var authKey = ""
		for name, headers := range req.Header {
			for _, h := range headers {
				if name == "Authentication"{
					authKey = h
				}
			}
		}
		if (authKey == cfg.Panel.AuthenticationTokenId) == false{
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "{'status':'error','message':'Invalid Authentication'}")
			return
		}
		args := strings.Split(req.URL.Path, "/")
		if args[1] == "api" {
			if args[2] == "status" {
				getStatus(w)
				return
			}
			if args[2] == "server" {
				if req.Method == "GET" {
					logMsg("SNEDING SERVER INFO FOR "+args[3], "prosess")
					fmt.Fprint(w, "{'status':'success','message':'see server info'}")
					return
				}
				if req.Method == "POST" {
					logMsg("CREATING NEW SERVER", "prosess")
					fmt.Fprint(w, "{'status':'success','message':'Server created'}")
					return
				}
				if req.Method == "DELETE" {
					fmt.Fprint(w, "{'status':'success','message':'Server delited'}")
					logMsg("DELETING SERVER "+args[3], "prosess")
					return
				}
				if req.Method == "PUT" {
					fmt.Fprint(w, "{'status':'success','message':'Server state changed'}")
					logMsg("ALTERING SERVER STATE "+args[3], "prosess")
					return
				}
				if req.Method == "PATCH" {
					fmt.Fprint(w, "{'status':'success','message':'Server EDETING'}")
					logMsg("EDETING SERVER "+args[3], "prosess")
					return
				}
			}
		} else {

		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{'status':'error','message':'No endpoint found'}")
	})
	logMsg("Damon listening on port "+string(cfg.Ports.Daemon), "startup")
	http.ListenAndServe("0.0.0.0:"+cfg.Ports.Daemon, nil)
}