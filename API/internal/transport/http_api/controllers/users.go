package controllers

import (
	"api/internal/transport/grpc_api"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	GRPCClient grpc_api.GRPCClient
}

func NewUserHandler(grpc *grpc_api.GRPCClient) *UserHandler {
	return &UserHandler{
		GRPCClient: *grpc,
	}
}

func (h *UserHandler) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	log.Printf("ID : %v\nError : %v\n", id, err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	user, err := h.GRPCClient.GetUserByID(int32(id))
	if err != nil {
		log.Printf("Error GRPC: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) Put(ctx *gin.Context) {
	//TODO
}

func (h *UserHandler) Post(ctx *gin.Context) {
	//TODO
}

func (h *UserHandler) Del(ctx *gin.Context) {
	//TODO
}
