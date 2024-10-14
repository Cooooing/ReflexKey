package com.example.kernel.common.util;

import lombok.extern.slf4j.Slf4j;
import org.apache.tomcat.util.buf.HexUtils;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.security.MessageDigest;

@Slf4j
public final class DigestUtils {

    private static final String SHA_1 = "SHA-1";
    private static final MessageDigest sha1;
    private static final String SHA_256 = "SHA-256";
    private static final MessageDigest sha256;
    private static final String MD_5 = "MD5";
    private static final MessageDigest md5;

    static {
        try {
            sha1 = MessageDigest.getInstance(SHA_1);
            sha256 = MessageDigest.getInstance(SHA_256);
            md5 = MessageDigest.getInstance(MD_5);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    public static byte[] toSHA1(File arg0) {
        try (FileInputStream fis = new FileInputStream(arg0)) {
            byte[] buffer = new byte[8192];
            int bytesRead;
            while ((bytesRead = fis.read(buffer)) != -1) {
                sha1.update(buffer, 0, bytesRead);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        return sha1.digest();
    }

    public static byte[] toSHA1(String arg0) {
        return sha1.digest(arg0.getBytes());
    }

    public static byte[] toSHA1(byte[] arg0) {
        return sha1.digest(arg0);
    }

    public static byte[] toSHA256(String arg0) {
        return sha256.digest(arg0.getBytes());
    }

    public static byte[] toSHA256(byte[] arg0) {
        return sha256.digest(arg0);
    }

    public static byte[] toMD5(String arg0) {
        return md5.digest(arg0.getBytes());
    }

    public static byte[] toMD5(byte[] arg0) {
        return md5.digest(arg0);
    }

    public static String toHexString(String arg0) {
        return HexUtils.toHexString(arg0.getBytes());
    }

    public static String toHexString(byte[] arg0) {
        return HexUtils.toHexString(arg0);
    }

}
