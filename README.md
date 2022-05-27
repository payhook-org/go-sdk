# Payhook SDK GO

go-sdk is an implementation of the Payhook API in Golang.

This version implement V1 Payhook API.

## Import

	import "github.com/payhook-org/go-sdk"

## Usage

~~~ go
package main

import (
	"fmt"
	"math/big"

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
}
~~~

## Examples
[Create With Webhook Example](https://github.com/payhook-org/go-sdk/blob/main/examples/main.go)

## Documentation
[Payhook Documentaion](https://docs.payhook.org)

## Licence

Copyright 2022 Payhook

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.