package person

import (
	"time"

	"github.com/build-tanker/passport/pkg/common/config"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// Person saves details for a user
type Person struct {
	ID         string    `db:"id" json:"id,omitempty"`
	Source     string    `db:"source" json:"source,omitempty"`
	Name       string    `db:"name" json:"name,omitempty"`
	Email      string    `db:"email" json:"email,omitempty"`
	PictureURL string    `db:"pictureURL" json:"pictureURL,omitempty"`
	Gender     string    `db:"gender" json:"gender,omitempty"`
	SourceID   string    `db:"source_id" json:"source_id,omitempty"`
	Deleted    bool      `db:"deleted" json:"deleted,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// Datastore for people
type store interface {
	add(source, name, email, pictureURL, gender, sourceID string) (string, error)
	view(id string) (Person, error)
	viewBySourceID(sourceID string) (Person, error)
	delete(id string) error
}

type persistentStore struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewDatastore - create a new datastore for people
func newStore(conf *config.Config, db *sqlx.DB) store {
	return &persistentStore{
		conf: conf,
		db:   db,
	}
}

func (s *persistentStore) add(source, name, email, pictureURL, gender, sourceID string) (string, error) {
	id := s.generateUUID()

	_, err := s.db.Queryx("INSERT INTO person (id, source, name, email, picture_url, gender, source_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", id, source, name, email, pictureURL, gender, sourceID)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *persistentStore) view(id string) (Person, error) {
	rows, err := s.db.Queryx("SELECT * FROM person WHERE deleted = FALSE AND id=$1 LIMIT 1", id)
	if err != nil {
		return Person{}, err
	}

	var person Person
	for rows.Next() {
		err = rows.StructScan(&person)
		if err != nil {
			return Person{}, err
		}
	}

	return person, nil
}

func (s *persistentStore) viewBySourceID(sourceID string) (Person, error) {
	rows, err := s.db.Queryx("SELECT * FROM person WHERE deleted = FALSE AND source_id=$1 LIMIT 1", sourceID)
	if err != nil {
		return Person{}, err
	}

	var person Person
	for rows.Next() {
		err = rows.StructScan(&person)
		if err != nil {
			return Person{}, err
		}
	}

	return person, nil
}

func (s *persistentStore) delete(id string) error {
	_, err := s.db.Queryx("UPDATE person SET deleted = TRUE WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *persistentStore) generateUUID() string {
	return uuid.NewV4().String()
}
