package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	// "strings"
)

// --- DCR types (commented out for middleware task) ---
// type AuthResponse struct {
// 	AccessToken string `json:"access_token"`
// 	Scope       string `json:"scope"`
// 	ExpiresIn   int    `json:"expires_in"`
// 	TokenType   string `json:"token_type"`
// }

func main() {
	// --- DCR flow (commented out for middleware task) ---
	// domain := "dev-ftk4njr62natyltu.eu.auth0.com"
	// clientID := "NGcAvlkguhAuiEjwq8Fulr9RORM3HbkK"
	// clientSecret := "GsXPtx_ZejSYWADZm4Ojca7ClfL-5uoB5s97tZy22yJ_Kh8SetnI32b7txdRUnvw"
	// audience := "https://" + domain + "/api/v2/"
	// token, err := getAccessToken(domain, clientID, clientSecret, audience)
	// if err != nil {
	// 	log.Fatalf("Failed to get token: %v", err)
	// }
	// fmt.Println("Successfully obtained Access Token.")
	// endpoint := "clients"
	// callManagementAPI(domain, token, endpoint)

// --- DCR: getAccessToken (commented out for middleware task) ---
// func getAccessToken(domain, id, secret, audience string) (string, error) {
// 	url := fmt.Sprintf("https://%s/oauth/token", domain)
// 	payloadString := fmt.Sprintf(
// 		`{"client_id":"%s","client_secret":"%s","audience":"%s","grant_type":"client_credentials"}`,
// 		id, secret, audience,
// 	)
// 	payload := strings.NewReader(payloadString)
// 	req, _ := http.NewRequest("POST", url, payload)
// 	req.Header.Add("content-type", "application/json")
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)
// 	var authRes AuthResponse
// 	err = json.Unmarshal(body, &authRes)
// 	if err != nil {
// 		return "", err
// 	}
// 	return authRes.AccessToken, nil
// }

// --- DCR: callManagementAPI (commented out for middleware task) ---
// func callManagementAPI(domain, token, endpoint string) {
// 	url := fmt.Sprintf("https://%s/api/v2/%s", domain, endpoint)
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Add("authorization", "Bearer "+token)
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Printf("API Request failed: %v", err)
// 		return
// 	}
// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)
// 	var prettyJSON bytes.Buffer
// 	error := json.Indent(&prettyJSON, body, "", "    ")
// 	if error != nil {
// 		log.Println("JSON parse error: ", error)
// 		return
// 	}
// 	fmt.Println(prettyJSON.String())
// }