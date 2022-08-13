package server

import (
	"fmt"
	db "manyu/collect/mongo"
)

func Init() {
	db.Init()

	r := NewRouter()
	r.Run(":" + "4000")
}