package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Test 1: Simple GET request
	resp1, err := http.Get("http://localhost:3000/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp1.Body.Close()

	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Test 1 - Main page:")
	fmt.Println("Status:", resp1.Status) // Should be 200
	fmt.Println("Body:", string(body1))  // Should be "This is the main page"

	// Test 2: HTTP 301 Redirect
	resp2, err := http.Get("http://localhost:3000/redirect301")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("\nTest 2 - HTTP 301 Redirect:")
	fmt.Println("Status:", resp2.Status)         // Should be 200 (after following redirect)
	fmt.Println("Final URL:", resp2.Request.URL) // Should be "http://localhost:3000/redirected"
	fmt.Println("Body:", string(body2))          // Should be "This is the redirected page"

	// Test 3: HTTP 302 Redirect
	resp3, err := http.Get("http://localhost:3000/redirect302")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp3.Body.Close()

	body3, err := io.ReadAll(resp3.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("\nTest 3 - HTTP 302 Redirect:")
	fmt.Println("Status:", resp3.Status)         // Should be 200 (after following redirect)
	fmt.Println("Final URL:", resp3.Request.URL) // Should be "http://localhost:3000/redirected"
	fmt.Println("Body:", string(body3))          // Should be "This is the redirected page"
}
