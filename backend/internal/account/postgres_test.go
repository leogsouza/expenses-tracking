package account

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/expenses-tracking/backend/internal/entity"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	account    *entity.Account
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.repository, err = NewRepository(s.DB)
	require.NoError(s.T(), err)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestFind() {
	var (
		id   = entity.GenerateID()
		name = "Wallet"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT * FROM "accounts" WHERE (id = $1)`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(id.String(), name))

	res, err := s.repository.Find(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(entity.Account{ID: id, Name: name}, res))
}

func (s *Suite) TestStore() {
	var (
		account = &entity.Account{
			ID:        entity.GenerateID(),
			Name:      "Wallet",
			CreatedAt: time.Now().UTC(),
		}
	)

	s.mock.ExpectBegin() // begin transaction
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "accounts" ("id","name","created_at") VALUES ($1,$2,$3) RETURNING "accounts"."id"`)).
		WithArgs(account.ID, account.Name, account.CreatedAt).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(account.ID.String()))

	s.mock.ExpectCommit()
	id, err := s.repository.Store(account)
	require.NoError(s.T(), err)
	require.Equal(s.T(), account.ID, id)
}
