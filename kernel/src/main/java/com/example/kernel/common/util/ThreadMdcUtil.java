package com.example.kernel.common.util;

import com.example.kernel.common.entity.base.Constant;
import org.slf4j.MDC;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.Callable;

public final class ThreadMdcUtil {

    public static void setTraceIdIfAbsent() {
        if (MDC.get(Constant.MDC_TRACE) == null || MDC.get(Constant.MDC_TRACE).isEmpty()) {
            String tid = UUID.randomUUID().toString().replace("-", "");
            MDC.put(Constant.MDC_TRACE, tid);
        }
    }

    public static <T> Callable<T> wrap(final Callable<T> callable, final Map<String, String> context) {
        return () -> {
            if (context == null) {
                MDC.clear();
            } else {
                MDC.setContextMap(context);
            }
            setTraceIdIfAbsent();
            try {
                return callable.call();
            } finally {
                MDC.clear();
            }
        };
    }

    public static Runnable wrap(final Runnable callable, final Map<String, String> context) {
        if (context == null) {
            MDC.clear();
        } else {
            MDC.setContextMap(context);
        }
        setTraceIdIfAbsent();
        try {
            return callable;
        } finally {
            MDC.clear();
        }
    }
}
