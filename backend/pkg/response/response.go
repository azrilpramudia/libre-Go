package response

type Envelope struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(data interface{}) Envelope {
	return Envelope{Success: true, Data: data}
}

func Error(message string) Envelope {
	return Envelope{Success: false, Error: message}
}
