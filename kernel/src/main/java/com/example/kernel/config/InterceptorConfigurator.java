package com.example.kernel.config;

import com.example.kernel.interceptor.LogInterceptor;
import com.example.kernel.interceptor.PermissionInterceptor;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

/**
 * 拦截器配置类
 **/
@RequiredArgsConstructor
@Configuration
public class InterceptorConfigurator implements WebMvcConfigurer {

    private final PermissionInterceptor permissionInterceptor;
    private final LogInterceptor logInterceptor;

    @Value("${spring.profiles.active}")
    private String active;

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(logInterceptor).order(-100);
        registry.addInterceptor(permissionInterceptor).addPathPatterns("/**")
                .excludePathPatterns(
                        "/doc.html", "/webjars/**", "/v3/api-docs/**",
                        "/favicon.ico")
                .order(0);
    }
}
