package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// func init() {
// 	functions.HTTP("HandleReq", HandleReq)
// }

type parameters struct {
	DelayForCcaas string `json:"delay_for_ccaas"`
}

type sessionInfo struct {
	Parameters parameters `json:"parameters"`
}

type resJson struct {
	SessionInfo sessionInfo `json:"sessionInfo"`
}

func HandleReq(w http.ResponseWriter, r *http.Request) {

	reqBody, reqReadErr := io.ReadAll(r.Body)

	if reqReadErr != nil {
		fmt.Printf("Error reading the request body : %s", reqReadErr)
		return
	}

	fmt.Printf("incoming request : %s", string(reqBody))

	fmt.Println("sleeping...")

	time.Sleep(1 * time.Second)

	fmt.Println("waking...")

	w.Header().Set("Content-Type", "application/json")

	resp := resJson{
		SessionInfo: sessionInfo{
			Parameters: parameters{
				DelayForCcaas: "True",
			},
		},
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		fmt.Printf("Error in JSON marshal. Err is : %s", err)
	}

	w.Write(jsonResp)

	fmt.Println("success")
}

func main() {

	http.HandleFunc("/", HandleReq)

	http.ListenAndServe(":8080", nil)
}