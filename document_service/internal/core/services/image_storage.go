package services

import (
	"bytes"
	"document_service/internal/core/ports"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type GetImageUrlResponse struct {
	ImageURL string `json:"image_url"`
}

type imageStorageService struct {
	client http.Client
}

func (s *imageStorageService) GetImageURL(docID string) (imageURL string, err error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("http://image_storage_service:8080/image/%s", docID), nil)
	if err != nil {
		return
	}
	res, err := s.client.Do(request)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		imageURL = ""
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var getImageUrlResponse GetImageUrlResponse
	if err = json.Unmarshal(body, &getImageUrlResponse); err != nil {
		return
	}

	imageURL = getImageUrlResponse.ImageURL
	return
}

func (s *imageStorageService) SaveImage(docID string, file *multipart.File, header *multipart.FileHeader) (err error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	fileWriter, err := writer.CreateFormFile("image", header.Filename)
	if err != nil {
		return
	}

	_, err = io.Copy(fileWriter, *file)
	if err != nil {
		return
	}
	writer.Close()

	request, err := http.NewRequest("POST", fmt.Sprintf("http://image_storage_service:8080/image/%s", docID), &buffer)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := s.client.Do(request)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		err = fmt.Errorf("unexpected status code: %d", res.StatusCode)
		return
	}

	return
}

func (s *imageStorageService) DeleteImage(docID string) (err error) {
	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://image_storage_service:8080/image/%s", docID), nil)
	if err != nil {
		return
	}
	res, err := s.client.Do(request)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return
	} else if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", res.StatusCode)
		return
	}

	return
}

func NewImageStorageService(client http.Client) ports.ImageStorageService {
	return &imageStorageService{
		client: client,
	}
}
