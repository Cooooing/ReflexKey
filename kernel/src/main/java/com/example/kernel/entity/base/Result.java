package com.example.kernel.entity.base;

import com.alibaba.fastjson2.JSON;
import com.example.kernel.entity.enums.ExceptionEnum;
import com.example.kernel.entity.enums.ResultEnum;
import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Data;
import lombok.experimental.Accessors;

import java.util.Date;

@Schema(description = "统一返回格式")
@Data
@Accessors(chain = true)
public class Result<T> {
    @Schema(description = "是否成功")
    private Boolean success;
    @Schema(description = "状态码")
    private Integer code;
    @Schema(description = "提示信息")
    private String message;
    @Schema(description = "数据")
    private T data;
    @Schema(description = "时间戳")
    private String time = Global.dateFormat.format(new Date());

    public Result() {
    }


    public Result(Boolean success, ResultEnum resultEnum) {
        this.success = success;
        this.code = resultEnum.getKey();
        this.message = resultEnum.getValue();
    }

    public Result(Boolean success, ResultEnum resultEnum, T data) {
        this.success = success;
        this.code = resultEnum.getKey();
        this.message = resultEnum.getValue();
        this.data = data;
    }

    public Result(boolean success, ExceptionEnum exceptionEnum) {
        this.success = success;
        this.code = exceptionEnum.getKey();
        this.message = exceptionEnum.getValue();
    }

    public Result(Boolean success, ExceptionEnum exceptionEnum, T data) {
        this.success = success;
        this.code = exceptionEnum.getKey();
        this.message = exceptionEnum.getValue();
        this.data = data;
    }

    public Result(Boolean success, String message) {
        this.success = success;
        this.message = message;
    }

    public static <T> Result<T> error(ExceptionEnum exceptionEnum) {
        return new Result<>(false, exceptionEnum);
    }

    public static <T> Result<T> error(ExceptionEnum exceptionEnum, T data) {
        return new Result<>(false, exceptionEnum, data);
    }

    public static <T> Result<T> error(Exception e) {
        return new Result<>(false, e.getMessage());
    }

    public static <T> Result<T> success(ResultEnum resultEnum) {
        return new Result<>(true, resultEnum);
    }

    public static <T> Result<T> success(ResultEnum resultEnum, T data) {
        return new Result<>(true, resultEnum, data);
    }

    public static <T> Result<T> success(T data) {
        return simpleJudge(true, data);
    }

    public static <T> Result<T> fail(ResultEnum resultEnum) {
        return new Result<>(false, resultEnum);
    }

    public static <T> Result<T> fail(ResultEnum resultEnum, T data) {
        return new Result<>(false, resultEnum, data);
    }

    /**
     * 根据boolean简单判断返回
     */
    public static <T> Result<T> simpleJudge(boolean success) {
        return success ? success(ResultEnum.SUCCESS) : fail(ResultEnum.FAIL);
    }

    public static <T> Result<T> simpleJudge(boolean success, T data) {
        return success ? success(ResultEnum.SUCCESS, data) : fail(ResultEnum.FAIL, data);
    }

    @Override
    public String toString() {
        return JSON.toJSONString(this);
    }


}
