// Package src provides HTTP handler functions for interacting with an Elasticsearch server.
package src

import (
	"encoding/json"
	"net/http"
)

/*
HTTP HANDLER FUNCTIONS

- CreateIndex: Handles the creation of an index in Elasticsearch.
- InsertMany: Handles the insertion of multiple documents into an index in Elasticsearch.
- InsertOne: Handles the insertion of a single document into an index in Elasticsearch.
- Search: Handles the search operation in Elasticsearch.
- UpdateOne: Handles the update of a single document in an index in Elasticsearch.
- DeleteOne: Handles the deletion of a single document from an index in Elasticsearch.
- DeleteIndex: Handles the deletion of an index from Elasticsearch.
*/

var es Elasticsearch

// CreateIndex handles the creation of an index in Elasticsearch.
// It expects the index name in the URL path and index mappings in the request body.
func CreateIndex(w http.ResponseWriter, r *http.Request) error {
	// Extract index name from URL path.
	index := r.PathValue("index")

	// Decode index mappings from request body.
	var mappings interface{}
	err := json.NewDecoder(r.Body).Decode(&mappings)
	if err != nil {
		return err
	}

	// Call Elasticsearch CreateIndex method.
	var respData map[string]interface{}
	if respData, err = es.CreateIndex(index, mappings); err != nil {
		return err
	}

	// Write success response.
	w.WriteHeader(http.StatusCreated)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

// InsertMany handles the insertion of multiple documents into an index in Elasticsearch.
// It expects the index name in the URL path and an array of documents in the request body.
func InsertMany(w http.ResponseWriter, r *http.Request) error {
	// Extract index name from URL path.
	index := r.PathValue("index")

	// Decode array of documents from request body.
	var dataArray []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&dataArray)
	if err != nil {
		return err
	}

	// Call Elasticsearch InsertMany method.
	var respData map[string]interface{}
	if respData, err = es.InsertMany(index, dataArray); err != nil {
		return err
	}

	// Write success response.
	w.WriteHeader(http.StatusCreated)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

// InsertOne handles the insertion of a single document into an index in Elasticsearch.
// It expects the index name in the URL path and the document ID in the request body.
func InsertOne(w http.ResponseWriter, r *http.Request) error {
	// Extract index name from URL path.
	index := r.PathValue("index")

	// Decode document from request body.
	var doc map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		return err
	}

	// Call Elasticsearch InsertOne method.
	var respData map[string]interface{}
	if respData, err = es.InsertOne(index, doc["mongo_id"].(string), doc); err != nil {
		return err
	}

	// Write success response.
	w.WriteHeader(http.StatusCreated)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

// Search handles the search operation in Elasticsearch.
// It expects the index name in the URL path and the search query in the request body.
func Search(w http.ResponseWriter, r *http.Request) error {
	// Extract index name from URL path.
	index := r.PathValue("index")

	// Decode search query from request body.
	var query map[string]interface{}
	var err error
	if err = json.NewDecoder(r.Body).Decode(&query); err != nil {
		return err
	}

	// Call Elasticsearch Search method.
	var respData map[string]interface{}
	if respData, err = es.Search(index, query); err != nil {
		return err
	}

	// Write success response.
	w.WriteHeader(http.StatusOK)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

// UpdateOne handles the update of a single document in an index in Elasticsearch.
// It expects the index name and document ID in the URL path and the update data in the request body.
func UpdateOne(w http.ResponseWriter, r *http.Request) error {
	// Extract index name and document ID from URL path.
	index := r.PathValue("index")
	docId := r.URL.Query().Get("docId")

	// Decode update data from request body.
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	// Call Elasticsearch UpdateOne method.
	var respData map[string]interface{}
	if respData, err = es.UpdateOne(index, docId, body); err != nil {
		return err
	}

	// Write success response.
	w.WriteHeader(http.StatusOK)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)

	return nil
}

// DeleteOne handles the deletion of a single document from an index in Elasticsearch.
// It expects the index name and document ID in the URL path.
func DeleteOne(w http.ResponseWriter, r *http.Request) error {
	// Extract index name and document ID from URL path.
	index := r.PathValue("index")
	docId := r.URL.Query().Get("docId")

	// Call Elasticsearch DeleteOne method.
	var respData map[string]interface{}
	var err error
	if respData, err = es.DeleteOne(index, docId); err != nil {
		return err
	}

	// Write success response.
	w.WriteHeader(http.StatusOK)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)

	return nil
}

// DeleteIndex handles the deletion of an index from Elasticsearch.
// It expects the index name in the URL path.
func DeleteIndex(w http.ResponseWriter, r *http.Request) error {
	// Extract index name from URL path.
	index := r.PathValue("index")

	// Call Elasticsearch DeleteIndex method.
	var respData map[string]interface{}
	var err error
	if respData, err = es.DeleteIndex(index); err != nil {
		return err
	}

	// Write success response.
	w.WriteHeader(http.StatusOK)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}
