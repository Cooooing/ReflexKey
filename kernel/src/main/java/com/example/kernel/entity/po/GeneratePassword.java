package com.example.kernel.entity.po;


import com.baomidou.mybatisplus.annotation.TableName;
import com.example.kernel.entity.base.Entity;
import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.experimental.Accessors;

@Data
@EqualsAndHashCode(callSuper = true)
@Accessors(chain = true)
@TableName("generate_password")
public class GeneratePassword extends Entity {
    @Schema(description = "密码")
    private String password;
}
