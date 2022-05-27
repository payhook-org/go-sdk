package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	payhook "github.com/payhook-org/go-sdk"
)

const apiKey = "YOUR_API_KEY"

func main() {
	// payhook client
	client := payhook.New(apiKey)

	createParams := payhook.CreatePaymentParams{
		Title:    "Test Payment",
		Currency: "TON",
		Amount:   big.NewInt(1000000000).String(), // 1 TON in nanos
	}

	// create new payment
	payment, err := client.CreatePayment(createParams)
	if err != nil {
		panic(err)
	}

	fmt.Println("payment URL:", payment.URL)

	// create handler for webhook
	http.HandleFunc("/payhook/webhook", webhook)

	// start server on 10000 port
	err = http.ListenAndServe("0.0.0.0:10000", nil)
	panic(err)
}

func webhook(w http.ResponseWriter, r *http.Request) {
	webhookID := r.Header.Get("X-Payhook-Webhook-ID")
	webhookEvent := r.Header.Get("X-Payhook-Webhook-Event")
	webhookSignature := r.Header.Get("X-Payhook-Webhook-Signature")

	var body json.RawMessage
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	serverSignature := payhook.GenerateSignature(apiKey, webhookID, webhookEvent, string(body))
	if serverSignature != webhookSignature {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("webhook signature bad")
		return
	}

	var payment payhook.PaymentWebhook
	err = json.Unmarshal(body, &payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Println(payment.ID, payment.Status)

	w.WriteHeader(http.StatusOK)
}
