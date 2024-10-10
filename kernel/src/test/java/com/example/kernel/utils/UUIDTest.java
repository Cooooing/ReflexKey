package com.example.kernel.utils;

import com.example.kernel.util.UUIDUtils;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;

import java.util.HashMap;
import java.util.Map;

@Slf4j
public class UUIDTest {


    /**
     * 测试 生成重复的uuid
     */
    @Test
    public void repeatUUID() {
        int count = 1000 * 10000;
        Map<String, Integer> map = new HashMap<>();
        int i = 0;
        while (true) {
            i++;
            String ticket = UUIDUtils.generateShortUuid();
            if (map.containsKey(ticket)) {
                log.info("repeat uuid:{} number:{}", ticket, i);
                break;
            } else {
                map.put(ticket, i);
            }
            if (i >= count) {
                log.info("{} uuid was no repeat", i);
                break;
            }
        }
        assert i >= count;
    }


}
