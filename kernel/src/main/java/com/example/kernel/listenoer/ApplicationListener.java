package com.example.kernel.listenoer;

import com.example.kernel.util.InitUtils;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

/**
 * 监听容器启动关闭
 **/
@Log4j2
@RequiredArgsConstructor
@Component
public class ApplicationListener implements CommandLineRunner, DisposableBean {

    @Value("${spring.profiles.active}")
    private String active;

    private final InitUtils initUtils;

    /**
     * 应用启动成功后的回调
     */
    @Override
    public void run(String... args) {
        log.info("application started success and to init data...");
        initUtils.initTables();
        if ("dev".equals(active)) {

        }


        log.info("application init data success");
    }

    /**
     * 应用启动关闭前的回调
     */
    @Override
    public void destroy() {
        log.info("application is closing and to close source...");

        log.info("application is closed");
    }




}