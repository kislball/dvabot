package services

type AtomicResponse struct {
	Ok         bool   `json:"ok"`
	FailReason string `json:"fail_reason"`
}
