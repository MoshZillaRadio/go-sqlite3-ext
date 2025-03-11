package sqlite3_ext

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

type entrypoint struct {
	lib  string
	proc string
}

var libNames = []entrypoint{
	{"libgo-sqlite3-ext.so", "sqlite3_fileio_init"},
}

var ErrNotFound = errors.New("MoshZillaRadio/go-sqlite3-ext: sqlite3-ext extension not found")

func init() {
	sql.Register("sqlite3-ext",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				for _, v := range libNames {
					if err := conn.LoadExtension(v.lib, v.proc); err == nil {
						return nil
					}
				}
				return ErrNotFound
			},
		})
}
