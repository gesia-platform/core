package handler

import (
	"github.com/gesia-platform/core/apps/notary/chainclient"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	chainClient *chainclient.ChainClient
	logger      *logrus.Logger
}

func NewHandler(chainClient *chainclient.ChainClient, logger *logrus.Logger) *Handler {
	return &Handler{chainClient, logger}
}
