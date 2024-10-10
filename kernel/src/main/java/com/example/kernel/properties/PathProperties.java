package com.example.kernel.properties;

import com.example.kernel.entity.base.Constant;
import com.example.kernel.util.FileUtils;
import jakarta.annotation.PostConstruct;
import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;

@Configuration
@ConfigurationProperties(prefix = "base.boot.system.path")
@Data
public class PathProperties {
    private String workspace;

    @PostConstruct
    public void init() throws IOException {

        Constant.resetWorkspacePath(workspace);

        FileUtils.createDirectory(workspace);
        FileUtils.createDirectory(Constant.dataPath);
        FileUtils.createFile(Constant.sqliteFilePath);
    }
}
