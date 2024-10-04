package common

import (
	"context"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

func (g *GuSqliteLog) LogMode(logLevel gormLog.LogLevel) gormLog.Interface {
	return g
}

func (g *GuSqliteLog) Info(ctx context.Context, msg string, data ...interface{}) {
	Log.Info(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (g *GuSqliteLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	Log.Warn(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (g *GuSqliteLog) Error(ctx context.Context, msg string, data ...interface{}) {
	Log.Error(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (g *GuSqliteLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if rows == -1 {
		Log.Trace(utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
	} else {
		Log.Trace(utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}

}
