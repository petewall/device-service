package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	URL string
}

func (c *Client) get(path, name string) ([]byte, error) {
	res, err := http.Get(c.URL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to request the %s: %w", name, err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s request failed: %s", name, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the %s response: %w", name, err)
	}

	return body, nil
}

func (c *Client) post(path, name string, body []byte) error {
	res, err := http.Post(c.URL+path, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to request the %s: %w", name, err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s request failed: %s", name, res.Status)
	}

	return nil
}

func (c *Client) GetAllDevices() ([]*Device, error) {
	name := "device list"
	body, err := c.get("/", name)
	if err != nil {
		return nil, err
	}

	var deviceList []*Device
	err = json.Unmarshal(body, &deviceList)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the %s response: %w", name, err)
	}

	return deviceList, nil
}

func (c *Client) GetDevice(mac string) (*Device, error) {
	name := "device " + mac
	body, err := c.get("/"+mac, name)
	if err != nil {
		return nil, err
	}

	var device *Device
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the %s response: %w", name, err)
	}

	return device, nil
}

func (c *Client) UpdateDevice(mac, firmwareType, firmwareVersion string) error {
	name := "update device " + mac
	payload := &UpdateDevicePayload{
		Firmware: firmwareType,
		Version:  firmwareVersion,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to create the payload for %s: %w", name, err)
	}

	return c.post("/"+mac, name, body)
}

func (c *Client) SetDeviceName(mac, name string) error {
	return c.post(
		"/"+mac+"/name?val="+url.QueryEscape(name),
		"set device "+mac+" name to "+name,
		nil)
}

func (c *Client) SetDeviceFirmwareType(mac, firmwareType string) error {
	return c.post(
		"/"+mac+"/firmware?val="+url.QueryEscape(firmwareType),
		"set device "+mac+" firmware type to "+firmwareType,
		nil)
}

func (c *Client) SetDeviceFirmwareVersion(mac, firmwareVersion string) error {
	return c.post(
		"/"+mac+"/version?val="+url.QueryEscape(firmwareVersion),
		"set device "+mac+" firmware version to "+firmwareVersion,
		nil)
}
