package repository

import (
	"errors"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type TesterRepository struct {
	DB *gorm.DB
}




func (r *TesterRepository) Findall() ([]db.TestModel, error) {
	var test []db.TestModel
	err := r.DB.Find(&test).Error
	return test , err

}

func (r *TesterRepository) Create(d *domain.TestInput) (*db.TestModel, error) {
	test := db.TestModel{
       Name: d.Name,
	   Message: d.Message,
	}
	err := r.DB.Create(&test).Error
	return  &test, err
}

func (r *TesterRepository) Update(id uint, d *domain.TestInputUpdate ) (*db.TestModel , error) {
   u_r := &db.TestModel{}
   result :=  r.DB.Model(u_r).Where("id = ?", id).Updates(d)
   if result.Error != nil {
	 return nil , result.Error
   }
   if result.RowsAffected == 0 {
	return nil , errors.New("Query Object Not Found")
   }

   u_r.ID = id
   err := r.DB.Find(u_r).Error
   return u_r , err

}



func (r *TesterRepository)Delete(id uint) (error) {
  u_r := &db.TestModel{}
  result := r.DB.Model(u_r).Where("id= ?", id).Delete(u_r)

  if result.Error != nil {
	return result.Error
  }
  if result.RowsAffected == 0 {
	return  errors.New("Query Object Not Found")
  }
  return nil
}


func (r *TesterRepository) Get (id uint) (*db.TestModel , error) {
	u_r := &db.TestModel{
		ID: id,
	}
	result := r.DB.Find(u_r)
    if result.Error != nil {
	   return nil , result.Error
    }
   if result.RowsAffected == 0 {
	return nil , errors.New("Query Object Not Found")
   }
	return u_r , nil
}