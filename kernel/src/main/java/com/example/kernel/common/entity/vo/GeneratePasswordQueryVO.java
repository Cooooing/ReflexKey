package com.example.kernel.common.entity.vo;

import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Data;

@Data
public class GeneratePasswordQueryVO {
    @Schema(description = "生成数量")
    private Integer count;
    @Schema(description = "密码长度")
    private Integer length;
    @Schema(description = "是否包含小写字母")
    private Boolean hasLowerCase;
    @Schema(description = "是否包含大写字母")
    private Boolean hasUpperCase;
    @Schema(description = "是否包含数字")
    private Boolean hasNumber;
    @Schema(description = "是否包含特殊字符")
    private Boolean hasSpecialChar;
    @Schema(description = "包含的字符")
    private String includeChars;
    @Schema(description = "排除的字符")
    private String excludeChars;

    public GeneratePasswordQueryVO() {
        this.count = 0;
        this.length = 0;
        this.hasLowerCase = false;
        this.hasUpperCase = false;
        this.hasNumber = false;
        this.hasSpecialChar = false;
        this.includeChars = "";
        this.excludeChars = "";
    }
}
