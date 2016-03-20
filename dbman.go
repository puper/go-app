package goapp

import (
	"github.com/go-ozzo/ozzo-dbx"
	"math/rand"
)

type DBMan struct {
	write *dbx.DB
	read  []*dbx.DB
}

func NewDBMan() *DBMan {
	return &DBMan{
		write: dbx.Open()
		read: func () {
			read := []*dbx.DB
			for xxx {
				read = append(read, )
			}
		}()
	}
}

func (d *DBMan) GetDB(write bool) *dbx.DB {
	if write {
		return d.write
	}
	return d.read[rand.Intn(len(d.read))]
}
