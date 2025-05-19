package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/errgroup"
)

// DB Config for Oracle
const (
	DBDriver = "godror"
	DBSource = "ebatest/password@localhost:1521/ORCLPDB1" // Update your Oracle DB connection string
)

// Global DB variable
var db *sqlx.DB

// Struct for the token generation request and response
type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

// Struct for the payload request and response
type Payload struct {
	Data string `json:"data"`
}

type ApiResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

// Struct for monitoring the new DB entries
type DBEntry struct {
	ID   int    `db:"id"`
	Data string `db:"data"`
}

func initDB() (*sqlx.DB, error) {
	// Connect to Oracle DB
	fmt.Println("db intiate")
	db, err := sqlx.Open(DBDriver, DBSource)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the Oracle database.", db)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the Oracle database.")
	return db, nil
}

func generateToken(username, password string) (string, error) {
	// Prepare the request to get the token
	reqBody := TokenRequest{Username: username, Password: password}
	reqBodyBytes, _ := json.Marshal(reqBody)

	// HTTP request to generate token
	resp, err := http.Post("http://localhost:8080/api/generate-token", "application/json", bytes.NewReader(reqBodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to call generate token API: %v", err)
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %v", err)
	}

	// Token is valid for 5 minutes, check if expiration logic works
	expirationTime, _ := time.Parse(time.RFC3339, tokenResp.Expires)
	if time.Now().After(expirationTime) {
		return "", fmt.Errorf("token has expired")
	}

	return tokenResp.Token, nil
}

func fetchApiData(token string) (*ApiResponse, error) {
	// Prepare the request payload
	payload := Payload{Data: "Request data from DB"}
	payloadBytes, _ := json.Marshal(payload)

	// Make the HTTP request to the API using the token
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/fetch-data", bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create API request: %v", err)
	}

	// Add Authorization header with the token
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	var apiResp ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	return &apiResp, nil
}

func updateDBEntry(entry *DBEntry) error {
	// Example of DB update for Oracle DB
	_, err := db.Exec("UPDATE my_table SET status = :1 WHERE id = :2", "processed", entry.ID)
	if err != nil {
		return fmt.Errorf("failed to update DB: %v", err)
	}
	return nil
}

// Goroutine to fetch data and update DB
func processEntry(entry *DBEntry, token string, wg *sync.WaitGroup, g errgroup.Group) {
	defer wg.Done()

	// Fetch data from API using the token
	apiResp, err := fetchApiData(token)
	if err != nil {
		log.Printf("Error fetching data for entry %d: %v", entry.ID, err)
		return
	}

	// Update DB with the response from the API
	err = updateDBEntry(entry)
	if err != nil {
		log.Printf("Error updating DB for entry %d: %v", entry.ID, err)
		return
	}

	log.Printf("Successfully processed entry %d, API response: %v", entry.ID, apiResp)
}

// Function to monitor DB for new entries and trigger the process
func MonitorNewEntries(wg *sync.WaitGroup, g errgroup.Group) {
	fmt.Println("inside moniter")
	// Generate token for the API calls
	// token, err := generateToken("your-username", "your-password")
	// if err != nil {
	// 	log.Printf("Error generating token: %v", err)
	// 	time.Sleep(5 * time.Second)
	// }
	for {
		// Query DB for new entries
		var entries []DBEntry
		// err := db.Select(&entries, "SELECT id, data FROM my_table WHERE status = 'pending' FETCH FIRST 10 ROWS ONLY")
		// if err != nil {
		// 	log.Printf("Error querying DB: %v", err)
		// 	time.Sleep(5 * time.Second)
		// 	continue
		// }

		// If no new entries, wait and retry
		// if len(entries) == 0 {
		// 	fmt.Println("no entries found")
		// 	time.Sleep(5 * time.Second)
		// 	continue
		// }
		entries = []DBEntry{{ID: 1, Data: "data"},{ID: 1, Data: "data"}}
        // entries = append(entries,entry)
		// Process each entry concurrently using goroutines and errgroup
		for _, entry := range entries {
			entry := entry // Capture entry in closure
			wg.Add(1)

			// Goroutine to process entry
			g.Go(func() error {
				processEntry(&entry, "token", wg, g)
				return nil
			})
		}

		// Wait for all goroutines to finish
		wg.Wait()
	}
}

func main() {
	fmt.Println("main")
	// var err error
	// Step 1: Initialize DB connection
	// db, err = initDB()
	// if err != nil {
	// 	log.Fatalf("Error initializing DB: %v", err)
	// 	os.Exit(1)
	// }

	// Step 2: Use an errgroup for concurrent execution
	var g errgroup.Group
	var wg sync.WaitGroup

	// Start monitoring new DB entries in a goroutine
	fmt.Println("start process")
	 MonitorNewEntries(&wg, g)
    fmt.Println("processEntry")
	// Wait for all processes to complete
	if err := g.Wait(); err != nil {
		log.Fatalf("Error in processing: %v", err)
	}
}
