package secret

import (
	"encoding/csv"
	"fmt"
	"os"
)

func GetCredentials() (string, string) {
	// Open the CSV file
	file, err := os.Open("rzp.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return "Exception", err.Error()
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return "Exception", err.Error()
	}

	var keyID string
	var keySecret string

	// Iterate over each record and print key_id and key_secret
	for _, record := range records {
		keyID = record[0]
		keySecret = record[1]
	}

	return keyID, keySecret
}
