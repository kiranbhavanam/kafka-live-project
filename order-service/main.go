//

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct{
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter,r *http.Request){
	response:=HealthResponse{Status:"healthy"}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func main(){
	http.HandleFunc("/health",healthHandler)
	log.Printf("Server listening on 8080")
	if err:=http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}