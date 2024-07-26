package service

import (
	"context"

	"courier_delivery/config/logger"
	"courier_delivery/genproto"
	"courier_delivery/storage"
)

// TaskService vazifalar uchun gRPC xizmatini taqdim etadi
type TaskService struct {
	store storage.StorageInterface
	genproto.UnimplementedTaskServiceServer
	log logger.Logger
}

// NewTaskService yangi TaskService yaratadi
func NewTaskService(store storage.StorageInterface, log logger.Logger) *TaskService {
	return &TaskService{store: store, log: log}
}

// CreateTask RPC chaqiruvini bajaradi va yangi vazifa yozuvini yaratadi
func (s *TaskService) CreateTask(ctx context.Context, req *genproto.CreateTaskRequest) (*genproto.TaskResponse, error) {
	res, err := s.store.Task().CreateTask(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to create task: %v", err)
		return nil, err
	}
	s.log.INFO.Println("Successfully created task:", res.Id)
	return res, nil
}

// GetTask RPC chaqiruvini bajaradi va vazifa yozuvini qaytaradi
func (s *TaskService) GetTask(ctx context.Context, req *genproto.GetTaskRequest) (*genproto.TaskResponse, error) {
	res, err := s.store.Task().GetTask(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to get task with ID %s: %v", req.Id, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully retrieved task:", res.Id)
	return res, nil
}

// UpdateTask RPC chaqiruvini bajaradi va vazifa yozuvini yangilaydi
func (s *TaskService) UpdateTask(ctx context.Context, req *genproto.UpdateTaskRequest) (*genproto.TaskResponse, error) {
	res, err := s.store.Task().UpdateTask(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to update task with ID %s: %v", req.Id, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully updated task:", res.Id)
	return res, nil
}

// DeleteTask RPC chaqiruvini bajaradi va vazifa yozuvini o'chiradi
func (s *TaskService) DeleteTask(ctx context.Context, req *genproto.DeleteTaskRequest) (*genproto.DeleteTaskResponse, error) {
	res, err := s.store.Task().DeleteTask(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to delete task with ID %s: %v", req.Id, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully deleted task:", req.Id)
	return res, nil
}

// GetAllTasks RPC chaqiruvini bajaradi va berilgan foydalanuvchiga tayinlangan barcha vazifalarni qaytaradi
func (s *TaskService) GetAllTasks(ctx context.Context, req *genproto.GetAllTasksRequest) (*genproto.GetAllTasksResponse, error) {
	res, err := s.store.Task().GetAllTasks(ctx, req)
	if err != nil {
		s.log.ERROR.Printf("Failed to get all tasks for assigned_to %d: %v", req.AssignedTo, err)
		return nil, err
	}
	s.log.INFO.Println("Successfully retrieved all tasks for assigned_to:", req.AssignedTo)
	return res, nil
}
