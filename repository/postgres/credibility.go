package postgres

import (
	"time"

	"github.com/BranDebs/Legitly-Backend/credibility"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type credibilityRepository struct {
	client  *gorm.DB
	db      string
	timeout time.Duration
}

func NewRepository(connStr string, db string, timeout time.Duration) credibility.Repository {
	return &credibilityRepository{
		db:      db,
		timeout: timeout,
	}
}

func (c *credibilityRepository) Add(*credibility.Credibility) error {
	return nil
}

func (c *credibilityRepository) Get(URL string) (*credibility.Credibility, error) {
	return nil, nil
}

func (c *credibilityRepository) Update(*credibility.Credibility) (*credibility.Credibility, error) {
	return nil, nil
}
