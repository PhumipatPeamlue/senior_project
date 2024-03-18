package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type flexMessage struct {
	To       string `json:"to"`
	Messages []struct {
		Type     string `json:"type"`
		AltText  string `json:"altText"`
		Contents struct {
			Type string `json:"type"`
			Body struct {
				Type     string `json:"type"`
				Layout   string `json:"layout"`
				Contents []struct {
					Type     string `json:"type"`
					Text     string `json:"text,omitempty"`
					Weight   string `json:"weight,omitempty"`
					Size     string `json:"size,omitempty"`
					Layout   string `json:"layout,omitempty"`
					Margin   string `json:"margin,omitempty"`
					Spacing  string `json:"spacing,omitempty"`
					Contents []struct {
						Type     string `json:"type"`
						Layout   string `json:"layout"`
						Contents []struct {
							Type  string `json:"type"`
							Text  string `json:"text"`
							Color string `json:"color"`
							Flex  int    `json:"flex"`
							Size  string `json:"size"`
						} `json:"contents"`
						Spacing string `json:"spacing,omitempty"`
					} `json:"contents,omitempty"`
				} `json:"contents"`
			} `json:"body"`
			Footer struct {
				Type     string `json:"type"`
				Layout   string `json:"layout"`
				Spacing  string `json:"spacing"`
				Contents []struct {
					Type     string        `json:"type"`
					Layout   string        `json:"layout"`
					Contents []interface{} `json:"contents"`
					Margin   string        `json:"margin"`
				} `json:"contents"`
				Flex int `json:"flex"`
			} `json:"footer"`
		} `json:"contents"`
	} `json:"messages"`
}

func NewFlexMessage(lineUserID, petName, drugName, drugUsage string, notifyAt time.Time) (flexMsg flexMessage, err error) {
	file, err := os.Open("flex_message.json")
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&flexMsg)
	if err != nil {
		return
	}

	flexMsg.To = lineUserID
	flexMsg.Messages[0].AltText = fmt.Sprintf("แจ้งเตือนการใช้ยาของน้อง %s", petName)
	flexMsg.Messages[0].Contents.Body.Contents[1].Contents[0].Contents[1].Text = drugName
	flexMsg.Messages[0].Contents.Body.Contents[1].Contents[1].Contents[1].Text = notifyAt.String()
	flexMsg.Messages[0].Contents.Body.Contents[1].Contents[2].Contents[1].Text = petName
	flexMsg.Messages[0].Contents.Body.Contents[1].Contents[3].Contents[1].Text = drugUsage

	return
}
