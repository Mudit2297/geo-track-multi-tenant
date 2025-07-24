package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	u "simulator/utils"
	"time"
)

const (
	authURL          = "http://localhost:8084/login"
	locationURL      = "http://localhost:8082/api/location"
	createTenantURL  = "http://localhost:8081/api/tenants"
	getAllTenantsURL = "http://localhost:8081/api/tenants"
	getTenantByIdURL = "http://localhost:8081/api/tenant/"
)

func login(payload u.AuthPayload) (string, error) {
	_, body, err := u.ExecRequest("auth", authURL, "POST", payload, "")
	if err != nil {
		return "", err
	}

	var loginResp u.LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return "", err
	}

	if loginResp.Response.IDToken == "" {
		return "", fmt.Errorf("no token")
	}
	return loginResp.Response.IDToken, nil
}

func sendLocationUpdates(token string) error {
	for i := 0; i < 10; i++ {
		loc := u.Location{
			Latitude:  rand.Float64() + float64(i)*0.001,
			Longitude: rand.Float64() + float64(i)*0.001,
		}

		r, b, err := u.ExecRequest("loc", locationURL, "POST", loc, token)
		if err != nil {
			return err
		}

		if r.StatusCode != 200 {
			fmt.Printf("Status code: %v, body: %v\n", r.StatusCode, string(b))
			continue
		}

		fmt.Printf("Sent location: %+v\n", loc)
		time.Sleep(2 * time.Second)
	}
	return nil
}

func createTenant(payload u.CreateTenantPayload, token string) error {
	req, body, err := u.ExecRequest("auth", createTenantURL, "POST", payload, token)
	if err != nil {
		return err
	}

	if req.StatusCode != 201 {
		return fmt.Errorf("failed to create tenant record in DB. status code: %v. body: %v", req.StatusCode, string(body))
	}

	return nil
}

func listAllTenants(token string) error {
	_, body, err := u.ExecRequest("auth", getAllTenantsURL, "GET", nil, token)
	if err != nil {
		return err
	}

	var resp u.TenantsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	for _, v := range resp {
		fmt.Printf("Name: %v\nID: %v\nEmail: %v\n\n", v.Name, v.ID, v.ContactEmail)
	}

	return nil
}

func getTenantDetailsByID(url string, token string) error {
	_, body, err := u.ExecRequest("auth", url, "GET", nil, token)
	if err != nil {
		return err
	}

	var resp u.TenantByIDResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %v\nID: %v\nEmail: %v\n\n", resp.Name, resp.ID, resp.ContactEmail)

	return nil
}

func main() {
	// Define credentails for authentication
	pl := u.AuthPayload{
		Username: "admin-t1",
		Password: "Admin@1234",
	}

	fmt.Println("Logging in to auth-service...")

	// Login using the creds and generate the token
	token, err := login(pl)
	if err != nil {
		log.Fatalf("Login failed: %v\n", err)
		return
	} else {
		fmt.Println("Login successful")
		fmt.Println()
	}

	// Create a new tenant - only works for admin users
	tpl := u.CreateTenantPayload{
		Name:         "b3systems",
		ContactEmail: "admin@b3systems.com",
	}
	err = createTenant(tpl, token)
	if err != nil {
		fmt.Printf("Tenant creation failed: %v\n", err)
		return
	} else {
		fmt.Println("Tenant created")
		fmt.Println()
	}

	// List all tenants - only works for admin users
	err = listAllTenants(token)
	if err != nil {
		fmt.Printf("Listing tenants failed: %v\n", err)
		return
	} else {
		fmt.Println("All tenants listed")
		fmt.Println()
	}

	// Get tenant for a particular ID - only works for admin users
	err = getTenantDetailsByID(getTenantByIdURL+"6bab4ba9-54de-4c02-b7ab-99efd58fcc91", token)
	if err != nil {
		fmt.Printf("Getting tenant details failed: %v\n", err)
		return
	} else {
		fmt.Println("Tenant details retrieved")
		fmt.Println()
	}

	// Send location updates for 20 sec. Time be changed in function logic
	err = sendLocationUpdates(token)
	if err != nil {
		fmt.Printf("Location update failed: %v\n", err)
		return
	} else {
		fmt.Println("Location submission successful")
		fmt.Println()
	}

}
