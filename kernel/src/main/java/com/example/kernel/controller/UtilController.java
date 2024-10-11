package com.example.kernel.controller;


import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.example.kernel.entity.base.PageVO;
import com.example.kernel.entity.base.Result;
import com.example.kernel.entity.po.GeneratePassword;
import com.example.kernel.entity.vo.GeneratePasswordQueryVO;
import com.example.kernel.service.GeneratePasswordService;
import com.example.kernel.util.RandomUtils;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Slf4j
@RequiredArgsConstructor
@RestController
@Tag(name = "工具")
@RequestMapping("/util")
public class UtilController {

    private final GeneratePasswordService generatePasswordService;

    @Operation(summary = "生成密码")
    @PostMapping("/generatePassword")
    public Result<List<String>> generatePassword(@RequestBody GeneratePasswordQueryVO vo) {
        if (vo.getCount() <= 0 || vo.getLength() <= 0 || (!vo.getHasNumber() && !vo.getHasLowerCase() && !vo.getHasUpperCase() && !vo.getHasSpecialChar() && "".equals(vo.getIncludeChars()) && "".equals(vo.getExcludeChars()))) {
            return Result.simpleJudge(false);
        }
        List<String> passwords = new ArrayList<>(vo.getCount());
        for (int i = 0; i < vo.getCount(); i++) {
            passwords.add(RandomUtils.generatePassword(vo.getLength(), vo.getHasLowerCase(), vo.getHasUpperCase(), vo.getHasNumber(), vo.getHasSpecialChar()
                    , vo.getIncludeChars().chars().mapToObj(c -> (char) c).collect(Collectors.toList())
                    , vo.getExcludeChars().chars().mapToObj(c -> (char) c).collect(Collectors.toList())));
        }
        generatePasswordService.saveBatch(passwords.stream().map(password -> new GeneratePassword().setPassword(password)).collect(Collectors.toList()));
        return Result.success(passwords);
    }

    @Operation(summary = "获取生成密码历史")
    @GetMapping("/getGeneratePassword")
    public Result<Page<GeneratePassword>> generatePassword(PageVO<GeneratePassword> pageVO) {
        Page<GeneratePassword> page = generatePasswordService.page(pageVO.toMybatisPlusPage());
        return Result.success(page);
    }


}
