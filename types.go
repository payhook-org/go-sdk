package payhook

const (
	PendingStatus    PaymentStatus = "pending"
	SuccessfulStatus PaymentStatus = "successful"
	FailedStatus     PaymentStatus = "failed"
)

// Response wraps the JSON response
type Response struct {
	// will return if it contains an error
	Message *string `json:"message"`
	// response result
	Result interface{} `json:"result"`
}

// PaymentWebhook send on webhook addr
type PaymentWebhook struct {
	Payment `json:"payment"`
}

type PaymentMethod struct {
	// name of the method
	Name string `json:"name"`
	// currency of the method
	Currency string `json:"currency"`
	// qr code url as png
	QRCodeURLPNG string `json:"qr_code_url_png"`
	// qr code url as svg
	QRCodeURLSVG string `json:"qr_code_url_svg"`
	// method url
	URL string `json:"URL"`
}

// Payment contains returned payment information
type Payment struct {
	// id of the payment
	ID uint64 `json:"id"`
	// hash id of the payment
	HashID string `json:"hash_id"`
	// payment creation time (unix)
	CreatedAt uint64 `json:"created_at"`
	// payment expiration time (unix)
	ExpiresAt uint64 `json:"expires_at"`
	// payment time (unix)
	PaidAt *uint64 `json:"paid_at"`
	// testnet payment
	Test bool `json:"test"`
	// title of the payment
	Title string `json:"title"`
	// description of the payment (empty if not provided)
	Description string `json:"description"`
	// currency of the payment (USD, TON etc)
	Currency string `json:"currency"`
	// amount of the payment (in nanos)
	Amount string `json:"amount"`
	// status of the payment (pending, successful, failed)
	Status PaymentStatus `json:"status"`
	// reason of failure (empty if not failed)
	FailReason string `json:"fail_reason"`
	// payload of the payment (empty if not provided)
	Payload string `json:"payload"`
	// comment of the payment
	Comment string `json:"comment"`
	// selected payment currency
	ToPayCurrency *string `json:"to_pay_currency"`
	// selected payment currency amount (in nanos)
	ToPayAmount *string `json:"to_pay_amount"`
	// elected payment time (unix)
	ToPayDate *uint64 `json:"to_pay_date"`
	// payment exchange rate (in nanos)
	ExchangeRate *string `json:"exchange_rate"`
	// sender wallet address
	SenderAddress *string `json:"sender_address"`
	// receiver wallet address
	ReceiverAddress *string `json:"receiver_address"`
	// url to redirect after payment (successful or failed)
	RedirectURL *string `json:"redirect_url"`
	// checkout url
	URL string `json:"url"`
	// payment transactions (separated by coma)
	BlockchainTransactions *string `json:"blockchain_transactions"`
	// payment transactions
	Methods []PaymentMethod `json:"methods"`
}

// PaymentStatus status of the payment (PendingStatus, SuccessfulStatus, FailedStatus)
type PaymentStatus string

// CreatePaymentParams params to the creation payment
type CreatePaymentParams struct {
	// title of the payment
	Title string `json:"title"`
	// description of the payment, not required
	Description string `json:"description,omitempty"`
	// currency of the payment (USD, TON etc)
	Currency string `json:"currency"`
	// amount of the payment (in nanos)
	Amount string `json:"amount"`
	// timeout in minutes of the payment, not required (default 15)
	Timeout uint64 `json:"timeout,omitempty"`
	// payload of the payment (max length 2048), not required
	Payload string `json:"payload,omitempty"`
	// url to redirect after payment (successful or failed), not required
	RedirectURL string `json:"redirect_url,omitempty"`
	// use testnet, not required
	Test bool `json:"test,omitempty"`
}
