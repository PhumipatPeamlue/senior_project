package handler

import "user_management_service/internal/line"

type Handler struct {
	lineClient line.LineClientInterface
}

func New(line line.LineClientInterface) *Handler {
	return &Handler{
		lineClient: line,
	}
}
