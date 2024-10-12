package com.example.kernel.util;

import com.sun.tools.javac.Main;
import lombok.extern.slf4j.Slf4j;

import java.io.File;
import java.io.IOException;
import java.io.InputStream;

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

    public static byte[] readResourceFileAsBytes(String fileName) {
        byte[] bytes = new byte[0];
        try (InputStream is = Main.class.getClassLoader().getResourceAsStream(fileName)) {
            if (is != null) {
                bytes = is.readAllBytes();
            }
        } catch (IOException e) {
            log.error("read resource file error:{}", fileName, e);
        }
        return bytes;
    }

}
