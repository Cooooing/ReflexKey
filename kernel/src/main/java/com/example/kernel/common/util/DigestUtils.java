package com.example.kernel.common.util;

import lombok.extern.slf4j.Slf4j;
import org.apache.tomcat.util.buf.HexUtils;

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

    public static String toSHA1(String arg0) {
        return HexUtils.toHexString(sha1.digest(arg0.getBytes()));
    }

    public static String toSHA1(byte[] arg0) {
        return HexUtils.toHexString(sha1.digest(arg0));
    }

    public static String toSHA256(String arg0) {
        return HexUtils.toHexString(sha256.digest(arg0.getBytes()));
    }

    public static String toSHA256(byte[] arg0) {
        return HexUtils.toHexString(sha256.digest(arg0));
    }

    public static String toMD5(String arg0) {
        return HexUtils.toHexString(md5.digest(arg0.getBytes()));
    }

    public static String toMD5(byte[] arg0) {
        return HexUtils.toHexString(md5.digest(arg0));
    }

}
