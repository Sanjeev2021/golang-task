package service

import (
	"errors"
	"log"
	"time"

	"task/database"
	"task/model"

)

var allTask []*model.Task

func NewTask(id uint, userid uint, dueDate time.Time, title string, description string, priority string, status string) (*model.Task, error) {

	user, err := GetUserById(userid)
	if err != nil {
		return nil, errors.New("cant find the user")
	}

	if user == nil {
		return nil, errors.New("user dosent exist")
	}

	if id <= 0 {
		return nil, errors.New("the id cant be less than 0 or 0")
	}

	if title == "" {
		return nil, errors.New("the title cant be empty")
	}

	if description == "" {
		return nil, errors.New("the description cant be empty")
	}

	if priority == "" {
		return nil, errors.New("the priority cant be empty , give it a priority")
	}

	if status == "" {
		return nil, errors.New("The status cant be empty")
	}

	if dueDate.Before(time.Unix(0, 0)) {
		return nil, errors.New("the due date cant be less than 0")
	}

	createdAt := time.Now().Truncate(24 * time.Hour)
	updatedAt := createdAt

	var tempTask = &model.Task{
		ID:          id,
		UserID:      userid,
		DueDate:     dueDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      status,
	}

	db := database.Get_GormDB()
	log.Println(db)
	if err := db.Create(tempTask).Error; err != nil {
		return nil, err
	}

	allTask = append(allTask, tempTask)
	return tempTask, nil

}

func DeleteTask(t *model.Task) error {
	if t.Title == "" {
		return errors.New("the task dosent exist")
	}

	t.ID = 0
	t.Title = ""

	return nil
}

func UpdateTaskTitle(t *model.Task, title string) error {
	if t.Title == "" {
		return errors.New("the task dosent exist")
	}

	if title == "" {
		return errors.New("the title cant be empty")
	}

	t.Title = title
	return nil
}

func GetAllTask() ([]*model.Task, error) {
	var task []*model.Task

	return task, nil

}
