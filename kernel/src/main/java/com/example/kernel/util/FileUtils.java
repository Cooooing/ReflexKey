package com.example.kernel.util;

import lombok.extern.slf4j.Slf4j;

import java.io.File;
import java.io.IOException;

@Slf4j
public class FileUtils {

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
