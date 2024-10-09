package com.example.kernel.entity.enums;

import lombok.Getter;

@Getter
public enum ExceptionEnum {

    // 异常
    EXCEPTION(500, "意料之外的异常"),


    NO_PERMISSION(403, "你没得权限"),
    NO_AUTH(401, "你能不能先登录一下"),
    NOT_FOUND(404, "未找到该资源!"),
    INTERNAL_SERVER_ERROR(500, "服务器跑路了"),

    ;


    private final Integer key;
    private final String value;

    ExceptionEnum(Integer key, String value) {
        this.key = key;
        this.value = value;
    }

}
