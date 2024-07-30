package handlers

import (
	"context"
	"net/http"
	"gateway/genproto/courier"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new task
// @Description Add a new task to the system
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param task body courier.CreateTaskRequest true "Task Data"
// @Success 201 {object} courier.TaskResponse
// @Failure 400 {object} string "Invalid Data"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/tasks/create [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var req courier.CreateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Task_C.CreateTask(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get task by ID
// @Description Retrieve details of a task by its ID
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param taskId path string true "Task ID"
// @Success 200 {object} courier.TaskResponse
// @Failure 400 {object} string "Invalid Task ID"
// @Failure 404 {object} string "Task Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/tasks/{taskId} [get]
func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("taskId")
	req := &courier.GetTaskRequest{Id: id}
	resp, err := h.Task_C.GetTask(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get all tasks
// @Description Retrieve a list of all tasks
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Success 200 {object} courier.GetAllTasksResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/tasks [get]
func (h *Handler) GetTasks(c *gin.Context) {
	req := &courier.GetAllTasksRequest{}
	resp, err := h.Task_C.GetAllTasks(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update task by ID
// @Description Update details of an existing task by its ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param taskId path string true "Task ID"
// @Param task body courier.UpdateTaskRequest true "Task Data"
// @Success 200 {object} courier.TaskResponse
// @Failure 400 {object} string "Invalid Task ID"
// @Failure 404 {object} string "Task Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/tasks/{taskId} [put]
func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("taskId")
	var req courier.UpdateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = id
	resp, err := h.Task_C.UpdateTask(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete task by ID
// @Description Remove a task from the system by its ID
// @Tags Tasks
// @Produce json
// @Security BearerAuth
// @Param taskId path string true "Task ID"
// @Success 200 {object} string "Task deleted successfully"
// @Failure 400 {object} string "Invalid Task ID"
// @Failure 404 {object} string "Task Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/tasks/{taskId} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("taskId")
	req := &courier.DeleteTaskRequest{Id: id}
	_, err := h.Task_C.DeleteTask(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
