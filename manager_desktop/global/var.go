package global

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GORMClient  *gorm.DB = new(gorm.DB)
	Conn        *sql.DB  = new(sql.DB)
	Trans       ut.Translator
	RedisClient *redis.Client = new(redis.Client)
)
