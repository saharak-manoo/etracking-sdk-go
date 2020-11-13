package etracking

import (
	"bytes"
	"encoding/json"
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

func (client *Client) Find(courier string, trackingNumber string) Response {
	return client.track(courier, trackingNumber, APIEndpointTracksFind)
}

func (client *Client) KerryExpress(trackingNumber string) Response {
	return client.track("kerry_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) ShopeeExpress(trackingNumber string) Response {
	return client.track("shopee_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) DhlExpress(trackingNumber string) Response {
	return client.track("dhl_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) JtExpress(trackingNumber string) Response {
	return client.track("jt_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) AlphaFast(trackingNumber string) Response {
	return client.track("alpha_fast", trackingNumber, APIEndpointTracks)
}

func (client *Client) BestExpress(trackingNumber string) Response {
	return client.track("best_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) CjLogistics(trackingNumber string) Response {
	return client.track("cj_logistics", trackingNumber, APIEndpointTracks)
}

func (client *Client) SpeedD(trackingNumber string) Response {
	return client.track("speed_d", trackingNumber, APIEndpointTracks)
}

func (client *Client) NinjaVan(trackingNumber string) Response {
	return client.track("ninja_van", trackingNumber, APIEndpointTracks)
}

func (client *Client) Ups(trackingNumber string) Response {
	return client.track("ups", trackingNumber, APIEndpointTracks)
}

func (client *Client) ScgExpress(trackingNumber string) Response {
	return client.track("scg_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) BeeExpress(trackingNumber string) Response {
	return client.track("bee_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) InterExpress(trackingNumber string) Response {
	return client.track("inter_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) TntExpress(trackingNumber string) Response {
	return client.track("tnt_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) Shippop(trackingNumber string) Response {
	return client.track("shippop", trackingNumber, APIEndpointTracks)
}

func (client *Client) Aramex(trackingNumber string) Response {
	return client.track("aramex", trackingNumber, APIEndpointTracks)
}

func (client *Client) FedEx(trackingNumber string) Response {
	return client.track("fed_ex", trackingNumber, APIEndpointTracks)
}

func (client *Client) TpLogistics(trackingNumber string) Response {
	return client.track("tp_logistics", trackingNumber, APIEndpointTracks)
}

func (client *Client) ItTransport(trackingNumber string) Response {
	return client.track("it_transport", trackingNumber, APIEndpointTracks)
}

func (client *Client) NakhonchaiAir(trackingNumber string) Response {
	return client.track("nakhonchai_air", trackingNumber, APIEndpointTracks)
}

func (client *Client) NimExpress(trackingNumber string) Response {
	return client.track("nim_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) FlashExpress(trackingNumber string) Response {
	return client.track("flash_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) BsExpress(trackingNumber string) Response {
	return client.track("bs_express", trackingNumber, APIEndpointTracks)
}

func (client *Client) Shipjung(trackingNumber string) Response {
	return client.track("shipjung", trackingNumber, APIEndpointTracks)
}

func (client *Client) GlobalCainiao(trackingNumber string) Response {
	return client.track("global_cainiao", trackingNumber, APIEndpointTracks)
}

func (client *Client) JapanPost(trackingNumber string) Response {
	return client.track("japan_post", trackingNumber, APIEndpointTracks)
}

func (client *Client) ChinaPost(trackingNumber string) Response {
	return client.track("china_post", trackingNumber, APIEndpointTracks)
}

func (client *Client) SkyBox(trackingNumber string) Response {
	return client.track("sky_box", trackingNumber, APIEndpointTracks)
}

func (client *Client) BookMyCargo(trackingNumber string) Response {
	return client.track("book_my_cargo", trackingNumber, APIEndpointTracks)
}

func (client *Client) track(courier, trackingNumber, endpoint string) Response {
	payload := Payload{
		Courier: courier,
		TrackingNumber: trackingNumber,
	}

	payloadByte, _ := json.Marshal(payload)
	resp, _ := client.post(endpoint, bytes.NewReader(payloadByte))
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	response := Response{}
	json.Unmarshal([]byte(bodyString), &response)

	return response
}