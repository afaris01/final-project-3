package controllers

import (
	"final-project-3/helpers"
	"final-project-3/models/inputs"
	"final-project-3/models/responses"
	"final-project-3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *taskController {
	return &taskController{taskService}
}

func (h *taskController) CreateNewTask(c *gin.Context) {
	var input inputs.TaskCreateInput

	err := c.ShouldBindJSON(&input)
	currentUser := c.MustGet("currentUser").(int)

	if err != nil {
		error_message := gin.H{
			"error": helpers.FormatValidationError(err),
		}

		resp := helpers.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newTask, err := h.taskService.CreateTask(input, currentUser)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helpers.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success create new task
	taskResponse := responses.TaskCreatedResponse{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Description: input.Description,
		Status:      newTask.Status,
		UserID:      newTask.UserID,
		CategoryID:  input.CategoryID,
		CreatedAt:   newTask.CreatedAt,
	}

	resp := helpers.APIResponse("success", taskResponse)
	c.JSON(http.StatusCreated, resp)
}

func (h *taskController) GetAllTask(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	allTask, err := h.taskService.GetTasks(currentUser)

	if err != nil {
		error_message := gin.H{
			"error": helpers.FormatValidationError(err),
		}

		resp := helpers.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helpers.APIResponse("success", allTask)
	c.JSON(http.StatusOK, resp)
}

func (h *taskController) UpdateTaskStatus(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	var inputTask inputs.TaskUpdateStatus
	var inputID inputs.TaskID

	err := c.ShouldBindJSON(&inputTask)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check first authenticate user & task
	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// will lolos

	updatedTask, err := h.taskService.UpdateStatusTask(currentUser, inputTask)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := responses.TaskUpdatedResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserID:      updatedTask.UserID,
		CategoryID:  updatedTask.CategoryID,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	response := helpers.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)
}

func (h *taskController) UpdateTask(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	var inputTask inputs.TaskEditInput
	var inputID inputs.TaskID

	err := c.ShouldBindJSON(&inputTask)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check first authenticate user & task
	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedTask, err := h.taskService.UpdateTask(inputID.ID, inputTask)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := responses.TaskUpdatedResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserID:      updatedTask.UserID,
		CategoryID:  updatedTask.CategoryID,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	response := helpers.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)
}

func (h *taskController) UpdateStatusTask(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputTaskStatus inputs.TaskUpdateStatus
	var inputID inputs.TaskID

	err := c.ShouldBindJSON(&inputTaskStatus)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedStatusTask, err := h.taskService.UpdateStatusTask(inputID.ID, inputTaskStatus)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := responses.TaskUpdatedResponse{
		ID:          updatedStatusTask.ID,
		Title:       updatedStatusTask.Title,
		Description: updatedStatusTask.Description,
		Status:      updatedStatusTask.Status,
		UserID:      updatedStatusTask.UserID,
		CategoryID:  updatedStatusTask.CategoryID,
		UpdatedAt:   updatedStatusTask.UpdatedAt,
	}

	response := helpers.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)

}

func (h *taskController) UpdateCategoryTask(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputCategory inputs.TaskUpdateCategory
	var inputID inputs.TaskID

	err := c.ShouldBindJSON(&inputCategory)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedStatusTask, err := h.taskService.UpdateStatusCategoryTask(inputID.ID, inputCategory)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := responses.TaskUpdatedResponse{
		ID:          updatedStatusTask.ID,
		Title:       updatedStatusTask.Title,
		Description: updatedStatusTask.Description,
		Status:      updatedStatusTask.Status,
		UserID:      updatedStatusTask.UserID,
		CategoryID:  updatedStatusTask.CategoryID,
		UpdatedAt:   updatedStatusTask.UpdatedAt,
	}

	response := helpers.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)

}

func (h *taskController) DeleteTask(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	var taskDeletedInput inputs.TaskID
	err := c.ShouldBindUri(&taskDeletedInput)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskWillDeleted, err := h.taskService.GetTaskDetail(taskDeletedInput.ID)

	if taskWillDeleted.ID == 0 {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillDeleted.UserID != currentUser {
		response := helpers.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.taskService.DeleteTask(taskDeletedInput.ID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := "Task has been successfully deleted"

	response := helpers.APIResponse("ok", message)
	c.JSON(http.StatusOK, response)
}