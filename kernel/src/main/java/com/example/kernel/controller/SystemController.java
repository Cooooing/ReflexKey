package com.example.kernel.controller;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.example.kernel.entity.base.Result;
import com.example.kernel.entity.Config;
import com.example.kernel.mapper.TestMapper;
import io.swagger.v3.oas.annotations.tags.Tag;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@Slf4j
@RequiredArgsConstructor
@RestController
@Tag(name = "系统")
@RequestMapping("/system")
public class SystemController {

    private final TestMapper testMapper;

    @GetMapping("/ping")
    public Result<Object> ping(){
        return Result.success("pong");
    }

    @GetMapping("/test2")
    public List<Config> test2(){
        List<Config> configs = testMapper.selectList(new LambdaQueryWrapper<>());
        for (Config config : configs) {
            log.info(config.toString());
        }
        return configs;
    }

}
