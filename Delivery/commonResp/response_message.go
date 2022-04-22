package commonResp

var (
	statusSuccess = "Success"
)

type SuccessMessage struct {
	ResponseSuccess bool        `json:"success"`
	StatusMessage   string      `json:"message"`
	Data            interface{} `json:"data"`
}

type SuccessMessage2 struct {
	ResponseSuccess bool   `json:"success"`
	StatusMessage   string `json:"message"`
}

type FailedMessage struct {
	ResponseFailed bool   `json:"success"`
	StatusMessage  string `json:"message"`
}

func NewSuccessMessage(data interface{}) *SuccessMessage {
	return &SuccessMessage{
		ResponseSuccess: true,
		StatusMessage:   statusSuccess,
		Data:            data,
	}
}

func NewSuccessMessage2() *SuccessMessage2 {
	return &SuccessMessage2{
		ResponseSuccess: true,
		StatusMessage:   statusSuccess,
	}
}

func NewFailedMessage(message string) *FailedMessage {
	return &FailedMessage{
		ResponseFailed: false,
		StatusMessage:  message,
	}
}
