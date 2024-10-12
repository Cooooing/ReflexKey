package com.example.kernel.utils;

import com.alibaba.fastjson2.JSON;
import com.example.kernel.util.BencodeUtils;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;

@Slf4j
public class BencodeUtilsTest {

    private final String str = "d4:nick8:Cooooing5:skilll6:Coding10:Basketballe4:blog26:https://cooooing.github.io3:agei22ee";
    private final Object object = new HashMap<String, Object>() {
        {
            put("nick", "Cooooing");
            put("age", 22);
            put("blog", "https://cooooing.github.io");
            put("skill", new ArrayList<>() {
                {
                    add("Coding");
                    add("Basketball");
                }
            });
        }
    };

    @Test
    public void testEncode() {
        String encode = BencodeUtils.encode(object);
        log.info("encode: {}", encode);
        assert encode.equals(str);
    }

    @Test
    public void testDecode() {
        Object decode = BencodeUtils.decode(str.getBytes());
        log.info("decode: {}", JSON.toJSONString(decode));
        assert JSON.toJSONString(decode).equals(JSON.toJSONString(object));
    }

}
