package usecase

import (
	"todo-app/model"
	"todo-app/repository"
	"todo-app/validator"
)

// ITaskUsecase インターフェース
type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

// taskUsecase ITaskUsecase を実装する構造体
type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

// NewTaskUsecase コンストラクタ
func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

// GetAllTasks タスクを全件取得する
func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	var tasks []model.Task
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return []model.TaskResponse{}, err
	}

	var resTasks []model.TaskResponse
	for _, task := range tasks {
		t := model.TaskResponse{
			Id:        task.Id,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

// GetTaskById タスクを1件取得する
func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

// CreateTask タスクを作成する
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.ValidateTask(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

// UpdateTask タスクを更新する
func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tv.ValidateTask(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

// DeleteTask タスクを削除する
func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
