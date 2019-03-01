package modes

import (
	"database/sql"
	"time"
)

type User struct {
	BaseModel

	UserName     string `gorm:"not null;size:12;unique"`
	UserPassWord string `gorm:"not null;size:40"`
	LoginIP      sql.NullString
	LoginTime    time.Time
}
