package utility

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	uuid "github.com/satori/go.uuid"
)

// Response ...
type Response struct {
	Ok   bool
	Data interface{}
}

// SendResponse ...
func SendResponse(w http.ResponseWriter, ok bool, data interface{}) {
	response := Response{
		Ok:   ok,
		Data: data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}

// ValidateConfig ...
func ValidateConfig(config *sarama.Config) error {
	err := config.Validate()
	if err != nil {
		log.Println("something wrong with config file", err)
		return err
	}
	return nil
}

// ToMap ...
func ToMap(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		log.Println("error in decoding request body", err)
		return nil, err
	}

	return resp, nil
}

// GenUUIDv4 ...
func GenUUIDv4() [16]byte {
	return uuid.NewV4()
}
