package com.example.kernel.common.util;

import lombok.extern.slf4j.Slf4j;

import java.io.File;
import java.io.IOException;
import java.nio.file.Path;

@Slf4j
public final class FileUtils {

    public static boolean lockDirectory(String path) throws IOException {
        path = path + File.separator + ".lock";
        if (isExist(path)) {
            return false;
        }
        return createFile(path);
    }

    public static boolean unlockDirectory(String path) {
        path = path + File.separator + ".lock";
        if (!isExist(path)) {
            return true;
        }
        return deleteFile(path);
    }

    public static boolean isExist(String path) {
        return new File(path).exists();
    }

    public static boolean createDirectory(String path) {
        boolean flag;
        File file = new File(path);
        if (!file.exists()) {
            flag = file.mkdirs();
            log.info("create directory {}:{}", flag, path);
        } else {
            flag = true;
        }
        return flag;
    }

    public static boolean createDirectory(Path path) {
        return createDirectory(path.toString());
    }

    public static boolean createFile(String path) throws IOException {
        boolean flag = false;
        File file = new File(path);
        if (createDirectory(file.getParentFile().getPath())) {
            if (!file.exists()) {
                flag = file.createNewFile();
                log.info("create file {}:{}", flag, path);
            }
        }
        return flag;
    }

    public static boolean createFile(Path path) throws IOException {
        return createFile(path.toString());
    }

    public static boolean deleteFile(String path) {
        boolean flag = false;
        File file = new File(path);
        if (file.exists()) {
            flag = file.delete();
            log.info("delete file {}:{}", flag, path);
        }
        return flag;
    }

}
