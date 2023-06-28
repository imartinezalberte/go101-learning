// chapter2
//
// Requirements:
//   - The to-do app simply allows a user to add a text to a to-do, to overwrite all content.
package chapter2

import "errors"

var unauthorizedErr = errors.New("unauthorized")

type (
	// Here we are using named parameters to clarify what's in here
	AuthFn func(username string) bool
	Db     struct {
		F AuthFn
	}

	ToDo struct {
		Text string
		Db   *Db
	}
)

func NewToDo(db *Db) *ToDo {
	return &ToDo{
		Text: "",
		Db:   db,
	}
}

func (t *ToDo) Write(username, input string) error {
	if t.Db.IsAuth(username) {
		t.Text = input
		return nil
	}
	return unauthorizedErr
}

func (t *ToDo) Append(username, input string) error {
	if t.Db.IsAuth(username) {
		t.Text += input
		return nil
	}
	return unauthorizedErr
}

func (d Db) IsAuth(username string) bool {
	return d.F(username)
}
