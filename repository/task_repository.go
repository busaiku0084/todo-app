package repository

import (
	"fmt"
	"todo-app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// tasksRepository は ITasksRepository を実装する構造体
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository コンストラクタ
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// GetAllTasks タスクを全件取得する
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// GetTaskById タスクを1件取得する
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

// CreateTask タスクを作成する
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTask タスクを更新する
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

// DeleteTask タスクを削除する
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}
