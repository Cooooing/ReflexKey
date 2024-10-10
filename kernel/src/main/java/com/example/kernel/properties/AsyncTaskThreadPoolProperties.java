package com.example.kernel.properties;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
@ConfigurationProperties(prefix = "base.boot.system.async")
@Data
public class AsyncTaskThreadPoolProperties {
    private Integer corePoolSize = 5;
    private Integer maxPoolSize = 50;
    private Integer keepAliveSeconds = 60;
    private Integer queueCapacity = 10000;
    private String threadNamePrefix = "async-";
}
