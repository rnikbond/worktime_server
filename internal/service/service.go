package service

import (
	"context"
	"fmt"

	tasksPkg "worktime_server/internal/service/tasks"
	tickerPkg "worktime_server/internal/service/tasks/tickerTask"
	userPkg "worktime_server/internal/user"
)

type Service struct {
	tasks map[*userPkg.User]tasksPkg.Tasks

	ctx context.Context
}

func NewService(ctx context.Context) *Service {

	return &Service{
		tasks: make(map[*userPkg.User]tasksPkg.Tasks),
		ctx:   ctx,
	}
}

// Login Авторизация пользователя
func (s *Service) Login(username string) error {

	if _, err := s.FindUser(username); err != nil {
		newUser := userPkg.CreateUser(s.ctx, username, 0)

		tasks := tasksPkg.Tasks{
			ITickerTask: tickerPkg.NewTask(newUser),
		}

		s.tasks[newUser] = tasks
	}

	return nil
}

// RunTasks Запуск задач пользователя
func (s *Service) RunTasks(username string) error {

	user, err := s.FindUser(username)
	if err != nil {
		return err
	}

	fmt.Println("run task")
	s.tasks[user].ITickerTask.Start()

	return nil
}

// StopTasks Остановка задач пользователя
func (s *Service) StopTasks(username string) error {

	user, err := s.FindUser(username)
	if err != nil {
		return err
	}

	fmt.Println("stop task")
	s.tasks[user].ITickerTask.Stop()

	return nil
}

// FindUser Поиск авторизованного пользователя
func (s *Service) FindUser(username string) (*userPkg.User, error) {

	for lookUser := range s.tasks {
		if lookUser.Username == username {
			return lookUser, nil
		}
	}

	return nil, ErrUserUnauthorized
}
