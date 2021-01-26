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
	TrackingNo        string     `json:"trackingNo"`
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
	TrackingNo    string       `json:"trackingNo"`
	Courier       string       `json:"courier"`
	CourierKey    string       `json:"courierKey"`
	Color         string       `json:"color"`
	Status        string       `json:"status"`
	CurrentStatus string       `json:"currentStatus"`
	Detail        ParcelDetail `json:"detail"`
	Timelines     []Timeline   `json:"timelines"`
}

type ParcelDetail struct {
	Sender                         string          `json:"sender"`
	Recipient                      string          `json:"recipient"`
	Qty                            int64           `json:"qty"`
	Address                        string          `json:"address"`
	OriginCity                     string          `json:"originCity"`
	OriginProvince                 string          `json:"originProvince"`
	OriginCountry                  string          `json:"originCountry"`
	DestinationCity                string          `json:"destinationCity"`
	DestinationProvince            string          `json:"destinationProvince"`
	DestinationCountry             string          `json:"destinationCountry"`
	Signer                         string          `json:"signer"`
	SignerImageURL                 string          `json:"signerImageURL"`
	Weight                         int64           `json:"weight"`
	ShippingService                string          `json:"shippingService"`
	ReturnShippingService          string          `json:"returnShippingService"`
	DeliveryType                   string          `json:"deliveryType"`
	SenderPhoneNumber              string          `json:"senderPhoneNumber"`
	RecipientPhoneNumber           string          `json:"recipientPhoneNumber"`
	DueDate                        string          `json:"dueDate"`
	CashOnDelivery                 float64         `json:"cashOnDelivery"`
	IsPayCashOnDelivery            bool            `json:"isPayCashOnDelivery"`
	DeliveryStaffName              string          `json:"deliveryStaffName"`
	DeliveryStaffPhoneNumber       string          `json:"deliveryStaffPhoneNumber"`
	DeliveryStaffBranchPhoneNumber string          `json:"deliveryStaffBranchPhoneNumber"`
	SenderCompany                  string          `json:"senderCompany"`
	SenderAddress                  string          `json:"senderAddress"`
	RecipientCompany               string          `json:"recipientCompany"`
	RecipientAddress               string          `json:"recipientAddress"`
	CourierTrackingNumber          string          `json:"courierTrackingNumber"`
	CourierPartner                 string          `json:"courierPartner"`
}

type Timeline struct {
	Date     		 string          `json:"date"`
	Details      []Event         `json:"details"`
}

type Event struct {
	DateTime    time.Time `json:"dateTime"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
}

func (client *Client) Find(courier string, trackingNo string) (*Response, error) {
	return client.track(courier, trackingNo, APIEndpointTracksFind)
}

func (client *Client) KerryExpress(trackingNo string) (*Response, error) {
	courier := "kerry-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ShopeeExpress(trackingNo string) (*Response, error) {
	courier := "shopee-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) DhlExpress(trackingNo string) (*Response, error) {
	courier := "dhl-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) JtExpress(trackingNo string) (*Response, error) {
	courier := "jt-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) AlphaFast(trackingNo string) (*Response, error) {
	courier := "alpha-fast"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BestExpress(trackingNo string) (*Response, error) {
	courier := "best-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) CjLogistics(trackingNo string) (*Response, error) {
	courier := "cj-logistics"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) SpeedD(trackingNo string) (*Response, error) {
	courier := "speed-d"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) NinjaVan(trackingNo string) (*Response, error) {
	courier := "ninja-van"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Ups(trackingNo string) (*Response, error) {
	courier := "ups"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ScgExpress(trackingNo string) (*Response, error) {
	courier := "scg-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BeeExpress(trackingNo string) (*Response, error) {
	courier := "bee-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) InterExpress(trackingNo string) (*Response, error) {
	courier := "inter-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) TntExpress(trackingNo string) (*Response, error) {
	courier := "tnt-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Shippop(trackingNo string) (*Response, error) {
	courier := "shippop"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Aramex(trackingNo string) (*Response, error) {
	courier := "aramex"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) FedEx(trackingNo string) (*Response, error) {
	courier := "fed-ex"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) TpLogistics(trackingNo string) (*Response, error) {
	courier := "tp-logistics"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ItTransport(trackingNo string) (*Response, error) {
	courier := "it-transport"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) NakhonchaiAir(trackingNo string) (*Response, error) {
	courier := "nakhonchai-air"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) NimExpress(trackingNo string) (*Response, error) {
	courier := "nim-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) FlashExpress(trackingNo string) (*Response, error) {
	courier := "flash-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BsExpress(trackingNo string) (*Response, error) {
	courier := "bs-express"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) Shipjung(trackingNo string) (*Response, error) {
	courier := "shipjung"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) GlobalCainiao(trackingNo string) (*Response, error) {
	courier := "global-cainiao"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) JapanPost(trackingNo string) (*Response, error) {
	courier := "japan-post"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) ChinaPost(trackingNo string) (*Response, error) {
	courier := "china-post"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) SkyBox(trackingNo string) (*Response, error) {
	courier := "sky-box"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) BookMyCargo(trackingNo string) (*Response, error) {
	courier := "book-my-cargo"
	return client.track(courier, trackingNo, fmt.Sprintf(APIEndpointTracks, courier))
}

func (client *Client) track(courier, trackingNo, endpoint string) (response *Response, err error) {
	payload := Payload{
		Courier: courier,
		TrackingNo: trackingNo,
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