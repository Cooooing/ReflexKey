package com.example.kernel.config;

import com.example.kernel.util.ThreadMdcUtil;
import lombok.NonNull;
import org.slf4j.MDC;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

import java.util.concurrent.Callable;
import java.util.concurrent.Future;

public final class OverrideThreadPoolExecutor extends ThreadPoolTaskExecutor {

    @Override
    public void execute(@NonNull Runnable task) {
        super.execute(ThreadMdcUtil.wrap(task, MDC.getCopyOfContextMap()));
    }

    @Override
    @NonNull
    public <T> Future<T> submit(@NonNull Callable<T> task) {
        return super.submit(ThreadMdcUtil.wrap(task, MDC.getCopyOfContextMap()));
    }

    @Override
    @NonNull
    public Future<?> submit(@NonNull Runnable task) {
        return super.submit(ThreadMdcUtil.wrap(task, MDC.getCopyOfContextMap()));
    }
}
