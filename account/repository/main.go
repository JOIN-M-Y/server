package repository

import (
	"github.com/JOIN-M-Y/server/account/entity"
	"github.com/jinzhu/gorm"
)

// Interface repository inteface
type Interface interface {
	Create(entity *entity.Account) error
	Update(entity *entity.Account) error
	Delete(accountID string) error
	FindByID(accountID string, unscoped bool) (*entity.Account, error)
	FindByEmailAndProvider(email string, provider string, unscoped bool) (*[]entity.Account, error)
	FindByEmail(email string) (*[]entity.Account, error)
}

// Repository repository for query to database
type Repository struct{ connection *gorm.DB }

// New create repository instance
func New(connection *gorm.DB) Interface {
	return &Repository{connection: connection}
}

// Create create account
func (repository *Repository) Create(entity *entity.Account) error {
	return repository.connection.Create(entity).Error
}

// Update update account
func (repository *Repository) Update(entity *entity.Account) error {
	return repository.connection.Save(entity).Error
}

// Delete delete account by accountId
func (repository *Repository) Delete(accountID string) error {
	condition := entity.Account{ID: accountID}
	return repository.connection.Delete(condition).Error
}

// FindByID find account by accountId
func (repository *Repository) FindByID(accountID string, unscoped bool) (*entity.Account, error) {
	output := entity.Account{}
	condition := entity.Account{ID: accountID}
	if unscoped == true {
		err := repository.connection.Unscoped().Where(condition).First(&output).Error
		return &output, err
	}
	err := repository.connection.Unscoped().Where(condition).First(&output).Error
	return &output, err
}

// FindByEmailAndProvider find all account
func (repository *Repository) FindByEmailAndProvider(email string, provider string, unscoped bool) (*[]entity.Account, error) {
	output := []entity.Account{}
	condition := entity.Account{Email: email, Provider: provider}
	if unscoped == true {
		err := repository.connection.Unscoped().Where(condition).Find(output).Error
		return &output, err
	}
	err := repository.connection.Where(condition).Find(&output).Error
	return &output, err
}

// FindByEmail find account by email
func (repository *Repository) FindByEmail(email string) (*[]entity.Account, error) {
	output := []entity.Account{}
	condition := entity.Account{Email: email}
	err := repository.connection.Where(condition).Find(&output).Error
	return &output, err
}
