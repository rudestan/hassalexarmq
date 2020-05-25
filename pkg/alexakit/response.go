package alexakit

const (
	version = "1.0"
	OutputSpeechTypePlainText = "PlainText"
)

const (
	SpeechTextConfirmation = "Ok."
	SpeechTextFailed = "Operation failed."
)

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type RespBody struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}

type Response struct {
	Version string `json:"version"`
	Response RespBody `json:"response"`
}

func NewPlainTextSpeechResponse(speechText string) Response {
	return Response{
		Version:  version,
		Response: RespBody{
			OutputSpeech: OutputSpeech{
				Type: OutputSpeechTypePlainText,
				Text: speechText,
			},
		},
	}
}
