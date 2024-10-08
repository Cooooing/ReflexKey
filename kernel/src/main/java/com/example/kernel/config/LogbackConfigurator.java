package com.example.kernel.config;

import ch.qos.logback.classic.spi.LoggerContextListener;
import ch.qos.logback.core.Context;
import com.example.kernel.entity.Global;
import ch.qos.logback.classic.Level;
import ch.qos.logback.classic.Logger;
import ch.qos.logback.classic.LoggerContext;
import ch.qos.logback.core.spi.ContextAwareBase;
import ch.qos.logback.core.spi.LifeCycle;

public class LogbackConfigurator extends ContextAwareBase implements LoggerContextListener, LifeCycle {
    private boolean started = false;

    @Override
    public void start() {
        if (started) {
            return;
        }
        Context context = getContext();
        context.putProperty("LOG_PATH", Global.logPath);
        started = true;
    }

    @Override
    public void stop() {

    }

    @Override
    public boolean isStarted() {
        return false;
    }

    @Override
    public boolean isResetResistant() {
        return false;
    }

    @Override
    public void onStart(LoggerContext loggerContext) {

    }

    @Override
    public void onReset(LoggerContext loggerContext) {

    }

    @Override
    public void onStop(LoggerContext loggerContext) {

    }

    @Override
    public void onLevelChange(Logger logger, Level level) {

    }
}