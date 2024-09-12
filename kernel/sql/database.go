package sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"kernel/common"
	"kernel/model"
	"kernel/util"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

var db *sql.DB

var initDatabaseLock = sync.Mutex{}

func InitDatabase(forceRebuild bool) (err error) {
	initDatabaseLock.Lock()
	defer initDatabaseLock.Unlock()

	initDBConnection()

	if !forceRebuild {
		// 检查数据库结构版本，如果版本不一致的话说明改过表结构，需要重建
		//if conf.DatabaseVer == getDatabaseVer() {
		return
		//}
		common.Info("the database structure is changed, rebuilding database...")
	}

	// 不存在库或者版本不一致都会走到这里

	//_ = closeDatabase()
	//if common.IsExist(util.DBPath) {
	//	if err = removeDatabaseFile(); nil != err {
	//		common.Error("remove database file [%s] failed: %s", util.DBPath, err)
	//		err = nil
	//	}
	//}

	//initDBConnection()
	//initDBTables()

	common.Info("reinitialized database [%s]", util.DBPath)
	return
}

func initDBConnection() {
	if nil != db {
		_ = closeDatabase()
	}

	var err error
	dsn := util.DBPath + "?_journal_mode=WAL" +
		"&_synchronous=OFF" +
		"&_mmap_size=2684354560" +
		"&_secure_delete=OFF" +
		"&_cache_size=-20480" +
		"&_page_size=32768" +
		"&_busy_timeout=7000" +
		"&_ignore_check_constraints=ON" +
		"&_temp_store=MEMORY" +
		"&_case_sensitive_like=OFF"
	db, err = sql.Open("sqlite3", dsn)
	if nil != err {
		common.Fatal(common.ExitCodeReadOnlyDatabase, "create database failed: %s", err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(365 * 24 * time.Hour)
}

func closeDatabase() (err error) {
	if nil == db {
		return
	}

	err = db.Close()
	debug.FreeOSMemory()
	runtime.GC() // 没有这句的话文件句柄不会释放，后面就无法删除文件
	return
}

func removeDatabaseFile() (err error) {
	err = os.RemoveAll(util.DBPath)
	if nil != err {
		return
	}
	err = os.RemoveAll(util.DBPath + "-shm")
	if nil != err {
		return
	}
	err = os.RemoveAll(util.DBPath + "-wal")
	if nil != err {
		return
	}
	return
}

func initTables() {

}

func QueryForList(sql string, params ...any) []map[string]any {
	res := make([]map[string]any, 0)
	rows, err := db.Query(sql, params...)
	if err != nil {
		common.Error("query failed: %s", err)
		return nil
	}
	defer rows.Close()
	logSql(sql, params...)

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]any, count)
	valPointers := make([]any, count)
	for rows.Next() {

		// 获取各列的值的地址
		for i := 0; i < count; i++ {
			valPointers[i] = &values[i]
		}

		// 获取各列的值，放到对应的地址中
		_ = rows.Scan(valPointers...)

		// 一条数据的Map (列名和值的键值对)
		entry := make(map[string]any)

		// Map 赋值
		for i, col := range columns {
			var v any

			// 值复制给val(所以Scan时指定的地址可重复使用)
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				// 字符切片转为字符串
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}

		res = append(res, entry)
	}
	return res
}

func QueryForPage(current int64, size int64, sql string, params ...any) model.Page {
	pageSql := "select * from (" + sql + ") limit ? offset ?"
	PageParams := append(params, size, (current-1)*size)
	res := model.Page{
		Total:   QueryForCount(sql, params...),
		Data:    QueryForList(pageSql, PageParams...),
		Page:    0,
		Current: current,
		Size:    size,
	}
	if res.Total != 0 && res.Size > 0 {
		res.Page = (res.Total-1)/size + 1
	}
	//(page.totalRecord - 1L) / page.pageSize + 1L;
	return res
}

func QueryForCount(sql string, params ...any) int64 {
	var res int64
	sql = "select count(*) from (" + sql + ")"
	row := db.QueryRow(sql, params...)
	_ = row.Scan(&res)
	return res
}

func logSql(sql string, params ...any) {
	sql = sqlStringFormat(sql, params...)
	common.Info(sql, params...)
}

func sqlStringFormat(s string, args ...any) string {
	for _, arg := range args {
		f := "!"
		switch arg.(type) {
		case string:
			f = "%s"
		case float64:
			f = "%f"
		case float32:
			f = "%f"
		case int:
			f = "%d"
		case int64:
			f = "%d"
		case int32:
			f = "%d"
		case int16:
			f = "%d"
		case int8:
			f = "%d"
		case uint:
			f = "%d"
		case uint64:
			f = "%d"
		case uint32:
			f = "%d"
		case uint16:
			f = "%d"
		case uint8:
			f = "%d"
		}
		s = strings.Replace(s, "?", f, 1)
	}
	return s
}
