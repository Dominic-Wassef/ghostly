package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type RememberToken struct {
	ID            int       `db:"id,omitempty"`
	UserID        int       `db:"user_id"`
	RememberToken string    `db:"remember_token"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}


func (t *RememberToken) Table() string {
	return "remember_tokens"
}

func (t *RememberToken) InsertToken(userID int, token string) error {
	collection := upper.Collection(t.Table())
	rememberToken := RememberToken{
		UserID: userID,
		RememberToken: token,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := collection.Insert(rememberToken)
	if err != nil {
		return err
	}
	return nil
}

func (t *RememberToken) Delete(rememberToken string) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{"remember_token": rememberToken})
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}
