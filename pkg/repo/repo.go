package repo

import (
	"errors"
	"log"
	"xmgo/pkg/apperror"
	"xmgo/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repo struct {
	dbh *gorm.DB
}

type IRepository interface {
	GetCompany(ID uint) (models.Company, error)
	CreateCompany(rec *models.Company) error
}

func NewRepo(dsn string) (*Repo, error) {
	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	return &Repo{
		dbh: db,
	}, err
}

func (r *Repo) AutoMigrate(c models.Company) {
	r.dbh.AutoMigrate(&c)
}

func (r *Repo) GetCompany(ID uint) (models.Company, error) {
	company := models.Company{}

	err := r.dbh.First(&company, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return company, apperror.ErrNotFound
	} else {
		log.Println(err)
	}
	return company, err
}

func (r *Repo) CreateCompany(rec *models.Company) error {
	result := r.dbh.Create(&rec)
	return result.Error
}
