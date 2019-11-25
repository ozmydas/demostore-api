package app

import (
	"github.com/jinzhu/gorm"
)

/****
* disini kita deklarasikan variable buat diakses handler
****/

type Env struct {
	DB          *gorm.DB
	PORT        string
	UPLOAD_PATH string
}
