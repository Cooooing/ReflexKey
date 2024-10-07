package com.example.kernel.listenoer;

import com.example.kernel.KernelApplication;
import com.example.kernel.entity.SqliteConfig;
import com.zaxxer.hikari.HikariDataSource;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;
import org.springframework.beans.factory.DisposableBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.beans.factory.support.DefaultListableBeanFactory;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.sql.SQLException;

/**
 * 监听容器启动关闭
 **/
@Log4j2
@RequiredArgsConstructor
@Component
public class ApplicationListener implements CommandLineRunner, DisposableBean {

    @Value("${spring.profiles.active}")
    private String active;

    private final SqliteConfig sqliteConfig;
    private final DefaultListableBeanFactory beanFactory;

    /**
     * 应用启动成功后的回调
     */
    @Override
    public void run(String... args) {
        log.info("application started success and to init data...");

        initDatabase();
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

    public void initDatabase() {
        System.out.println(sqliteConfig);
        sqliteConfig.initDir();

        String dir = System.getProperty("user.dir");
        HikariDataSource dataSource = new HikariDataSource();
        dataSource.setDriverClassName("org.sqlite.JDBC");
        dataSource.setJdbcUrl("jdbc:sqlite:" + dir + java.io.File.separator + "workspace" + "/db/kernel.db");
//        dataSource.setUsername("root");
//        dataSource.setPassword("123456");
        try {
            dataSource.getConnection();
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
        System.out.println("db3 创建完成！");
        beanFactory.registerSingleton("db3", dataSource);
        // 重新执行依赖注入
        beanFactory.autowireBean(beanFactory.getBean(KernelApplication.class));

    }


}