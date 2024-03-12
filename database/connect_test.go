package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	db := PostgresDB{}
	err := db.Connect("test")
	assert.Nil(t, err)

	err = db.Close()
	assert.Nil(t, err)
}
