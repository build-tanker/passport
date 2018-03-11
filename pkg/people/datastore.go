package people

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
type Datastore interface {
	Add(source, name, email, pictureURL, gender, sourceID string) (string, error)
	View(id string) (Person, error)
	ViewBySourceID(sourceID string) (Person, error)
	Delete(id string) error
}

type datastore struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewDatastore - create a new datastore for people
func NewDatastore(conf *config.Config, db *sqlx.DB) Datastore {
	return &datastore{
		conf: conf,
		db:   db,
	}
}

func (s *datastore) Add(source, name, email, pictureURL, gender, sourceID string) (string, error) {
	id := s.generateUUID()
	return id, nil
}

func (s *datastore) View(id string) (Person, error) {
	return Person{}, nil
}

func (s *datastore) ViewBySourceID(sourceID string) (Person, error) {
	return Person{}, nil
}

func (s *datastore) Delete(id string) error {
	return nil
}

func (s *datastore) generateUUID() string {
	return uuid.NewV4().String()
}
