package com.example.kernel.interceptor;


import com.example.kernel.entity.base.Constant;
import com.example.kernel.util.RandomUtils;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.NonNull;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.slf4j.MDC;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.util.AntPathMatcher;
import org.springframework.util.CollectionUtils;
import org.springframework.util.PathMatcher;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.util.ContentCachingRequestWrapper;
import org.springframework.web.util.ContentCachingResponseWrapper;

import java.net.InetAddress;
import java.net.URLDecoder;
import java.net.UnknownHostException;
import java.nio.charset.StandardCharsets;
import java.util.*;


/**
 * 日志拦截器 追踪请求
 **/
@Slf4j
@Component
public class LogInterceptor implements HandlerInterceptor {
    private ContentCachingRequestWrapper cachingRequest;
    private ContentCachingResponseWrapper cachingResponse;

    private String requestIp;
    private String requestUri;
    private String requestMethod;
    private String requestHeaders;
    private String requestParams;
    private Map<String, String> bodyContent;
    private String requestBody;
    private long startTime;

    @Override
    public boolean preHandle(@NonNull HttpServletRequest request, @NonNull HttpServletResponse response, @NonNull Object handler) {
        MDC.put(Constant.MDC_TRACE, RandomUtils.generateShortUuid());

        // 如果是被排除的uri，不记录 access_log
        if (matchExclude(request.getRequestURI())) {
            return true;
        }

        // Wrapper 封装 Request 和 Response
        cachingRequest = new ContentCachingRequestWrapper(request);
        cachingResponse = new ContentCachingResponseWrapper(response);

        // 获取请求内容
        requestIp = getRequestIp(request);
        requestUri = request.getRequestURI();
        requestMethod = request.getMethod();
        requestHeaders = getRequestHeaders(cachingRequest);
        requestParams = getRequestParams(cachingRequest).replaceAll("\n", "");
        bodyContent = getBodyContent(cachingRequest, cachingResponse);
        requestBody = bodyContent.get("requestBody");
        // 记录 access_log
        final List<String> logs = new ArrayList<>();
        logs.add("ip --> " + requestIp);
        logs.add("; uri = " + requestUri);
        logs.add("; method = " + requestMethod);
        logs.add(";\nrequestParams = " + requestParams);
        logs.add(";\nrequestHeaders = " + requestHeaders);
        logs.add(";\nrequestBody = " + requestBody);
        startTime = System.currentTimeMillis();
        log.info(String.join("", logs));
        return true;
    }

    @Override
    public void postHandle(@NonNull HttpServletRequest request, @NonNull HttpServletResponse response, @NonNull Object handler, ModelAndView modelAndView) throws Exception {
        final long useTime = System.currentTimeMillis() - startTime;

        String responseHeaders = getResponseHeaders(cachingResponse);
        String responseBody = bodyContent.get("responseBody");

        final List<String> logs = new ArrayList<>();
        logs.add("ip <-- " + requestIp);
        logs.add("; uri = " + requestUri);
        logs.add("; method = " + requestMethod);
        logs.add("; useTime=" + useTime + "ms");
        logs.add(";\nresponseStatus = " + cachingResponse.getStatus());
        logs.add(";\nresponseHeaders = " + responseHeaders);
        logs.add(";\nresponseBody = " + responseBody);
        log.info(String.join("", logs));

        // 复制响应体到响应
        cachingResponse.copyBodyToResponse();

        MDC.remove(Constant.MDC_TRACE);
    }


    public static final Set<String> excludeUris = new HashSet<>();
    private static final PathMatcher URI_PATH_MATCHER = new AntPathMatcher();

    private static final List<String> DEFAULT_DOWNLOAD_CONTENT_TYPE = List.of("application/vnd.ms-excel",//.xls
            "application/msexcel",//.xls
            "application/cvs",//.cvs
            MediaType.APPLICATION_OCTET_STREAM_VALUE,//.*（ 二进制流，不知道下载文件类型）
            "application/x-xls",//.xls
            "application/msword",//.doc
            MediaType.TEXT_PLAIN_VALUE,//.txt
            "application/x-gzip"//.gz
    );

