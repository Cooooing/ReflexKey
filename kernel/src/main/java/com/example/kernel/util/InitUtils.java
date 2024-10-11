package com.example.kernel.util;

import com.example.kernel.entity.base.Constant;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.jdbc.ScriptRunner;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.sql.DataSource;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;
import java.sql.Connection;
import java.util.Objects;

@Slf4j
@RequiredArgsConstructor
@Component
public final class InitUtils {

    private final DataSource dataSource;

    @Value("${spring.profiles.active}")
    private String active;

    /**
     * 初始化表结构
     */
    public void initTables() {
        log.info("init tables...");
        try {
            Connection conn = dataSource.getConnection();
            ScriptRunner runner = new ScriptRunner(conn);
            runner.setErrorLogWriter(null);
            runner.setLogWriter(null);
            runner.setStopOnError(true);
            Resources.setCharset(StandardCharsets.UTF_8);
            runner.runScript(new InputStreamReader(Objects.requireNonNull(InitUtils.class.getClassLoader().getResourceAsStream(Constant.sqliteInitFile)), StandardCharsets.UTF_8));
            conn.close();
        } catch (Exception e) {
            log.error("init tables error", e);
        }
    }

    /**
     * 重建数据库 会导致数据丢失，仅在开发环境使用
     */
    public void rebuildDataSource() {
        if (!"dev".equals(active)) {
            return;
        }
        log.info("rebuild datasource...");
        try {
            FileUtils.deleteFile(Constant.sqliteFilePath);
            initTables();
        } catch (Exception e) {
            log.error("rebuild datasource error", e);
        }
    }

}
