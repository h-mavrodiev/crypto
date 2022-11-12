package stex

import "errors"

func (c *StexClient) GetProfileInfo(ch chan<- interface{}) error {

	resource := "info"

	req, err := c.CreateGetRequest(c.Endpoints.Profile, resource, "", "")
	if err != nil {
		return errors.New("faild create get request for stex profile info")
	}

	c.Authenticate(req)

	res := InfoData{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed get request for stex profile info")
	}

	ch <- res

	return nil
}
