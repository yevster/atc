package handler

import (
	"sync"

	"github.com/winston-ci/logbuffer"
	"github.com/winston-ci/winston/db"
)

type Handler struct {
	db db.DB

	logs      map[string]*logbuffer.LogBuffer
	logsMutex *sync.RWMutex
}

func NewHandler(db db.DB) *Handler {
	return &Handler{
		db: db,

		logs:      make(map[string]*logbuffer.LogBuffer),
		logsMutex: new(sync.RWMutex),
	}
}
