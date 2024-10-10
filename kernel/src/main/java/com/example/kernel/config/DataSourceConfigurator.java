package com.example.kernel.config;


import com.example.kernel.entity.base.Constant;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.sql.DataSource;

@Configuration
public class DataSourceConfigurator {

    @Bean(name = "sqliteDataSource")
    public DataSource sqliteDataSource() {
        return DataSourceBuilder.create().driverClassName(Constant.sqliteDriverName).url(Constant.sqliteUrl).build();
    }


}
