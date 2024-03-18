package core

import (
	"bytes"
	"context"
	"cronjob/internal/core/domains"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type LineNotificationServiceInterface interface {
	PushMessage(ctx context.Context, lineUserID, petName string, drugInfo domains.DrugInfo, notifyAt time.Time) (err error)
}

type lineNotificationService struct {
	httpClient *http.Client
}

func (l *lineNotificationService) PushMessage(ctx context.Context, lineUserID, petName string, drugInfo domains.DrugInfo, notifyAt time.Time) (err error) {
	drugName, drugUsage := drugInfo.DrugName, drugInfo.DrugUsage

	flexMsg, err := domains.NewFlexMessage(lineUserID, petName, drugName, drugUsage, notifyAt)
	if err != nil {
		return
	}

	b, err := json.Marshal(flexMsg)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", os.Getenv("LINE_MESSAGE_PUSH_URL"), bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("LINE_ACCESS_TOKEN")))

	resp, err := l.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
	}

	return
}

func NewLineNotificationService(httpClient *http.Client) LineNotificationServiceInterface {
	return &lineNotificationService{
		httpClient: httpClient,
	}
}
