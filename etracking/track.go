package etracking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Payload struct {
	Courier           string     `json:"courier"`
	TrackingNumber    string     `json:"trackingNumber"`
}

type Response struct {
	Meta              Meta       `json:"meta"`
	Data              Data       `json:"data"`
}

type Meta struct {
	Code              int        `json:"code"`
	Message           string     `json:"message"`
}

type Data struct {
	Courier           string     `json:"courier"`
	CurrentStatus     string     `json:"currentStatus"`
	Detail            Detail     `json:"detail"`
	Timelines         []Timeline `json:"timelines"`
}

type Detail struct {
	Sender            string     `json:"sender"`
	Recipient         string     `json:"recipient"`
}

type Timeline struct {
	Date     		 string          `json:"date"`
	Details      []Event         `json:"details"`
}

type Event struct {
	DateTime     time.Time       `json:"dateTime"`
	Date     		 string          `json:"date"`
	Time     		 string          `json:"time"`
	Description  string          `json:"description"`
}

func (client *Client) Find(courier string, trackingNumber string) (*Response, error) {
	return client.track(courier, trackingNumber, APIEndpointTracksFind)
}

func (client *Client) KerryExpress(trackingNumber string) (*Response, error) {
	courier := "kerry_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ShopeeExpress(trackingNumber string) (*Response, error) {
	courier := "shopee_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) DhlExpress(trackingNumber string) (*Response, error) {
	courier := "dhl_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) JtExpress(trackingNumber string) (*Response, error) {
	courier := "jt_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) AlphaFast(trackingNumber string) (*Response, error) {
	courier := "alpha_fast"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BestExpress(trackingNumber string) (*Response, error) {
	courier := "best_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) CjLogistics(trackingNumber string) (*Response, error) {
	courier := "cj_logistics"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) SpeedD(trackingNumber string) (*Response, error) {
	courier := "speed_d"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) NinjaVan(trackingNumber string) (*Response, error) {
	courier := "ninja_van"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Ups(trackingNumber string) (*Response, error) {
	courier := "ups"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ScgExpress(trackingNumber string) (*Response, error) {
	courier := "scg_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BeeExpress(trackingNumber string) (*Response, error) {
	courier := "bee_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) InterExpress(trackingNumber string) (*Response, error) {
	courier := "inter_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) TntExpress(trackingNumber string) (*Response, error) {
	courier := "tnt_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Shippop(trackingNumber string) (*Response, error) {
	courier := "shippop"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Aramex(trackingNumber string) (*Response, error) {
	courier := "aramex"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) FedEx(trackingNumber string) (*Response, error) {
	courier := "fed_ex"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) TpLogistics(trackingNumber string) (*Response, error) {
	courier := "tp_logistics"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ItTransport(trackingNumber string) (*Response, error) {
	courier := "it_transport"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) NakhonchaiAir(trackingNumber string) (*Response, error) {
	courier := "nakhonchai_air"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) NimExpress(trackingNumber string) (*Response, error) {
	courier := "nim_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) FlashExpress(trackingNumber string) (*Response, error) {
	courier := "flash_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BsExpress(trackingNumber string) (*Response, error) {
	courier := "bs_express"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Shipjung(trackingNumber string) (*Response, error) {
	courier := "shipjung"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) GlobalCainiao(trackingNumber string) (*Response, error) {
	courier := "global_cainiao"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) JapanPost(trackingNumber string) (*Response, error) {
	courier := "japan_post"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ChinaPost(trackingNumber string) (*Response, error) {
	courier := "china_post"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) SkyBox(trackingNumber string) (*Response, error) {
	courier := "sky_box"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BookMyCargo(trackingNumber string) (*Response, error) {
	courier := "book_my_cargo"
	return client.track(courier, trackingNumber, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) track(courier, trackingNumber, endpoint string) (response *Response, err error) {
	payload := Payload{
		Courier: courier,
		TrackingNumber: trackingNumber,
	}

	payloadByte, _ := json.Marshal(payload)
	resp, err := client.post(endpoint, bytes.NewReader(payloadByte))

	if err != nil {
		return nil, err
	}

	defer closeResponse(resp)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	json.Unmarshal([]byte(bodyString), &response)

	return response, nil
}