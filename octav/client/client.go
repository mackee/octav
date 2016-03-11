// DO NOT EDIT. Automatically generated by hsup at Fri, 11 Mar 2016 13:43:55 JST
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/builderscon/octav/octav"
	"github.com/lestrrat/go-urlenc"
)

type Client struct {
	Client   *http.Client
	Endpoint string
}

func New(s string) *Client {
	return &Client{
		Client:   &http.Client{},
		Endpoint: s,
	}
}

func (c *Client) CreateConference(in *octav.CreateConferenceRequest) (*octav.Conference, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/conference/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Conference
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateRoom(in *octav.Room) (*octav.Room, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/room/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Room
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateSession(in *octav.CreateSessionRequest) (*octav.Session, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/session/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Session
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateUser(in *octav.CreateUserRequest) (*octav.User, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/user/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.User
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) CreateVenue(in *octav.Venue) (*octav.Venue, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/venue/create")
	if err != nil {
		return nil, err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Venue
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) DeleteConference(in *octav.DeleteConferenceRequest) error {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/conference/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteRoom(in *octav.DeleteRoomRequest) error {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/room/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteUser(in *octav.DeleteUserRequest) error {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/user/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) DeleteVenue(in *octav.DeleteVenueRequest) error {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/venue/delete")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}

func (c *Client) ListRooms(in *octav.ListRoomRequest) ([]octav.Room, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/room/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload []octav.Room
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListSessionsByConference(in *octav.ListSessionsByConferenceRequest) (interface{}, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/schedule/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload interface{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) ListVenues(in *octav.ListVenueRequest) ([]octav.Venue, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/venue/list")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload []octav.Venue
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) LookupConference(in *octav.LookupConferenceRequest) (*octav.Conference, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/conference/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Conference
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupRoom(in *octav.LookupRoomRequest) (*octav.Room, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/room/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Room
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupSession(in *octav.LookupSessionRequest) (*octav.Session, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/session/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Session
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupUser(in *octav.LookupUserRequest) (*octav.User, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/user/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.User
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) LookupVenue(in *octav.LookupVenueRequest) (*octav.Venue, error) {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/venue/lookup")
	if err != nil {
		return nil, err
	}
	buf, err := urlenc.Marshal(in)
	if err != nil {
		return nil, err
	}
	u.RawQuery = string(buf)
	res, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	var payload octav.Venue
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) UpdateConference(in *octav.UpdateConferenceRequest) error {
	var err error
	u, err := url.Parse(c.Endpoint + "/v1/conference/update")
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(in)
	if err != nil {
		return err
	}
	res, err := c.Client.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`Invalid response: '%s'`, res.Status)
	}
	return nil
}
