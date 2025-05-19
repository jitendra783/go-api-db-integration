package api

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// 	"github.com/spf13/viper"
// )

// // Define Prometheus metrics

// func main() {
// 	// Create a new Gin router
// 	r := gin.Default()

// 	// Register a /metrics endpoint for Prometheus to scrape
// 	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

// 	// Define a sample handler for /hello
// 	r.GET("/hello", func(c *gin.Context) {
// 		// Record the start time of the request
// 		start := time.Now()

// 		// Simulate some processing
// 		time.Sleep(500 * time.Millisecond)

// 		// Record the request details in Prometheus metrics
// 		httpRequests.WithLabelValues(c.Request.Method, "200").Inc()                         // Increment the counter
// 		httpDuration.WithLabelValues(c.Request.Method).Observe(time.Since(start).Seconds()) // Record duration

// 		// Respond to the client
// 		c.String(http.StatusOK, "Hello, Prometheus with Gin!")
// 	})

// 	// Start the server on port 8080
// 	fmt.Println("Starting server on :8080")
// 	log.Fatal(r.Run(":8080"))
// }

// func fetchDataFromDB() ([]DataRow, error) {
// 	var rows []DataRow
// 	if err := db.Where("processed = ?", false).Find(&rows).Error; err != nil {
// 		return nil, err
// 	}
// 	return rows, nil
// }

// func callAPI(payload []byte) (map[string]interface{}, error) {
// 	apiURL := viper.GetString("api.url") // Get the API URL from the config
// 	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var response map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }
// func updateDB(id int, status string) error {
// 	if err := db.Model(&DataRow{}).Where("id = ?", id).Update("processed", true).Update("status", status).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
// func processRow(row DataRow, resultChan chan<- error) {
// 	defer wg.Done() // Signal that this goroutine is done when it finishes

// 	// Prepare the payload for the API request
// 	payload := map[string]interface{}{
// 		"id":   row.ID,
// 		"name": row.Name,
// 		"info": row.Info,
// 	}

// 	payloadBytes, err := json.Marshal(payload)
// 	if err != nil {
// 		resultChan <- fmt.Errorf("Error marshalling payload for row %d: %v", row.ID, err)
// 		return
// 	}

// 	// Call the API
// 	apiResponse, err := callAPI(payloadBytes)
// 	if err != nil {
// 		resultChan <- fmt.Errorf("Error calling API for row %d: %v", row.ID, err)
// 		return
// 	}

// 	// Check response and decide on the status
// 	status := "success"
// 	if apiResponse["status"] != "OK" {
// 		status = "failure"
// 	}

// 	// Update the database with the response status
// 	if err := updateDB(row.ID, status); err != nil {
// 		resultChan <- fmt.Errorf("Error updating DB for row %d: %v", row.ID, err)
// 		return
// 	}

// 	log.Printf("Successfully processed row %d with status %s", row.ID, status)
// 	resultChan <- nil // Send nil to indicate success
// }
