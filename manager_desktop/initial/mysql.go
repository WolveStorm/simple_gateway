package initial

import (
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"simple_gateway/global"
	"time"
)

func InitMysql() {
	conf := global.DebugFullConfig.MysqlConfig
	var err error
	global.Conn, err = sql.Open("mysql", conf.Dsn)
	if err != nil {
		zap.S().Error("open sql connection error")
		return
	}
	global.Conn.SetMaxOpenConns(conf.MaxOpenConn)
	global.Conn.SetMaxIdleConns(conf.MaxIdleConn)
	global.Conn.SetConnMaxLifetime(time.Duration(conf.MaxConnLifeTime) * time.Second)
	global.GORMClient, _ = gorm.Open(mysql.New(mysql.Config{Conn: global.Conn}), &gorm.Config{
		Logger: getLogger(),
	})
}

func getLogger() logger.Interface {
	file, _ := os.OpenFile(GetCurrentPath()+"/logs/mysql/"+"mysql.log", os.O_CREATE|os.O_APPEND, 0644)
	return logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info})
}

func CloseConn() {
	global.Conn.Close()
}
