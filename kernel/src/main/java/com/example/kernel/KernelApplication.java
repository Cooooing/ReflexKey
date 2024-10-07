package com.example.kernel;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

import javax.sql.DataSource;
import java.sql.SQLException;
import java.util.Map;

@SpringBootApplication
public class KernelApplication {

    public static void main(String[] args) throws SQLException {
        ConfigurableApplicationContext ctx = SpringApplication.run(KernelApplication.class, args);
        Map<String, DataSource> dataSourceMap = ctx.getBeansOfType(DataSource.class);
        for (Map.Entry<String, DataSource> entry : dataSourceMap.entrySet()) {
            String name = entry.getKey();
            DataSource dataSource = entry.getValue();
            System.out.println(name);
            System.out.println(dataSource.getConnection()); // 这里会抛出异常，直接throws走了
        }
    }

}
