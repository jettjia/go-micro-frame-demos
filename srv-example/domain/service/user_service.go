package service

import (
	"srv-example/domain/model"
	"srv-example/domain/repository"
)

type IUserService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(*model.User) error
	FindUserByID(int64) (*model.User, error)
	FindAllUser() ([]model.User, error)
	FindPage() ([]model.User, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository}
}

type UserService struct {
	UserRepository repository.IUserRepository
}

// 插入
func (u *UserService) AddUser(user *model.User) (int64, error) {
	return u.UserRepository.CreateUser(user)
}

// 删除
func (u *UserService) DeleteUser(ID int64) error {
	return u.UserRepository.DeleteUserByID(ID)
}

// 修改
func (u *UserService) UpdateUser(user *model.User) error {
	return u.UserRepository.UpdateUser(user)
}

// 通过ID 查找
func (u *UserService) FindUserByID(ID int64) (*model.User, error) {
	return u.UserRepository.FindUserByID(ID)
}

// 所有
func (u *UserService) FindAllUser() ([]model.User, error) {
	panic("implement me")
}

// 分页
func (u *UserService) FindPage() ([]model.User, error) {
	panic("implement me")
}
