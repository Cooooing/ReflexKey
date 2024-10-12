package com.example.kernel.common.config;

import com.example.kernel.common.properties.AsyncTaskThreadPoolProperties;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.aop.interceptor.AsyncUncaughtExceptionHandler;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.AsyncConfigurer;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

import java.util.concurrent.Executor;
import java.util.concurrent.ThreadPoolExecutor;

/**
 * 异步任务处理线程池类
 */
@Slf4j
@EnableAsync
@RequiredArgsConstructor
@Configuration
public class AsyncTaskExecuteConfigurator implements AsyncConfigurer {
    private final AsyncTaskThreadPoolProperties config;

    @Override
    public Executor getAsyncExecutor() {
        ThreadPoolTaskExecutor executor = new OverrideThreadPoolExecutor();
        executor.setCorePoolSize(config.getCorePoolSize());
        executor.setMaxPoolSize(config.getMaxPoolSize());
        executor.setQueueCapacity(config.getQueueCapacity());
        executor.setKeepAliveSeconds(config.getKeepAliveSeconds());
        executor.setThreadNamePrefix(config.getThreadNamePrefix());
        executor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        executor.initialize();
        return executor;
    }

    @Override
    public AsyncUncaughtExceptionHandler getAsyncUncaughtExceptionHandler() {
        return (arg0, arg1, arg2) -> log.error("async exception:{} class:{} method:{} args:{}", arg0.getMessage(), arg1.getDeclaringClass(), arg1.getName(), arg2);
    }
}
