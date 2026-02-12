package services

import (
	"encoding/json"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	payload "github.com/marveldo/gogin/internal/application/payloads"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/db"
)

type TesterService struct {
	R *repository.TesterRepository
	C *asynq.Client
}

func (s *TesterService) Hello() string {
	return "Hello, world"
}

func (s *TesterService) Message() string {
	return "This is a message from TesterService"
}

func (s *TesterService) RunAddition(i *domain.Addition) error {
	payload := payload.AdditionPayload{}
	err := copier.Copy(&payload, i)
	if err != nil {
		return err
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	task := asynq.NewTask("add", b)
	s.C.Enqueue(task)
	return nil

}

func (s *TesterService) GetAllTests() ([]db.TestModel, error) {
	return s.R.Findall()
}

func (s *TesterService) CreateTest(d *domain.TestInput) (*db.TestModel, error) {
	return s.R.Create(d)
}

func (s *TesterService) UpdateTest(Id string, d *domain.TestInputUpdate) (*db.TestModel, error) {
	id, err := strconv.ParseUint(Id, 10, 32)
	if err != nil {
		return nil, err
	}
	return s.R.Update(uint(id), d)
}

func (s *TesterService) DeleteTest(Id string) error {
	id, err := strconv.ParseUint(Id, 10, 32)
	if err != nil {
		return err
	}
	return s.R.Delete(uint(id))
}

func (s *TesterService) GetTest(Id string) (*db.TestModel, error) {
	id, err := strconv.ParseUint(Id, 10, 32)
	if err != nil {
		return nil, err
	}
	return s.R.Get(uint(id))
}
