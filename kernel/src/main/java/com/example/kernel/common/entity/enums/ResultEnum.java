package com.example.kernel.common.entity.enums;

import lombok.Getter;

@Getter
public enum ResultEnum {
    // other
    SUCCESS(200, "nice"),
    FAIL(400, "!nice"),




    ;


    private final Integer key;
    private final String value;

    ResultEnum(Integer key, String value) {
        this.key = key;
        this.value = value;
    }

}
