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

type Elasticsearch struct{}

func (es *Elasticsearch) request(method, url string, data interface{}) (map[string]interface{}, error) {
	ES_URL := os.Getenv("ES_URL")
	ES_USERNAMES := os.Getenv("ES_USERNAME")
	ES_PASSWORD := os.Getenv("ES_PASSWORD")

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, ES_URL+url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(ES_USERNAMES, ES_PASSWORD)
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		log.Printf("Error : %v", result)
		return nil, fmt.Errorf("HTTP error: %v", resp.Status)
	}
	log.Printf("Response : %v", result)
	return result, nil
}

func (es *Elasticsearch) CreateIndex(index string, mappings interface{}) (map[string]interface{}, error) {
	return es.request("PUT", "/"+index, map[string]interface{}{"mappings": mappings})
}

func (es *Elasticsearch) InsertMany(index string, dataArray []map[string]interface{}) (map[string]interface{}, error) {
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

func (es *Elasticsearch) InsertOne(index string, docId string, doc map[string]interface{}) (map[string]interface{}, error) {
	return es.request("POST", fmt.Sprintf("/%v/_doc/%v", index, docId), doc)
}

func (es *Elasticsearch) Search(index string, query map[string]interface{}) (map[string]interface{}, error) {
	return es.request("GET", fmt.Sprintf("/%v/_search", index), query)
}

func (es *Elasticsearch) UpdateOne(index string, docId string, doc map[string]interface{}) (map[string]interface{}, error) {
	return es.request("PUT", fmt.Sprintf("/%v/_update/%v", index, docId), doc)
}

func (es *Elasticsearch) DeleteOne(index string, docId string) (map[string]interface{}, error) {
	return es.request("DELETE", fmt.Sprintf("/%v/_doc/%v", index, docId), nil)
}

func (es *Elasticsearch) DeleteIndex(index string) (map[string]interface{}, error) {
	return es.request("DELETE", "/"+index, nil)
}
