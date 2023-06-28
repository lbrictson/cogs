package pkg

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/lbrictson/cogs/ent"
	_ "modernc.org/sqlite"
	"os"
)

type NewDatabaseConnectionInput struct {
	InMemory bool
	Location string
}

func NewDatabaseConnection(input NewDatabaseConnectionInput) (*ent.Client, error) {
	params := "_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)&_pragma=synchronous(NORMAL)&_pragma=journal_size_limit(100000000)"
	if input.InMemory {
		db, err := sql.Open("sqlite", "file:test.db?mode=memory&_pragma=foreign_keys(1)")
		if err != nil {
			return nil, err
		}
		drv := entsql.OpenDB(dialect.SQLite, db)
		conn := ent.NewClient(ent.Driver(drv))
		return conn, migrate(conn)
	}
	// Grab the last character of the location
	lastChar := input.Location[len(input.Location)-1:]
	// Add a slash if it's not there because this is the path to a directory
	if lastChar != "/" {
		input.Location = fmt.Sprintf("%s/", input.Location)
	}
	err := validatePathToDBLocationExists(input.Location)
	if err != nil {
		return nil, err
	}
	loc := input.Location + "data.db"
	db, err := sql.Open("sqlite", fmt.Sprintf("%v?%v", loc, params))
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.SQLite, db)
	conn := ent.NewClient(ent.Driver(drv))
	return conn, migrate(conn)
}

// migrate runs all database migrations
func migrate(conn *ent.Client) error {
	return conn.Schema.Create(context.Background())
}

// validatePathToDBLocationExists will create the path to the requested database location if it does not exist
func validatePathToDBLocationExists(path string) error {
	// Check if the path exists
	// If it does not, create it
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}
