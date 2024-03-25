package src

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Elasticsearch represents a client for interacting with Elasticsearch server.
type Elasticsearch struct{}

// request performs an HTTP request to the Elasticsearch server.
func (es *Elasticsearch) request(method, url string, data interface{}) (map[string]interface{}, error) {
	// Fetch Elasticsearch server URL, username, and password from environment variables.
	ES_URL := os.Getenv("ES_URL")
	ES_USERNAMES := os.Getenv("ES_USERNAME")
	ES_PASSWORD := os.Getenv("ES_PASSWORD")

	// Marshal data into JSON format.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create HTTP client with custom Transport to allow skipping certificate verification.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Create HTTP request.
	req, err := http.NewRequest(method, ES_URL+url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set request headers and basic authentication.
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(ES_USERNAMES, ES_PASSWORD)

	// Execute HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode response JSON.
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Check for HTTP errors.
	if resp.StatusCode >= 400 {
		log.Printf("Error : %v", result)
		return nil, fmt.Errorf("HTTP error: %v", resp.Status)
	}

	log.Printf("Response : %v", result)
	return result, nil
}

// CreateIndex creates an index with specified mappings in Elasticsearch.
func (es *Elasticsearch) CreateIndex(index string, mappings interface{}) (map[string]interface{}, error) {
	return es.request("PUT", "/"+index, map[string]interface{}{"mappings": mappings})
}

// InsertMany inserts multiple documents into the specified index in Elasticsearch.
func (es *Elasticsearch) InsertMany(index string, dataArray []map[string]interface{}) (map[string]interface{}, error) {
	// Prepare bulk data for inserting multiple documents.
	bulkData := make([]string, 0, len(dataArray)*2)
	for _, doc := range dataArray {
		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": index,
				"_id":    doc["mongo_id"],
			},
		}
		metaJson, _ := json.Marshal(meta)
		docJson, _ := json.Marshal(doc)
		bulkData = append(bulkData, string(metaJson)+"\n"+string(docJson))
	}

	return es.request("POST", "/"+index+"/_bulk", strings.Join(bulkData, "\n")+"\n")
}

// InsertOne inserts a single document into the specified index in Elasticsearch.
func (es *Elasticsearch) InsertOne(index string, docId string, doc map[string]interface{}) (map[string]interface{}, error) {
	return es.request("POST", fmt.Sprintf("/%v/_doc/%v", index, docId), doc)
}

// Search performs a search query on the specified index in Elasticsearch.
func (es *Elasticsearch) Search(index string, query map[string]interface{}) (map[string]interface{}, error) {
	return es.request("GET", fmt.Sprintf("/%v/_search", index), query)
}

// UpdateOne updates a single document in the specified index in Elasticsearch.
func (es *Elasticsearch) UpdateOne(index string, docId string, doc map[string]interface{}) (map[string]interface{}, error) {
	return es.request("PUT", fmt.Sprintf("/%v/_update/%v", index, docId), doc)
}

// DeleteOne deletes a single document from the specified index in Elasticsearch.
func (es *Elasticsearch) DeleteOne(index string, docId string) (map[string]interface{}, error) {
	return es.request("DELETE", fmt.Sprintf("/%v/_doc/%v", index, docId), nil)
}

// DeleteIndex deletes the specified index from Elasticsearch.
func (es *Elasticsearch) DeleteIndex(index string) (map[string]interface{}, error) {
	return es.request("DELETE", "/"+index, nil)
}
