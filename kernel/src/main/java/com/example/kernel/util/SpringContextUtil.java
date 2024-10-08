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


@Slf4j
@Component
public class SpringContextUtil implements ApplicationContextAware {

    private static ApplicationContext applicationContext;

    public ApplicationContext getApplicationContext() {
        return applicationContext;
    }

    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        SpringContextUtil.applicationContext = applicationContext;
    }

    public Object getBean(String name) {
        return getApplicationContext().getBean(name);
    }

    public <T> T getBean(Class<T> clazz) {
        return getApplicationContext().getBean(clazz);
    }

    public <T> T getBean(String name, Class<T> clazz) {
        return getApplicationContext().getBean(name, clazz);
    }

    public Map<String, Object> getBeansWithAnnotation(Class<? extends Annotation> annotationClass) {
        return applicationContext.getBeansWithAnnotation(annotationClass);
    }
}