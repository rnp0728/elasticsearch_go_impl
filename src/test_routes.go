package src

import (
	"encoding/json"
	"net/http"
)

/*
FUNCTIONS
- CREATE INDEX
- INSERT MANY
- INSERT ONE
- SEARCH
- UPDATE ONE
- DELETE ONE
- DELETE INDEX
*/

var es Elasticsearch

// POST api/create/{index}
func CreateIndex(w http.ResponseWriter, r *http.Request) error {
	index := r.PathValue("index")

	var mappings interface{}
	err := json.NewDecoder(r.Body).Decode(&mappings)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if respData, err = es.CreateIndex(index, mappings); err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

func InsertMany(w http.ResponseWriter, r *http.Request) error {
	index := r.PathValue("index")

	var dataArray []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&dataArray)

	if err != nil {
		return err
	}
	var respData map[string]interface{}
	if respData, err = es.InsertMany(index, dataArray); err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

func InsertOne(w http.ResponseWriter, r *http.Request) error {
	// get the index from request "r" params
	index := r.PathValue("index")
	var doc map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if respData, err = es.InsertOne(index, doc["mongo_id"].(string), doc); err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

func Search(w http.ResponseWriter, r *http.Request) error {
	index := r.PathValue("index")
	var query map[string]interface{}
	var err error
	if err = json.NewDecoder(r.Body).Decode(&query); err != nil {
		return err
	}
	var respData map[string]interface{}
	if respData, err = es.Search(index, query); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}

func UpdateOne(w http.ResponseWriter, r *http.Request) error {
	index := r.PathValue("index")
	docId := r.URL.Query().Get("docId")
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if respData, err = es.UpdateOne(index, docId, body); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)

	return nil
}

func DeleteOne(w http.ResponseWriter, r *http.Request) error {
	index := r.PathValue("index")
    docId := r.URL.Query().Get("docId")

    var respData map[string]interface{}
	var err error

    if respData, err = es.DeleteOne(index, docId); err != nil {
        return err
    }

    w.WriteHeader(http.StatusOK)
    resp := NewSuccessResponse(respData)
    jsonEncoder := json.NewEncoder(w)
    jsonEncoder.Encode(resp)

    return nil
}

func DeleteIndex(w http.ResponseWriter, r *http.Request) error {
	index := r.PathValue("index")

	var respData map[string]interface{}
	var err error

	if respData, err = es.DeleteIndex(index); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	resp := NewSuccessResponse(respData)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(resp)
	return nil
}
