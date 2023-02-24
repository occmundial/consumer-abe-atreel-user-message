package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

var (
	dbTimeout time.Duration
)

func NewQueries(configuration *models.Configuration, db *SQLServerConnection) *Queries {
	cs := Queries{Configuration: configuration, Database: db.Connection}
	cs.init()
	return &cs
}

type Queries struct {
	Configuration *models.Configuration
	Database      *sql.DB
}

func (queries *Queries) init() {
	dbTimeout = time.Duration(queries.Configuration.DBTimeout) * time.Second
}

func (queries *Queries) GetDicState() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT StateID,StateName FROM dbo.LookupStates (NOLOCK) ls WHERE ls.CountryID = 'MX'`
	rows, err := queries.Database.QueryContext(ctx, query)
	if err != nil {
		return make(map[string]string), err
	}
	defer rows.Close()
	var states = make(map[string]string)
	for rows.Next() {
		var state = models.State{}
		if err = rows.Scan(&state.StateID, &state.StateName); err != nil {
			return states, err
		}
		states[state.StateID] = state.StateName
	}
	return states, nil
}

func (queries *Queries) CheckHealth() error {
	ctxPing, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	return queries.Database.PingContext(ctxPing)
}
