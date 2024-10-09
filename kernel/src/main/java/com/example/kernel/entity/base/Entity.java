package com.example.kernel.entity.base;

import com.baomidou.mybatisplus.annotation.*;
import lombok.*;
import lombok.experimental.Accessors;

@Data
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode
@ToString
@Accessors(chain = true)
public class Entity {
    // 主键
    @TableId(value = "`id`", type = IdType.ASSIGN_ID)
    private Integer id;
    // 创建时间
    @TableField(value = "`create_time`", fill = FieldFill.INSERT)
    private String createTime;
    // 更新时间
    @TableField(value = "`update_time`", fill = FieldFill.INSERT_UPDATE)
    private String updateTime;
    // 逻辑删除
    @TableField(value = "`deleted_time`")
    @TableLogic
    private Integer deleted;
}