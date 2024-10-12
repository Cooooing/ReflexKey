package com.example.kernel.common.entity.base;

import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class PageVO<T> {
    @Schema(description = "当前页")
    private Integer current;
    @Schema(description = "每页数量")
    private Integer size;

    public PageVO() {
        this.current = 1;
        this.size = 10;
    }

    public Page<T> toMybatisPlusPage() {
        current = current == null ? 1 : current;
        size = size == null ? 10 : size;
        return new Page<>(current, size);
    }
}
