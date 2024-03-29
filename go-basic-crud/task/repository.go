package task

import "gorm.io/gorm"

type Repository interface {
	Insert(task Task) (Task, error)
	Find(id string) (Task, error)
	Index() ([]Task, error)
	Update(id string, task Task) (Task, error)
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(task Task) (Task, error) {
	err := r.db.Create(&task).Error

	if err != nil {
		return task, err
	}
	return task, nil

}

func (r *repository) Index() ([]Task, error) {

	var tasks []Task
	err := r.db.Find(&tasks).Error

	if err != nil {
		return tasks, err

	}

	return tasks, nil

}

func (r *repository) Find(id string) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error

	if err != nil {
		return task, err

	}

	return task, nil
}

func (r *repository) Delete(id string) error {
	task, err := r.Find(id)

	if err != nil {
		return err
	}

	err = r.db.Delete(&task).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(id string, newTask Task) (Task, error) {
	task, err := r.Find(id)

	if err != nil {
		return task, err
	}

	r.db.Model(&task).Updates(newTask)

	return task, nil
}
