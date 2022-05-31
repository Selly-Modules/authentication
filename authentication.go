package authentication

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Selly-Modules/natsio"
)

// Client ...
type Client struct {
	Config        Config
	natsServer    natsio.Server
	natsJetStream natsio.JetStream
	Request       Request
}

var (
	client *Client
)

// NewClient
// Init client elasticsearch
func NewClient(config Config) (*Client, error) {
	if config.ApiKey == "" {
		return nil, errors.New("api key is required")
	}
	if config.Nats.URL == "" {
		return nil, errors.New("nats url is required")
	}
	if err := natsio.Connect(config.Nats); err != nil {
		return nil, fmt.Errorf("nats connect failed: %v", err)
	}

	client = &Client{
		Config:        config,
		natsServer:    natsio.GetServer(),
		natsJetStream: natsio.GetJetStream(),
		Request:       Request{},
	}

	return client, nil
}

// requestNats
// publish message to nats and waiting response
func requestNats(subject string, data []byte) (*Response, error) {
	var (
		req = RequestBody{
			ApiKey: client.Config.ApiKey,
			Body:   data,
		}
		res *Response
	)
	msg, err := client.natsServer.Request(subject, toBytes(req))
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(msg.Data, &res); err != nil {
		return nil, err
	}
	if res.Message != "" {
		return nil, errors.New(res.Message)
	}
	return res, nil
}

func toBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
