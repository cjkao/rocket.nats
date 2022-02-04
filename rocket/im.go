package rocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateImRequest struct {
	Username string `json:"username"`
}

type CreateImResponse struct {
	Room struct {
		T         string   `json:"t"`
		Rid       string   `json:"rid"`
		Usernames []string `json:"usernames"`
	} `json:"room"`
	Success bool `json:"success"`
}

// Creates a new channel.
func (c *Client) CreateIm(param *CreateImRequest) (*CreateImResponse, error) {
	opt, _ := json.Marshal(param)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/%s/im.create", c.baseURL, c.apiVersion),
		bytes.NewBuffer(opt))

	if err != nil {
		return nil, err
	}

	res := CreateImResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
