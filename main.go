package main

import (
	"fmt"
	"os"
	"strconv"
)

var maxUsers = 100
var totalUsers = 0
var currentUser = "unknown"

func initializeContract() {
	if os.Getenv("DB_INITIALIZED") == "true" {
		panic("Contract is already initialized")
	}

	os.Stdout.WriteString(fmt.Sprintf("DBW=MAX_USERS=%d\n", maxUsers))
	os.Stdout.WriteString("DBW=DB_INITIALIZED=true\n")
}

func registerUser(username string) {
	if os.Getenv("DB_INITIALIZED") != "true" {
		panic("Contract is not initialized")
	}

	totalUsers, _ = strconv.Atoi(os.Getenv("DB_TOTAL_USERS"))
	
	if totalUsers+1 > maxUsers {
		panic("Exceeded maximum user limit")
	}

	if currentUser = os.Getenv("DB_CURRENT_USER"); currentUser == "" {
		currentUser = "unknown"
	}

	os.Stdout.WriteString(fmt.Sprintf("OUT=previous_user: %s\n", currentUser))

	os.Stdout.WriteString(fmt.Sprintf("DBW=DB_CURRENT_USER=%s\n", username))
	os.Stdout.WriteString(fmt.Sprintf("DBW=USER_%d=%s\n", totalUsers, username))
	os.Stdout.WriteString(fmt.Sprintf("DBW=DB_TOTAL_USERS=%d\n", totalUsers+1))
}

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("Error: Missing arguments!\n")
		os.Exit(1)
	}

	if os.Args[1] == "initialize" {
		initializeContract()
		os.Exit(0)
	}

	if os.Args[1] == "register" && len(os.Args) == 3 {
		registerUser(os.Args[2])
		os.Exit(0)
	}

	os.Stderr.WriteString(fmt.Sprintf("Error: Invalid command or arguments: %s\n", os.Args[1]))
	os.Exit(1)
}