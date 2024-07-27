package pg

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"url_shortener/internal/database"
)

type Postgres struct {
	db *sql.DB
}

func New(storagePath string, log *slog.Logger) (*Postgres, error) {
	const op = "pg.Postgres::New()"
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, errors.New(op + err.Error())
	}
	log.Info("Database connected")
	return &Postgres{db: db}, nil
}

func (p *Postgres) SaveUrl(urlToSave string, alias string, token string) (id int64, err error) {
	const op = "pg.Postgres.SaveUrl()"
	stmt, err := p.db.Prepare("INSERT INTO url(alias, url, d_token) VALUES($1, $2, $3)")
	if err != nil {
		return 0, errors.New(op + err.Error())
	}
	defer Close(stmt)
	err = stmt.QueryRow(alias, urlToSave, token).Scan(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) && errors.Is(err, sql.ErrNoRows) {
			return 0, database.UrlAlreadyExist
		}
		return 0, errors.New(op + err.Error())
	}
	return id, nil
}

func (p *Postgres) DeleteByAlias(alias string) error {
	const op = "pg.Postgres.DeleteByAlias()"
	stmt, err := p.db.Prepare("DELETE FROM url WHERE alias = $1")
	if err != nil {
		return errors.New(op + err.Error())
	}
	defer Close(stmt)
	_, err = stmt.Exec(alias)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) && errors.Is(err, sql.ErrNoRows) {
			return database.AliasNotFound
		}
	}
	return nil
}

func (p *Postgres) DeleteByToken(token string) error {
	const op = "pg.Postgres.DeleteByToken()"
	stmt, err := p.db.Prepare("DELETE FROM url WHERE d_token = $1")
	if err != nil {
		return errors.New(op + err.Error())
	}
	defer Close(stmt)
	_, err = stmt.Exec(token)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) && errors.Is(err, sql.ErrNoRows) {
			return database.TokenNotFound
		}
	}
	return nil
}

func (p *Postgres) FindByAlias(alias string) (url string, err error) {
	const op = "pg.Postgres.FindByAlias()"
	stmt, err := p.db.Prepare("SELECT url FROM url WHERE alias = $1")
	if err != nil {
		return "", errors.New(op + err.Error())
	}
	defer Close(stmt)
	err = stmt.QueryRow(alias).Scan(&url)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) && errors.Is(err, sql.ErrNoRows) {
			return "", database.AliasNotFound
		}
	}
	return url, nil
}

func Close(stmt *sql.Stmt) {
	err := stmt.Close()
	if err != nil {
		log.Fatal("Failed to close statement")
	}
}
