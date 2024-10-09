package com.example.kernel.entity.base;

import java.io.File;
import java.text.SimpleDateFormat;

public class Global {

    public final static String NAME = "kernel";
    public final static String VERSION = "0.0.1";
    public final static String DATABASE_VERSION = "20240930";

    public final static SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");

    public final static String rootPath = System.getProperty("user.dir");
    public final static String workspacePath = rootPath + File.separator + "workspace";
    public final static String dataPath = workspacePath + File.separator + "data";
    public final static String logPath = workspacePath + File.separator + "logs";

    public final static String sqliteFileName = "kernel.db";
    public final static String sqliteFilePath = workspacePath + File.separator + sqliteFileName;
    public final static String sqliteDriverName = "org.sqlite.JDBC";
    public final static String sqliteUrl = "jdbc:sqlite:" + sqliteFilePath;
    public final static String sqliteInitFile = "sql/sqlite-init.sql";

}
