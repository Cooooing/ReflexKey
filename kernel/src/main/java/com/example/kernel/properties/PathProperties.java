package com.example.kernel.properties;

import com.example.kernel.entity.base.Global;
import com.example.kernel.util.FileUtils;
import jakarta.annotation.PostConstruct;
import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;

@Configuration
@ConfigurationProperties(prefix = "path")
@Data
public class PathProperties {
    private String workspace;
    private String data;
    private String log;

    @PostConstruct
    public void init() throws IOException {
        if (workspace == null || workspace.isEmpty()) {
            workspace = Global.workspacePath;
        }
        if (data == null || data.isEmpty()) {
            data = Global.dataPath;
        }
        if (log == null || log.isEmpty()) {
            log = Global.logPath;
        }

        FileUtils.createDirectory(workspace);
        FileUtils.createDirectory(data);
        FileUtils.createDirectory(Global.logPath);
        FileUtils.createFile(Global.sqliteFilePath);
    }
}
