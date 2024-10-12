package com.example.kernel.common.entity.po;


import com.baomidou.mybatisplus.annotation.TableName;
import com.example.kernel.common.entity.base.Entity;
import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.experimental.Accessors;

@Data
@EqualsAndHashCode(callSuper = true)
@Accessors(chain = true)
@TableName("history_record")
public class HistoryRecord extends Entity {

    @Schema(description = "value")
    private String value;

    @Schema(description = "类别")
    private String type;
}
