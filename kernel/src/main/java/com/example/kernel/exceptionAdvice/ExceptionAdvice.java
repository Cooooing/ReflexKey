package com.example.kernel.exceptionAdvice;

import com.example.kernel.entity.base.Result;
import com.example.kernel.entity.enums.ExceptionEnum;
import lombok.extern.log4j.Log4j2;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.servlet.resource.NoResourceFoundException;

@Log4j2
@RestControllerAdvice
public class ExceptionAdvice {

    @ExceptionHandler(Exception.class)
    public Result<Object> exceptionAdvice(Exception e) {
        log.error(e.getMessage(), e);
        return Result.error(ExceptionEnum.INTERNAL_SERVER_ERROR);
    }

    @ExceptionHandler(DefinedException.class)
    public Result<Object> definedExceptionAdvice(DefinedException e) {
        if (e.getException() == null) {
            log.error(e.getMessage(), e);
        } else {
            log.error(e.getException().getMessage(), e.getException());
        }
        return Result.error(e);
    }

    @ExceptionHandler(NoResourceFoundException.class)
    public Result<Object> NoResourceFoundExceptionAdvice(NoResourceFoundException e) {
        log.error("method = {} path = {} {}", e.getHttpMethod(), e.getResourcePath(), e.getMessage(), e);
        return Result.error(ExceptionEnum.NOT_FOUND);
    }
}
