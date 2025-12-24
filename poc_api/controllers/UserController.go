package controllers

import (
	"github.com/gin-gonic/gin"
	interfaceCommand "poc_api/poc_application/interfaces"
	dtosRequest "poc_api/poc_application/dataTransferObjects/userDtos/request"
	dtosResponse "poc_api/poc_application/dataTransferObjects/userDtos/response"

)

type UserController struct {
	commandService interfaceCommand.IUserCommandService
	queryService interfaceCommand.IUserQueryService
}

func NewUserController(
	commandService interfaceCommand.IUserCommandService,
	queryService interfaceCommand.IUserQueryService,
) *UserController {
	return &UserController{
		commandService: commandService,
		queryService: queryService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var req dtosRequest.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := c.commandService.Register(&req)

	ctx.JSON(200, gin.H{
		"message": result,
	})
}

func (c *UserController) FindAll(ctx *gin.Context){
	result := c.queryService.FindAll()

	if len(result) == 0 {
		ctx.JSON(204, []dtosResponse.UserSimpleResponse{})
		return
	}

	ctx.JSON(200, result)
}
