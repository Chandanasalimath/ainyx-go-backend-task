package service

import (
	"time"

	"ainyx-go-backend-task/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *repository.Queries
}

func NewUserService(repo *repository.Queries) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name string, dob time.Time) (repository.User, error) {
	return s.repo.CreateUser(nil, repository.CreateUserParams{
		Name: name,
		Dob: pgtype.Date{
			Time:  dob,
			Valid: true,
		},
	})
}

func (s *UserService) GetUserByID(id int32) (repository.User, int, error) {
	user, err := s.repo.GetUserByID(nil, id)
	if err != nil {
		return repository.User{}, 0, err
	}

	age := calculateAge(user.Dob.Time)
	return user, age, nil
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
