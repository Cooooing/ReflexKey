package com.example.kernel.controller;

import com.example.kernel.mapper.TestMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

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
    public Object test2(){
        return testMapper.selectList(null);
    }

}
