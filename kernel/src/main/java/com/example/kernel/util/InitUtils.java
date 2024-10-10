package com.example.kernel.util;

import com.example.kernel.entity.base.Constant;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.jdbc.ScriptRunner;
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

    public void rebuildDataSource() {
        log.info("rebuild datasource...");
        try {
            boolean b = FileUtils.deleteFile(Constant.sqliteFilePath);


//            conn.close();
        } catch (Exception e) {
            log.error("rebuild datasource error", e);
        }
    }

}
