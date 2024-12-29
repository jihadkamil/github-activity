package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// /*
type Response struct {
	Data []Data `json:"data"`
}

type Data struct {
	Data DataDetail `json:"data"`
}
type DataDetail struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor Actor  `json:"actor"`
	Repo  Repo   `json:"repo"`
	// Payload   Requ   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

// Define supporting structs for Actor, Repo, and Payload
type Actor struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// */

func main() {
	fmt.Println("github-activity")
	// /*
	// Create a new reader to read input from the standard input (terminal)
	reader := bufio.NewReader(os.Stdin)

	for {
		// Read the input from the user until a newline character is encountered
		fmt.Println("input username>")
		username, _ := reader.ReadString('\n')

		// Trim any leading and trailing whitespace characters (including the newline) from the input
		username = strings.TrimSpace(username)
		// */
		// username := "jihadkamil"

		if username == "" {
			fmt.Println("Please enter a valid username")
			continue
		} else if username == ":q" {
			fmt.Println("exit")
			return
		}

		url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("error fetch %s %s\n", url, err.Error())
		}

		defer response.Body.Close()

		// read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("error reading response body %s\n", err.Error())
		}
		var result interface{}
		// var result Response
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		switch result := result.(type) {
		case []interface{}:
			if len(result) == 0 {
				fmt.Println("No events found for this user.")
			} else {
				fmt.Println("This is the result:\n", result)
			}
		case map[string]interface{}:
			fmt.Println("Received an unexpected JSON object:", result)
		default:
			fmt.Println("Received an unexpected JSON type.")
		}
		fmt.Println("================================================================================================")
	}
}
