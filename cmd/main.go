package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ResultSummary struct {
	Body []BodyItem `json:""`
}

type BodyItem struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Assets []AssetItem `json:"assets"`
}

type AssetItem struct {
	Name          string `json:"name"`
	DownloadCount int    `json:"download_count"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: github-release OWNER REPO BEARER_TOKEN")
		os.Exit(1)
	}

	owner := os.Args[1]
	repo := os.Args[2]

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var result []BodyItem
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing body:", err)
		os.Exit(1)
	}

	for _, bodyItem := range result {
		fmt.Printf("ID: %d, Name: %s\n", bodyItem.ID, bodyItem.Name)
		for _, asset := range bodyItem.Assets {
			fmt.Printf("  Asset Name: %s, Download Count: %d\n", asset.Name, asset.DownloadCount)
		}
	}

	err = os.WriteFile("results.json", body, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Println("Results written to results.json")
}
