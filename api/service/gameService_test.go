package service_test

import (
	"database/sql"
	"github.com/420Nat20/Nat20/nat-20/data/model"
	"regexp"

	"github.com/420Nat20/Nat20/nat-20/service"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB        *gorm.DB
	mock      sqlmock.Sqlmock
	gameModel *model.GameModel
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	require.NoError(s.T(), err)
}

func (s *Suite) Test_GameService_Get() {
	var (
		serverId = 123456
		dm       = 1234
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "game_models" WHERE server_id = $1 AND "game_models"."deleted_at" IS NULL ORDER BY "game_models"."id" LIMIT 1`)).
		WithArgs(serverId).
		WillReturnRows(sqlmock.NewRows([]string{"dm"}).
			AddRow(dm))

	res, err := service.GetGameModelByServerID(s.DB, serverId)

	require.NoError(s.T(), err)
	require.True(s.T(), res.DM == dm)
}

func (s *Suite) Test_GameService_GetAll() {
	var (
		dm = 1234
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "game_models" WHERE "game_models"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"dm"}).
			AddRow(dm).
			AddRow(dm))

	res, err := service.GetAllGameModels(s.DB)

	require.NoError(s.T(), err)
	require.True(s.T(), len(res) == 2)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
