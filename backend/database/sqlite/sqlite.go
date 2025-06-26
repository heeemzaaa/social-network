package sqlite

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

// the struct that will hold the connection !!

type Database struct {
	Database *sql.DB
}

func InitDB(nameDB string) (*Database, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%v?_foreign_keys=1", nameDB))
	if err != nil {
		return nil, err
	}
	database := &Database{db}
	errMigr := database.RunMigrations("database/migrations/")
	if errMigr != nil {
		return nil, err
	}
	return database, nil
}

func (db *Database) RunMigrations(migrationsPath string) error {
	// creating the driver
	// it ties the existing conn with the migration system that runs .up.sql and .down.sql
	driver, err := sqlite3.WithInstance(db.Database, &sqlite3.Config{})
	if err != nil {
		log.Fatal("Error! While trying to set up the driver with the connected Database!", err)
	}
	// driver = wrapped database connection
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%v", migrationsPath), "sqlite3", driver)
	// m = migration manager object
	if err != nil {
		log.Fatal("Error! While setting up the migration manager object", err)
	}
	// check if there is empty up files

	ok, emptyFiles, errEmpty := HasEmptyFilesUp(migrationsPath)
	if errEmpty != nil {
		log.Fatal("Error! While trying to check the directory of the migrations!", errEmpty)
	}

	if ok {
		for _, file := range emptyFiles {
			fmt.Printf("Empty file: %v\n", file)
		}
		os.Exit(1)
	}

	// let's treat the case of a dirty database version and rollback to the previous version
	if err := m.Up(); err != nil {
		if IsDirtyErr(err) {
			version, dirty, _ := m.Version()
			fmt.Printf("Current migration version: %v, dirty: %v\n", version, dirty)
			if dirty {
				if forceErr := m.Force(int(version - 1)); forceErr != nil {
					log.Fatal("Error! Couldn't force database to previous version!", forceErr)
				}
				return nil
			}

		}
		if err == migrate.ErrNoChange {
			fmt.Println("Error! No change has been realized!", err)
			return nil
		}
		log.Fatal("Error! While trying to up the migrations!!", err)
	}
	return nil
}

func IsDirtyErr(err error) bool {
	return strings.Contains(err.Error(), "Dirty database")
}

// function katchekiii lina wash kaynin files fyal up khawyin

func HasEmptyFilesUp(migrationPath string) (bool, []string, error) {
	emptyFiles := []string{}
	if err := filepath.WalkDir(migrationPath, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, "up.sql") {
			fileInfo, errStat := os.Stat(path)
			if errStat != nil {
				return errStat
			}
			if fileInfo.Size() == 0 {
				emptyFiles = append(emptyFiles, fileInfo.Name())
				return nil
			}
		}
		return nil
	}); err != nil {
		return false, nil, err
	}
	return len(emptyFiles) != 0, emptyFiles, nil
}
