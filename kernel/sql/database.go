package sql

import (
	//_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kernel/common"
	"kernel/model"
	"kernel/util"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

var DB *gorm.DB

var initDatabaseLock = sync.Mutex{}

func InitDatabase(forceRebuild bool) (err error) {
	initDatabaseLock.Lock()
	defer initDatabaseLock.Unlock()

	initDBConnection()

	if !forceRebuild {
		// 检查数据库结构版本，如果版本不一致的话说明改过表结构，需要重建
		if model.DatabaseVer == getDatabaseVer() {
			return
		}
		common.Log.Info("the database structure is changed, rebuilding database...")
	}

	// 不存在库或者版本不一致都会走到这里
	_ = closeDatabase()
	if common.File.IsExist(util.DBPath) {
		if err = removeDatabaseFile(); nil != err {
			common.Log.Error("remove database file [%s] failed: %s", util.DBPath, err)
			err = nil
		}
	}

	initDBConnection()
	initDBTables()

	common.Log.Info("reinitialized database [%s]", util.DBPath)
	return
}

func initDBConnection() {
	if nil != DB {
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
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: &common.SqlLog,
	})
	if nil != err {
		common.Log.Fatal(common.ExitCodeReadOnlyDatabase, "create sqlite3 database failed: %s", err)
	}
}

func closeDatabase() (err error) {
	if nil == DB {
		return
	}

	db, _ := DB.DB()
	err = db.Close()
	if err != nil {
		return err
	}
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

func initDBTables() {
	sql :=
		``
	DB.Exec(sql)
}

func InsertOrUpdate() {

}

func QueryForListByMap(sql string, params ...any) []map[string]any {
	var res []map[string]any
	DB.Raw(sql, params...).Scan(&res)
	return res
}

func QueryForList(res any, sql string, params ...any) {
	// 检查 res 是否为切片类型
	v := reflect.ValueOf(res)
	if v.Kind() != reflect.Ptr {
		common.Log.Error("res must be a pointer to a slice")
		return
	}
	// 确保 elem 是可寻址的
	elem := v.Elem()
	if !elem.CanSet() {
		common.Log.Error("res must be addressable, got %s", v.Type())
		return
	}
	logSql(sql, params...)
	DB.Raw(sql, params...).Scan(res)
	common.Log.Info("QueryForListByMap: %v", res)
}

func QueryForPageByMap(current int64, size int64, sql string, params ...any) model.Page {
	pageSql := "select * from (" + sql + ") limit ? offset ?"
	PageParams := append(params, size, (current-1)*size)
	res := model.Page{
		Total:   QueryForCount(sql, params...),
		Data:    QueryForListByMap(pageSql, PageParams...),
		Page:    0,
		Current: current,
		Size:    size,
	}
	if res.Total != 0 && res.Size > 0 {
		res.Page = (res.Total-1)/size + 1
	}
	return res
}
func QueryForPage(current int64, size int64, scan any, sql string, params ...any) model.Page {
	pageSql := "select * from (" + sql + ") limit ? offset ?"
	PageParams := append(params, size, (current-1)*size)
	QueryForList(scan, pageSql, PageParams...)
	res := model.Page{
		Total:   QueryForCount(sql, params...),
		Data:    scan,
		Page:    0,
		Current: current,
		Size:    size,
	}
	if res.Total != 0 && res.Size > 0 {
		res.Page = (res.Total-1)/size + 1
	}
	return res
}

func QueryForCount(sql string, params ...any) int64 {
	var res int64
	sql = "select count(*) from (" + sql + ")"
	DB.Raw(sql, params...).Scan(&res)
	return res
}

func logSql(sql string, params ...any) {
	sql = sqlStringFormat(sql, params...)
	common.Log.Info(sql, params...)
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
