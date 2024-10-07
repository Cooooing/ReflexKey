package com.example.kernel.entity;

import com.example.kernel.util.FileUtils;
import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;


/**
 * Todo
 **/
@Data
@Component
@ConfigurationProperties(prefix = "path")
public class SqliteConfig {
    private String workspace;
    private String data;

    public void initDir() {
        String dir = System.getProperty("user.dir");
        FileUtils.createDirectory(dir + java.io.File.separator + workspace);

    }
}
