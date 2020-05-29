package msghandler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const EnvToken = "SUPERVISOR_TOKEN"
const SupervisorEndPoint = "http://supervisor/core/api/alexa"

type Handler struct {
	EndPoint string
}

func NewHandler() *Handler  {
	return &Handler{
		EndPoint: SupervisorEndPoint,
	}
}

// Handle initializes AlexaRequest struct with all intents and slots received in json message payload.
// Then it creates simplified filtered struct and performs execution with device control package.
func (h *Handler) Handle(req string) {
	log.Println("Request: " + req)

	err := h.postToApi(req)
	if err != nil {
		log.Println(err)
	}
}

func (h*Handler) postToApi(req string) error {
	token := os.Getenv(EnvToken)

	if token == "" {
		return errors.New("can not get supervisor token")
	}

	request, err := http.NewRequest("POST", h.EndPoint, strings.NewReader(req))
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer " + token)
	request.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)

	defer func() {
		if resp != nil && resp.Body != nil {
			err = resp.Body.Close()

			if err != nil {
				log.Println("error closing response body: ", err)
			}
		}
	}()

	if err != nil {
		return err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(responseBody))

	return nil
}