    /**
     * 从 HttpServletRequest 中获取 ip
     */
    private static String getRequestIp(HttpServletRequest request) {
        String ip = request.getHeader("x-forwarded-for");
        if (ip == null || ip.isEmpty() || "unknown".equalsIgnoreCase(ip)) {
            ip = request.getHeader("Proxy-Client-IP");
        }
        if (ip == null || ip.isEmpty() || "unknown".equalsIgnoreCase(ip)) {
            ip = request.getHeader("X-Forwarded-For");
        }
        if (ip == null || ip.isEmpty() || "unknown".equalsIgnoreCase(ip)) {
            ip = request.getHeader("WL-Proxy-Client-IP");
        }
        if (ip == null || ip.isEmpty() || "unknown".equalsIgnoreCase(ip)) {
            ip = request.getHeader("X-Real-IP");
        }
        if (ip == null || ip.isEmpty() || "unknown".equalsIgnoreCase(ip)) {
            ip = request.getRemoteAddr();
            if ("127.0.0.1".equalsIgnoreCase(ip) || "0:0:0:0:0:0:0:1".equalsIgnoreCase(ip)) {
                // 根据网卡取本机配置的 IP
                InetAddress iNet = null;
                try {
                    iNet = InetAddress.getLocalHost();
                } catch (UnknownHostException e) {
                    log.error("unknown host", e);
                }
                if (iNet != null) ip = iNet.getHostAddress();
            }
        }
        // 对于通过多个代理的情况，分割出第一个 IP
        if (ip != null && ip.length() > 15) {
            if (ip.indexOf(",") > 0) {
                ip = ip.substring(0, ip.indexOf(","));
            }
        }
        return "0:0:0:0:0:0:0:1".equals(ip) ? "127.0.0.1" : ip;
    }

    private String getRequestParams(final HttpServletRequest request) {
        final Map<String, String[]> requestParams = new HashMap<>(request.getParameterMap());
        final List<String> pairs = new ArrayList<>();
        if (!requestParams.isEmpty()) {
            for (final Map.Entry<String, String[]> entry : requestParams.entrySet()) {
                final String name = entry.getKey();
                final String[] value = entry.getValue();
                if (value == null) {
                    pairs.add(name + "=");
                } else {
                    for (final String v : value) {
                        pairs.add(name + "=" + StringUtils.trimToEmpty(v));
                    }
                }
            }
        }
        String requestParamsStr = CollectionUtils.isEmpty(pairs) ? StringUtils.EMPTY : String.join("&", pairs);
        if (StringUtils.equalsIgnoreCase(request.getContentType(), MediaType.APPLICATION_FORM_URLENCODED_VALUE)) {
            requestParamsStr = URLDecoder.decode(requestParamsStr, StandardCharsets.UTF_8);
        }
        return requestParamsStr;
    }

    private String getRequestHeaders(final HttpServletRequest request) {
        final Enumeration<String> headerNames = request.getHeaderNames();
        final List<String> headers = new ArrayList<>();
        while (headerNames.hasMoreElements()) {
            final String key = headerNames.nextElement();
            headers.add(key + ':' + request.getHeader(key));
        }
        return '[' + String.join(",", headers) + ']';
    }

    private String getResponseHeaders(final HttpServletResponse response) {
        Collection<String> headerNames = response.getHeaderNames();
        final List<String> headers = new ArrayList<>();
        for (String key : headerNames) {
            headers.add(key + ':' + response.getHeader(key));
        }
        return '[' + String.join(",", headers) + ']';
    }

    /**
     * 获取请求体/响应体内容
     */
    private Map<String, String> getBodyContent(ContentCachingRequestWrapper request, ContentCachingResponseWrapper response) {
        // 请求体
        String requestBody;
        if (isUpload(request)) {
            requestBody = "upload...";
        } else {
            byte[] requestContentAsByteArray = request.getContentAsByteArray();
            if (request.getContentLength() <= 1024) {
                requestBody = new String(requestContentAsByteArray);
            } else {
                requestBody = new String(requestContentAsByteArray, 0, 1024).replaceAll("[\n\r]", "") + "...";
            }
        }
        // 响应体
        String responseBody;
        if (isDownload(response)) {
            responseBody = "download...";
        } else {
            byte[] responseContentAsByteArray = response.getContentAsByteArray();
            if (response.getContentSize() <= 1024) {
                responseBody = new String(responseContentAsByteArray);
            } else {
                responseBody = new String(responseContentAsByteArray, 0, 1024).replaceAll("[\n\r]", "") + "...";
            }
        }
        HashMap<String, String> res = new HashMap<>(2);
        res.put("requestBody", requestBody);
        res.put("responseBody", responseBody);
        return res;
    }

    private boolean isUpload(final HttpServletRequest request) {
        final String contentType = request.getHeader(HttpHeaders.CONTENT_TYPE);
        if (StringUtils.isBlank(contentType)) {
            return false;
        }
        return StringUtils.containsIgnoreCase(contentType, MediaType.MULTIPART_FORM_DATA_VALUE);
    }

    private boolean isDownload(final HttpServletResponse response) {
        final String contentType = response.getContentType();
        if (StringUtils.isBlank(contentType)) {
            return false;
        }
        return DEFAULT_DOWNLOAD_CONTENT_TYPE.stream().anyMatch(it -> StringUtils.equalsIgnoreCase(it, contentType));
    }

    /**
     * 排除不需要记录日志的 uri
     */
    private boolean matchExclude(final String uri) {
        if (CollectionUtils.isEmpty(excludeUris)) {
            return false;
        }
        for (final String excludeUri : excludeUris) {
            if (URI_PATH_MATCHER.match(excludeUri, uri)) {
                return true;
            }
        }
        return false;
    }
}
