package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	fmt.Println("Starting the application...")
	client := &http.Client{}
	// response, err := client.Get("http://desktop-va5b2l8/v1/rest/auth")
	req, err := http.NewRequest("GET", "http://desktop-va5b2l8/v1/rest/auth", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	/*
		fmt.Println()
		fmt.Println("Request header", req.Header)
		fmt.Println("=========== Get Body is next =============")
		fmt.Println()
		for x, h := range req.Header {
			fmt.Println(x, h)
		}
	*/
	response, err := client.Do(req)
	/*
		fmt.Println("=========== Response is received =============")
		fmt.Println()
	*/
	anp := response.Header["Auth-Nonce"]
	// fmt.Println("Authentication Nonce:", anp)
	authNonceResp, err := strconv.Atoi(anp[0])
	if err != nil {
		fmt.Println("Error en la conversión del entero")
	}
	auth := ((authNonceResp / 13) % 99999) * 17
	// fmt.Println("Auth:", auth)

	req.Header.Set("auth-nonce", anp[0])
	req.Header.Set("auth-nonce-response", strconv.Itoa(auth))
	u, err := url.Parse("http://desktop-va5b2l8/v1/rest/auth?usr=sysadmin&pwd=sysadmin")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(u.RequestURI())
	req.URL = u
	// fmt.Println(req.Host, req.Method, req.URL)
	// fmt.Println(req.Header)
	// req.URL.Query.Add("http://desktop-va5b2l8/v1/rest/auth?user=sysadmin&pwd=sysadmin")

	response, err = client.Do(req)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	/*
		else {
			// defer response.Body.Close()
			fmt.Println("Response body")
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		for i, x := range response.Header {
			fmt.Println(i, x)
		}
		fmt.Println()
		fmt.Println("Auth-Nonce", response.Header["Auth-Nonce"])
		// Perform a GET on https://{SERVERNAME}/v1/rest/auth?usr={USER}&pwd={PASSWORD}
		// with both the Auth-Nonce and the Auth-Nonce-Response values in the header.
		// Replace the {USER} and {PASSWORD} with actual values.
	*/
	session := response.Header["Auth-Session"]

	u, err = url.Parse("http://desktop-va5b2l8/v1/rest/sit?ws=desktop-va5b2l8")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Session:", session)
	fmt.Println("================================ Get done ============================")
	/*
		response.Header["Auth-Nonce"] = anp
		response.Header["Auth-Nonce-Response"] = append(response.Header["Auth-Nonce-Response"], strconv.Itoa(auth))
		response.Header["Auth-Session"] = append(response.Header["Auth-Session"], "")
	*/
	req.Header.Set("Auth-Session", session[0])
	req.URL = u
	response, err = client.Do(req)
	if err != nil {
		fmt.Println("Authentication didn't work")
	} else {
		for i, x := range response.Header {
			fmt.Println(i, x)
		}
		fmt.Println()
	}

	u, err = url.Parse("http://desktop-va5b2l8/v1/rest/inventory/554139745000159172")
	if err != nil {
		log.Fatal(err)
	}
	req.URL = u
	response, err = client.Do(req)
	if err != nil {
		fmt.Println("Authentication didn't work")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read error is:", err)
	}
	fmt.Println(string(body))

	fmt.Println("Terminating the application...")
}
