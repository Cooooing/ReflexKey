package com.example.kernel.exceptionAdvice;

import com.example.kernel.entity.enums.ExceptionEnum;
import lombok.Getter;
import org.springframework.util.CollectionUtils;

import java.util.List;

@Getter
public class DefinedException extends RuntimeException {

    private final ExceptionEnum exceptionEnum;
    private final Exception exception;

    public DefinedException(ExceptionEnum exceptionEnum, Exception exception) {
        this.exceptionEnum = exceptionEnum;
        this.exception = exception;
    }

    public DefinedException(ExceptionEnum exceptionEnum) {
        super(exceptionEnum.getValue());
        this.exceptionEnum = exceptionEnum;
        this.exception = null;
    }

    public DefinedException(Exception exception) {
        this.exceptionEnum = ExceptionEnum.INTERNAL_SERVER_ERROR;
        this.exception = exception;
    }

    public DefinedException(String s) {
        super(s);
        this.exceptionEnum = null;
        this.exception = null;
    }

    public static void throwDefinedException(ExceptionEnum s, Boolean... isThrow) throws DefinedException {
        List<Boolean> isThrows = List.of(isThrow);
        if (!CollectionUtils.isEmpty(isThrows)) {
            if (isThrows.contains(false)) {
                throw new DefinedException(s);
            }
        }
    }

    public static void throwDefinedException(String s, Boolean... isThrow) throws DefinedException {
        List<Boolean> isThrows = List.of(isThrow);
        if (!CollectionUtils.isEmpty(isThrows)) {
            if (isThrows.contains(false)) {
                throw new DefinedException(s);
            }
        }
    }
}
