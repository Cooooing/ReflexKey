package com.example.kernel.entity.base;

import com.example.kernel.util.FileUtils;
import lombok.extern.slf4j.Slf4j;

import java.io.File;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.text.SimpleDateFormat;

@Slf4j
public class Constant {

    public final static String NAME = "kernel";
    public final static String VERSION = "0.0.1";
    public final static String DATABASE_VERSION = "20240930";

    public final static SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");

    public final static String rootPath = System.getProperty("user.dir");
    public static String workspacePath = rootPath + File.separator + "workspace";
    public static String dataPath = workspacePath + File.separator + "data";
    public static String logPath = workspacePath + File.separator + "logs";

    public final static String sqliteFileName = "kernel.db";
    public static String sqliteFilePath = workspacePath + File.separator + sqliteFileName;
    public final static String sqliteDriverName = "org.sqlite.JDBC";
    public final static String sqliteUrl = "jdbc:sqlite:" + sqliteFilePath;
    public final static String sqliteInitFile = "sql/sqlite-init.sql";

    public static void resetWorkspacePath(String path) {
        Path normalizePath = Paths.get(path).toAbsolutePath().normalize();
        if (!FileUtils.createDirectory(normalizePath.toString())) {
            return;
        }
        if (!normalizePath.isAbsolute()) {
            workspacePath = rootPath + File.separator + normalizePath;
        } else {
            workspacePath = normalizePath.toString();
        }
        dataPath = workspacePath + File.separator + "data";
        sqliteFilePath = workspacePath + File.separator + sqliteFileName;
    }

    public final static String MDC_TRACE = "tid"; // 日志追踪标识

}
