package com.example.kernel.service;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;


@Slf4j
@RequiredArgsConstructor
@Service
public class Test {
    @Async
    public void asyncTrace() {
        log.info("执行线程池中的方法asyncTrace，重写了线程池");
    }

}
