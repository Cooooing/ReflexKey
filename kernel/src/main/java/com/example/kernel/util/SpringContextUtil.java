package com.example.kernel.util;

import lombok.extern.slf4j.Slf4j;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.BeansException;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.stereotype.Component;

import java.lang.annotation.Annotation;
import java.util.Map;

/**
 * spring上下文工具类
 **/
@Slf4j
@Component
public class SpringContextUtil implements ApplicationContextAware {

    /**
     * 获取上下文对象
     */
    private static ApplicationContext applicationContext;

    /**
     * 获取 applicationContext
     */
    public ApplicationContext getApplicationContext() {
        return applicationContext;
    }

    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        SpringContextUtil.applicationContext = applicationContext;
    }

    /**
     * 通过 name 获取 bean 对象
     */
    public Object getBean(String name) {

        return getApplicationContext().getBean(name);
    }

    /**
     * 通过 class 获取 bean 对象
     */
    public <T> T getBean(Class<T> clazz) {
        return getApplicationContext().getBean(clazz);
    }

    /**
     * 通过 name，clazz  获取指定的 bean 对象
     */
    public <T> T getBean(String name, Class<T> clazz) {
        return getApplicationContext().getBean(name, clazz);
    }

    public Map<String, Object> getBeansWithAnnotation(Class<? extends Annotation> annotationClass) {
        return applicationContext.getBeansWithAnnotation(annotationClass);
    }
}