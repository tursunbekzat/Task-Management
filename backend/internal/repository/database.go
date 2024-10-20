package repository

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "backend/internal/models"
)

type Repository struct {
    db *gorm.DB
}

func (r *Repository) DB() *gorm.DB {
    return r.db
}

func NewRepository(dsn string) (*Repository, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Auto migrate schemas
    err = db.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        return nil, err
    }

    return &Repository{db: db}, nil
}

func (r *Repository) CreateTask(task *models.Task) error {
    return r.db.Create(task).Error
}

func (r *Repository) GetTaskByID(id uint) (*models.Task, error) {
    var task models.Task
    err := r.db.Preload("User").First(&task, id).Error
    return &task, err
}

func (r *Repository) UpdateTask(task *models.Task) error {
    return r.db.Save(task).Error
}

func (r *Repository) DeleteTask(id uint) error {
    return r.db.Delete(&models.Task{}, id).Error
}

func (r *Repository) ListTasks(userID uint) ([]models.Task, error) {
    var tasks []models.Task
    err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
    return tasks, err
}





type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
    return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
    return r.db.Delete(&models.User{}, id).Error
}