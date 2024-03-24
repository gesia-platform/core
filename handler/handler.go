package handler

import (
	"github.com/gesia-platform/core/chaintree"
)

type Handler struct {
	chaintree *chaintree.ChainTree
}

func NewHandler(
	chaintree *chaintree.ChainTree,

) *Handler {
	return &Handler{chaintree}
}
