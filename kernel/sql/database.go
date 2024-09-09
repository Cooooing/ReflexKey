package sql

import (
	"fmt"
	"kernel/common"
	"kernel/model"
	"kernel/util"
	"os"
	"runtime"
	"runtime/debug"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

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
	if nil != DB {
		_ = closeDatabase()
	}
	var err error
	DB, err = gorm.Open(sqlite.Open(util.DBPath), &gorm.Config{})
	if nil != err {
		common.Fatal(common.ExitCodeReadOnlyDatabase, "create database failed: %s", err)
	}
}

func closeDatabase() (err error) {
	if nil == DB {
		return
	}

	db, err := DB.DB()
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

func QueryForList(sql string, params ...any) []map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	db, _ := DB.DB()
	rows, err := db.Query(sql, params...)
	common.Info("query: %s", sql)
	if err != nil {
		common.Error("query failed: %s", err)
		return nil
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valPointers := make([]interface{}, count)
	for rows.Next() {

		// 获取各列的值的地址
		for i := 0; i < count; i++ {
			valPointers[i] = &values[i]
		}

		// 获取各列的值，放到对应的地址中
		_ = rows.Scan(valPointers...)

		// 一条数据的Map (列名和值的键值对)
		entry := make(map[string]interface{})

		// Map 赋值
		for i, col := range columns {
			var v interface{}

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

func QueryForPage(page int64, size int64, sql string, params ...any) model.Page {
	pageSql := "select * from (" + sql + ") limit ? offset ?"
	PageParams := append(params, size, (page-1)*size)
	fmt.Println(params)
	fmt.Println(PageParams)
	res := model.Page{
		Total: QueryForCount(sql, params...),
		Data:  QueryForList(pageSql, PageParams...),
		Page:  page,
		Size:  size,
	}
	return res
}

func QueryForCount(sql string, params ...any) int64 {
	var res int64
	db, _ := DB.DB()
	sql = "select count(*) from (" + sql + ")"
	row := db.QueryRow(sql, params...)
	common.Info("query: %s", sql)
	_ = row.Scan(&res)
	return res
}
