package repository

import (
	"srv-example/domain/model"
	"srv-example/global"
)

type IUserRepository interface {
	InitTable() error
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindUserByID(int64) (*model.User, error)
	FindAll() ([]model.User, error)
	FindPage() ([]model.User, error)
}

//创建UserRepository
func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

type UserRepository struct {
}

// 初始化表
func (u *UserRepository) InitTable() error {
	return global.DB.CreateTable(&model.User{}).Error
}

// 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, global.DB.Create(user).Error
}

// 根据用户ID删除用户
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return global.DB.Where("id = ?", userID).Delete(&model.User{}).Error
}

// 更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return global.DB.Model(user).Update(&user).Error
}

//根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, global.DB.First(user, userID).Error
}

func (u *UserRepository) FindAll() ([]model.User, error) {
	panic("implement me")
}

func (u *UserRepository) FindPage() ([]model.User, error) {
	panic("implement me")
}