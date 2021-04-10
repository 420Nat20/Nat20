package route

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type BaseController struct {
	DB     *gorm.DB
	Router *mux.Router
}
