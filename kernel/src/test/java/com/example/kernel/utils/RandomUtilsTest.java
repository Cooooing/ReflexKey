package com.example.kernel.utils;

import com.example.kernel.common.util.RandomUtils;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;

import java.util.HashMap;
import java.util.Map;

@Slf4j
public class RandomUtilsTest {

    /**
     * 测试 生成重复的uuid
     */
    @Test
    public void testRepeatUUID() {
        int count = 1000 * 10000;
        Map<String, Integer> map = new HashMap<>();
        int i = 0;
        while (true) {
            i++;
            String shortUuid = RandomUtils.generateShortUuid();
            if (map.containsKey(shortUuid)) {
                log.info("repeat uuid:{} number:{}", shortUuid, i);
                break;
            } else {
                map.put(shortUuid, i);
            }
            if (i >= count) {
                log.info("{} uuid was no repeat", i);
                break;
            }
        }
        assert i >= count;
    }

    /**
     * 测试 生成随机密码
     */
    @Test
    public void testRepeatPassword() {
        int count = 1000 * 10000;
        Map<String, Integer> map = new HashMap<>();
        int i = 0;
        while (true) {
            i++;
            String password = RandomUtils.generatePassword(8, true, true, true, false, null, null);
            if (map.containsKey(password)) {
                log.info("repeat password:{} number:{}", password, i);
                break;
            } else {
                map.put(password, i);
            }
            if (i >= count) {
                log.info("{} password was no repeat", i);
                break;
            }
        }
        assert i >= count;
    }


}
