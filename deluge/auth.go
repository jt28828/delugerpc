package deluge

import (
	"fmt"
)

func (c *Client) Login() (bool, error) {
	response, err := sendRequest[bool](c, "auth.login", c.password)
	if err != nil || response.Result == nil {
		return false, fmt.Errorf("login failed, subsequent requests not possible %v", err)
	}

	if *response.Result != true {
		return false, fmt.Errorf("login failed. Password may be incorrect")
	}

	return true, nil
}
