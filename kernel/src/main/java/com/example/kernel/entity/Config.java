package com.example.kernel.entity;

import com.baomidou.mybatisplus.annotation.TableName;
import com.example.kernel.entity.base.Entity;
import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.experimental.Accessors;

@Data
@EqualsAndHashCode(callSuper = true)
@Accessors(chain = true)
@TableName("configs")
public class Config extends Entity {

    @Schema(description = "环境(all,win,linux,mac,android,ios)")
    private String env;

    @Schema(description = "设备标识")
    private String device;

    @Schema(description = "类别")
    private String type;

    @Schema(description = "只读 0:可写 1:只读")
    private String onlyRead;

    @Schema(description = "key")
    private String key;

    @Schema(description = "value")
    private String value;

}
