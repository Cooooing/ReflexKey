package sql

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"kernel/common/log"
	"kernel/util"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

var (
	db             *sql.DB
	historyDB      *sql.DB
	assetContentDB *sql.DB
)

var initDatabaseLock = sync.Mutex{}

func init() {
	regex := func(re, s string) (bool, error) {
		re = strings.ReplaceAll(re, "\\\\", "\\")
		return regexp.MatchString(re, s)
	}

	sql.Register("sqlite3_extended", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			return conn.RegisterFunc("regexp", regex, true)
		},
	})
}

func InitDatabase(forceRebuild bool) (err error) {
	initDatabaseLock.Lock()
	defer initDatabaseLock.Unlock()

	//ClearCache()
	//disableCache()
	//defer enableCache()

	//if forceRebuild {
	//	ClearQueue()
	//}

	initDBConnection()
	//treenode.InitBlockTree(forceRebuild)

	//if !forceRebuild {
	//	// 检查数据库结构版本，如果版本不一致的话说明改过表结构，需要重建
	//	if util.DatabaseVer == getDatabaseVer() {
	//		return
	//	}
	//	log.Info("the database structure is changed, rebuilding database...")
	//}

	// 不存在库或者版本不一致都会走到这里

	closeDatabase()
	//if gulu.File.IsExist(util.DBPath) {
	//	if err = removeDatabaseFile(); nil != err {
	//		log.Error("remove database file [%s] failed: %s", util.DBPath, err)
	//		util.PushClearProgress()
	//		err = nil
	//	}
	//}

	initDBConnection()
	//initDBTables()

	log.Info("reinitialized database [%s]", util.DBPath)
	return
}

func initDBConnection() {
	if nil != db {
		_ = closeDatabase()
	}
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
	var err error
	db, err = sql.Open("sqlite3_extended", dsn)
	if nil != err {
		log.Fatal(log.ExitCodeReadOnlyDatabase, "create database failed: %s", err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(24 * time.Hour)
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
