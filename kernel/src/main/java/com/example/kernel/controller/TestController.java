package com.example.kernel.controller;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.example.kernel.entity.Test;
import com.example.kernel.mapper.TestMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

/**
 * Todo
 **/
@Slf4j
@RequiredArgsConstructor
@RestController
@RequestMapping("/test")
public class TestController {

    private final TestMapper testMapper;

    @GetMapping("/")
    public String test(){
        return "hello world";
    }

    @GetMapping("/test2")
    public List<Test> test2(){
        List<Test> tests = testMapper.selectList(new LambdaQueryWrapper<>());
        for (Test test : tests) {
            log.info(test.toString());
        }
        return tests;
    }

}
