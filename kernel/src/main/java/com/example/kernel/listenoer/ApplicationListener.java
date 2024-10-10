package com.example.kernel.listenoer;

import com.example.kernel.entity.base.Constant;
import com.example.kernel.util.FileUtils;
import com.example.kernel.util.InitUtils;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.io.IOException;

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
    public void run(String... args) throws IOException {
        FileUtils.lockDirectory(Constant.workspacePath);
        initUtils.initTables();
        if ("dev".equals(active)) {

        }


        log.info("application is inited");
    }

    /**
     * 应用启动关闭前的回调
     */
    @Override
    public void destroy() {

        FileUtils.unlockDirectory(Constant.workspacePath);
        log.info("application is closed");
    }


}