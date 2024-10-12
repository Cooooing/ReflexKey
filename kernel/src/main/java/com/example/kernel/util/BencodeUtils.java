package com.example.kernel.util;

import com.example.kernel.exceptionAdvice.DefinedException;
import lombok.extern.slf4j.Slf4j;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Slf4j
public class BencodeUtils {

    // 编码方法
    public static String encode(Object obj) {
        StringBuilder sb = new StringBuilder();
        if (obj == null) {
            return "";
        }
        switch (obj) {
            case Map<?, ?> map -> {
                sb.append("d");
                for (Map.Entry<?, ?> entry : map.entrySet()) {
                    sb.append(encode(entry.getKey()));
                    sb.append(encode(entry.getValue()));
                }
                sb.append("e");
            }
            case Integer ignored -> sb.append("i").append(obj).append("e");
            case List<?> list -> {
                sb.append("l");
                for (Object item : list) {
                    sb.append(encode(item));
                }
                sb.append("e");
            }
            case String str -> sb.append(str.length()).append(':').append(str);
            case null, default -> throw new IllegalArgumentException("Unsupported type: " + obj.getClass());
        }
        return sb.toString();
    }

    // 解码方法
    public static Object decode(byte[] s) {
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

    // 辅助方法：解码对象
    private static Object[] decodeObject(byte[] s, int index) {
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
                index = (int) keyResult[1];
                Object[] valueResult = decodeObject(s, index);
                Object value = valueResult[0];
                index = (int) valueResult[1];
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
            throw new IllegalArgumentException(String.format("Invalid Bencode format %s at index %d", b, index));
        }
    }
}


