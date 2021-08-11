package connectors

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var client = &http.Client{}

//Connectors return a list of all active connectors
func Connectors(conf Configuration) ([]string, error) {
	req, err := http.NewRequest("GET", conf.BaseUrl()+"/connectors", nil)
	if err != nil {
		return nil, err
	}

	basicAuth(req, conf.Auth())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		log.Print("Error in closing request body", err)
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var connectors []string
	err = json.Unmarshal(body, &connectors)
	if err != nil {
		return nil, err
	}

	return connectors, nil
}

func basicAuth(request *http.Request, auth Authentication) {
	encoded := base64.StdEncoding.EncodeToString([]byte(auth.UserName() + ":" + auth.Password()))
	request.Header.Add("Authorization", "Basic "+encoded)
}
