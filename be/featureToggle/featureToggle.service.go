package featuretoggle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/artubric/toggleWorld/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const featureToggleBasePath = "featuretoggle"

// SetupRouters ...
func SetupRouters(apiBasePath string) {
	featureToggleItemHandler := http.HandlerFunc(handleFeatureToggle)
	featureToggleListHandler := http.HandlerFunc(handleFeatureToggles)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, featureToggleBasePath), middleware.Cors(featureToggleListHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, featureToggleBasePath), middleware.Cors(featureToggleItemHandler))
}

func handleFeatureToggle(w http.ResponseWriter, r *http.Request) {
	urlPathParameter := strings.Split(r.URL.Path, featureToggleBasePath+"/")
	inputFeatureToggleID := urlPathParameter[len(urlPathParameter)-1]
	objectID, err := primitive.ObjectIDFromHex(inputFeatureToggleID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	foundFeatureToggle := getFeatureToggle(objectID)
	if foundFeatureToggle == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		foundFeatureToggleJSON, err := json.Marshal(&foundFeatureToggle)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(foundFeatureToggleJSON)
	case http.MethodPut:
		var updatedFeatureToggle FeatureToggle

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &updatedFeatureToggle)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updatedFeatureToggle.ID != foundFeatureToggle.ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		addOrUpdateFeatureToggle(updatedFeatureToggle)
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		removeFeatureToggle(foundFeatureToggle.ID)
		w.WriteHeader(http.StatusOK)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleFeatureToggles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fetchedToggles := getFeatureToggleList()
		featureToggleJSON, err := json.Marshal(&fetchedToggles)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(featureToggleJSON)
	case http.MethodPost:
		var newFeatureToggle FeatureToggle
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &newFeatureToggle)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newFeatureToggle.ID != primitive.NilObjectID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := addOrUpdateFeatureToggle(newFeatureToggle)
		w.WriteHeader(http.StatusCreated)
		jsonID, err := json.Marshal(id.Hex())
		w.Write([]byte(jsonID))
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
