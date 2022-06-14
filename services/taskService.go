package services

import (
	"final-project-3/models"
	"final-project-3/models/inputs"
	"final-project-3/repositories"
)

type TaskService interface {
	CreateTask(input inputs.TaskCreateInput, IDUser int) (models.Task, error)
	GetTasks(ID int) ([]models.Task, error)
	GetTaskDetail(ID int) (models.Task, error)
	UpdateTask(ID int, editTask inputs.TaskEditInput) (models.Task, error)
	UpdateStatusTask(ID int, statusTask inputs.TaskUpdateStatus) (models.Task, error)
	DeleteTask(ID int) (bool, error)
	UpdateStatusCategoryTask(ID int, statusTask inputs.TaskUpdateCategory) (models.Task, error)
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(taskRepository repositories.TaskRepository) *taskService {
	return &taskService{taskRepository}
}

func (s *taskService) CreateTask(input inputs.TaskCreateInput, IDUser int) (models.Task, error) {
	newTask := models.Task{
		Description: input.Description,
		Title:       input.Title,
		CategoryID: input.CategoryID,
		UserID:      IDUser,
		Status:      false,
	}

	createdTask, err := s.taskRepository.Save(newTask)

	if err != nil {
		return createdTask, err
	}

	return createdTask, nil
}

func (s *taskService) GetTasks(ID int) ([]models.Task, error) {
	tasks, err := s.taskRepository.Get(ID)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *taskService) GetTaskDetail(ID int) (models.Task, error) {
	task, err := s.taskRepository.GetDetail(ID)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *taskService) UpdateTask(ID int, editTask inputs.TaskEditInput) (models.Task, error) {
	// mapping struct
	updateTask := models.Task{
		Title:       editTask.Title,
		Description: editTask.Description,
	}

	tasks, err := s.taskRepository.Update(ID, updateTask)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *taskService) UpdateStatusTask(ID int, statusTask inputs.TaskUpdateStatus) (models.Task, error) {
	// mapping task
	newStatusTask := statusTask.Status

	taskUpdatedStatus, err := s.taskRepository.SwitchStatus(ID, newStatusTask)

	if err != nil {
		return taskUpdatedStatus, err
	}

	return taskUpdatedStatus, nil
}

func (s *taskService) UpdateStatusCategoryTask(ID int, statusTask inputs.TaskUpdateCategory) (models.Task, error) {
	// mapping task
	newCategoryTask := models.Task{
		CategoryID: statusTask.CategoryID,
	}

	taskUpdated, err := s.taskRepository.Update(ID, newCategoryTask)

	if err != nil {
		return taskUpdated, err
	}

	return taskUpdated, nil
}

func (s *taskService) DeleteTask(ID int) (bool, error) {
	deletedTask, err := s.taskRepository.Delete(ID)

	if err != nil {
		return deletedTask, err
	}

	return deletedTask, nil
}
