package services

import (
	"strconv"

	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/db"
)

type TesterService struct {
	R *repository.TesterRepository
}

func (s *TesterService) Hello() string {
	return "Hello, world"
}

func (s *TesterService) Message() string {
	return "This is a message from TesterService"
}

func (s *TesterService) GetAllTests() ([]db.TestModel, error) {
	return s.R.Findall()
}

func (s *TesterService) CreateTest(d *domain.TestInput) (*db.TestModel , error) {
	return s.R.Create(d)
}

func (s *TesterService) UpdateTest(Id string,d *domain.TestInputUpdate) (*db.TestModel , error){
    id , err:= strconv.ParseUint(Id, 10, 32)
	if err != nil {
		return nil , err
	}
	return s.R.Update(uint(id), d)
}