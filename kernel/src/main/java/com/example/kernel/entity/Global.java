package com.example.kernel.entity;

import java.io.File;

public class Global {

    public static String NAME = "kernel";
    public static String VERSION = "0.0.1";
    public static String DATABASE_VERSION = "20240930";


    public static String rootPath = System.getProperty("user.dir");
    public static String workspacePath = rootPath + File.separator + "workspace";
    public static String dataPath = workspacePath + File.separator + "data";
    public static String logPath = workspacePath + File.separator + "logs";

    public static String sqliteFileName = "kernel.db";
    public static String sqliteFilePath = workspacePath + File.separator + sqliteFileName;
    public static String sqliteDriverName = "org.sqlite.JDBC";
    public static String sqliteUrl = "jdbc:sqlite:" + sqliteFilePath;

}
