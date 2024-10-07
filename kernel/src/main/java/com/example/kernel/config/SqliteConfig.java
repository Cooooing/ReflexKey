package com.example.kernel.config;


import com.example.kernel.entity.Global;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.sql.DataSource;

@Configuration
public class SqliteConfig {

    /**
     * 配置sqlite数据源
     */
    @Bean(name = "sqliteDataSource")
    public DataSource sqliteDataSource() {

        String dir = System.getProperty("user.dir");
        Global.driverName = "org.sqlite.JDBC";
        Global.dbUrl = "jdbc:sqlite:" + dir + java.io.File.separator + "workspace" + "/db/kernel.db";

        return DataSourceBuilder.create().driverClassName(Global.driverName).url(Global.dbUrl).build();
    }


}
