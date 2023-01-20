package global

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GORMClient  *gorm.DB      = new(gorm.DB)
	Conn        *sql.DB       = new(sql.DB)
	RedisClient *redis.Client = new(redis.Client)
	// hget hash:gateway:service_info test_grpc_service
	HashServiceInfoKey = "hash:gateway:service_info"
	// hget hash:gateway:app_info tianjia
	HashAppInfoKey = "hash:gateway:app_info"
)
