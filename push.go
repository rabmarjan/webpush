package main

import (
	"bytes"
	"encoding/json"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
)

const (
	vapidPrivateKey = "9qFBi81Hi7Xw4ABn_HoCumXrQeIPW_TrNHunqRaP-Z4"
)

func main() {
	subJSON := `{"endpoint":"https://updates.push.services.mozilla.com/wpush/v2/gAAAAABbw_ghbcYp1JhCQk_Knc_yExbLRhZZqsrJ1Iv7svu9fvfJ-99eKkou1BLB3C2ii6fTGhcbXVVChtm7DEgo7Vp2UvXV9dz48OV3p-lOn8oimf_AiLlGTcfmSaVMBtWAFdPlkXEkjpXpyUA4UMqSnVhQGirmdSSz-k6Oq6pYss1r-YCtB0w","keys":{"auth":"QS6WY6oNDhw7XaTFYXEaUA","p256dh":"BJoGjSwXKXExiOzktBnsTI9TizMVuNVrTiNA09jKtnote-2Q-Uau2ohihkUkoQMBHPFAicnyR6u5XIiwYQjNoMg"}}`
	//sub := Push()
	//fmt.Println(string(sub))
	subJSONChrome := `{"endpoint":"https://fcm.googleapis.com/fcm/send/fPqQyt901IY:APA91bGzvC168YTBEivUAE4VEb4Vwcvg6Waz4kZBJ8NLWusB0vgcJb1Wh-BqKwNhaOqSazleJ5qwIxQiPjSMwhziJVdtB6CnNMbUvh9c4EWdSOIJP2WOAOdSgpqKCgij1Sbn8i6ctBIq","expirationTime":null,"keys":{"p256dh":"BCSvvrfLcQIEOMbdHeZFczQTzrvpP6BU-3m8KymGJZQpkEhizjsRUpBVvygEpQ7lDlkaejIMHBvyN7Z6iJbvG4k","auth":"jSn8lMApigroflK9WuB3Cg"}}`
	// Decode subscription
	s := webpush.Subscription{}
	if err := json.NewDecoder(bytes.NewBufferString(subJSONChrome)).Decode(&s); err != nil {
		log.Fatal(err)
	}
	m := webpush.Subscription{}
	if err := json.NewDecoder(bytes.NewBufferString(subJSON)).Decode(&m); err != nil {
		log.Fatal(err)
	}

	// Send Notification
	_, err := webpush.SendNotification([]byte("Test Marjan"), &s, &webpush.Options{
		Subscriber:      "rab.marjan@gmail.com",
		VAPIDPrivateKey: vapidPrivateKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	_, er := webpush.SendNotification([]byte("Test Sharif"), &m, &webpush.Options{
		Subscriber:      "rab.marjan@gmail.com",
		VAPIDPrivateKey: vapidPrivateKey,
	})
	if er != nil {
		log.Fatal(er)
	}
}
