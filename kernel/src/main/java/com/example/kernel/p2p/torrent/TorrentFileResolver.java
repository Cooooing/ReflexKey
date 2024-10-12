package com.example.kernel.p2p.torrent;

import com.alibaba.fastjson2.JSON;
import com.example.kernel.common.exceptionAdvice.DefinedException;
import com.example.kernel.p2p.entity.Torrent;
import lombok.extern.slf4j.Slf4j;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.*;

@Slf4j
public class TorrentFileResolver {

    private byte[] torrentBytes;
    private int infoStart;
    private int infoEnd;
    private Torrent result;

    private TorrentFileResolver() {
    }

    public static TorrentFileResolver init(String path) throws IOException {
        TorrentFileResolver resolver = new TorrentFileResolver();
        resolver.torrentBytes = Files.readAllBytes(Paths.get(path));
        resolver.infoStart = 0;
        resolver.infoEnd = 0;
        return resolver;
    }

    public Torrent read() {
        if (result != null) {
            return result;
        }
        result = JSON.parseObject(JSON.toJSONString(decode(torrentBytes)), Torrent.class);
        byte[] infoByte = new byte[infoEnd - infoStart + 1];
        System.arraycopy(torrentBytes, infoStart, infoByte, 0, infoEnd - infoStart + 1);
        result.setInfoBencode(infoByte);
        return result;
    }

    // Todo 生成torrent文件
    public void write(String path) {

    }

    private Object decode(byte[] s) {
        Object o = null;
        if (s == null || s.length == 0) {
            return o;
        }
        try {
            o = decodeObject(s, 0)[0];
        } catch (Exception e) {
            throw new DefinedException(e);
        }
        return o;
    }

    private Object[] decodeObject(byte[] s, int index) {
        byte b = s[index];
        if (b == 'i') {
            // 整数类型
            index++;
            int start = index;
            while (s[index] != 'e') {
                index++;
            }
            long value = Long.parseLong(new String(s, start, index - start));
            index++;
            return new Object[]{value, index};
        } else if ('0' <= b && b <= '9') {
            // 字符串类型
            int start = index;
            while (s[index] != ':') {
                index++;
            }
            int length = Integer.parseInt(new String(s, start, index - start));
            index++;
            String str = new String(s, index, length);
            index += length;
            return new Object[]{str, index};
        } else if (b == 'd') {
            // 字典类型
            index++;
            Map<String, Object> map = new LinkedHashMap<>();
            while (s[index] != 'e') {
                Object[] keyResult = decodeObject(s, index);
                String key = (String) keyResult[0];
                boolean record = key.equals("info");
                if (record) {
                    infoStart = index + 6;
                }
                index = (int) keyResult[1];

                Object[] valueResult = decodeObject(s, index);
                Object value = valueResult[0];
                index = (int) valueResult[1];
                if (record) {
                    infoEnd = index - 1;
                }
                map.put(key, value);
            }
            index++;
            return new Object[]{map, index};
        } else if (b == 'l') {
            // 列表类型
            index++;
            List<Object> list = new ArrayList<>();
            while (s[index] != 'e') {
                Object[] result = decodeObject(s, index);
                list.add(result[0]);
                index = (int) result[1];
            }
            index++;
            return new Object[]{list, index};
        } else {
            // 忽略无效的字符
            index++;
            return decodeObject(s, index);
        }
    }

    // Todo 将torrent进行编码
    private byte[] encode() {
        return null;
    }

}
