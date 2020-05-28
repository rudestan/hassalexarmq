package alexakit

import "encoding/json"

type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Body    ReqBody `json:"request"`
	Context Context `json:"context"`
}

type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

type Context struct {
	System struct {
		APIAccessToken string `json:"apiAccessToken"`
		Device         struct {
			DeviceID string `json:"deviceId,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
		ApiEndpoint string `json:"apiEndpoint"`
		ApiAccessToken string `json:"apiAccessToken"`
	} `json:"System,omitempty"`
}

type ReqBody struct {
	Type      string `json:"type"`
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
	Intent    Intent `json:"intent,omitempty"`
	Reason    string `json:"reason,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
}

type Intent struct {
	Name  string          `json:"name"`
	Slots map[string]Slot `json:"slots"`
}

type Slot struct {
	ConfirmationStatus string `json:"confirmationStatus"`
	Name  			  string `json:"name"`
	Value 			  string `json:"value"`
	Resolutions  Resolutions `json:"resolutions"`
	Source			  string `json:"source"`
}

type Resolutions struct {
	ResolutionPerAuthority []struct{
		Authority string `json:"authority"`
		Status struct{
			Code string `json:"code"`
		} `json:"status"`
		Values []struct{
			Value struct{
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}

func (r *Request) ToJson() (string, error) {
	content, err := json.Marshal(r)

	if err != nil {
		return "", err
	}

	return string(content), nil
}
