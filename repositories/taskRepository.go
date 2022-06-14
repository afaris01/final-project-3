package repositories

import (
	"final-project-3/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Save(task models.Task) (models.Task, error)
	Get(IDUser int) ([]models.Task, error)
	GetDetail(ID int) (models.Task, error)
	Update(ID int, taskEdit models.Task) (models.Task, error)
	SwitchStatus(ID int, status bool) (models.Task, error)
	Delete(ID int) (bool, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Save(task models.Task) (models.Task, error) {
	err := r.db.Create(&task).Error

	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Get(IDUser int) ([]models.Task, error) {
	allTask := []models.Task{}

	err := r.db.Where("user_id = ?", IDUser).Find(&allTask).Error

	if err != nil {
		return []models.Task{}, err
	}

	return allTask, nil

}

func (r *taskRepository) GetDetail(ID int) (models.Task, error) {
	task := models.Task{}

	err := r.db.Where("id = ?", ID).Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *taskRepository) Update(ID int, taskEdit models.Task) (models.Task, error) {
	err := r.db.Where("id = ?", ID).Updates(taskEdit).Error

	if err != nil {
		return models.Task{}, err
	}

	updatedTask := models.Task{}
	err = r.db.Where("id = ?", ID).Find(&updatedTask).Error

	if err != nil {
		return models.Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepository) SwitchStatus(ID int, status bool) (models.Task, error) {
	err := r.db.Where("id = ?", ID).Updates(models.Task{Status: status}).Error

	if err != nil {
		return models.Task{}, err
	}

	taskSwitched := models.Task{}
	err = r.db.Where("id = ?", ID).Find(&taskSwitched).Error

	if err != nil {
		return models.Task{}, err
	}

	return taskSwitched, nil
}

func (r *taskRepository) Delete(ID int) (bool, error) {

	taskDeleted := models.Task{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&taskDeleted).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
