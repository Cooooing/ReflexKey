package com.example.kernel.controller;


import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.example.kernel.entity.base.Constant;
import com.example.kernel.entity.base.PageVO;
import com.example.kernel.entity.base.Result;
import com.example.kernel.entity.po.HistoryRecord;
import com.example.kernel.entity.vo.GeneratePasswordQueryVO;
import com.example.kernel.service.HistoryRecordService;
import com.example.kernel.util.BencodeUtils;
import com.example.kernel.util.RandomUtils;
import com.github.xiaoymin.knife4j.annotations.DynamicParameter;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
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

    private final HistoryRecordService historyRecordService;

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
        historyRecordService.saveBatch(passwords.stream().map(password -> new HistoryRecord().setType(Constant.GENERATE_PASSWORD).setValue(password)).collect(Collectors.toList()));
        return Result.success(passwords);
    }

    @Operation(summary = "获取生成密码历史")
    @GetMapping("/getGeneratePassword")
    public Result<Page<HistoryRecord>> generatePassword(PageVO<HistoryRecord> pageVO) {
        Page<HistoryRecord> page = historyRecordService.page(pageVO.toMybatisPlusPage(), new LambdaQueryWrapper<HistoryRecord>().eq(HistoryRecord::getType, Constant.GENERATE_PASSWORD));
        return Result.success(page);
    }

    @Operation(summary = "添加剪贴板历史")
    @Parameter(name = "value", description = "剪贴板内容")
    @GetMapping("/addClipboard")
    public Result<Boolean> addClipboard(String value) {
        return Result.simpleJudge(historyRecordService.save(new HistoryRecord().setType(Constant.CLIPBOARD).setValue(value)));
    }

    @Operation(summary = "获取剪贴板历史")
    @GetMapping("/getClipboard")
    public Result<Page<HistoryRecord>> getClipboard(PageVO<HistoryRecord> pageVO) {
        Page<HistoryRecord> page = historyRecordService.page(pageVO.toMybatisPlusPage(), new LambdaQueryWrapper<HistoryRecord>().eq(HistoryRecord::getType, Constant.CLIPBOARD));
        return Result.success(page);
    }

    @Operation(summary = "bencode 编码")
    @DynamicParameter(example = "{\"nick\":\"Cooooing\",\"skill\":[\"Coding\",\"Basketball\"],\"blog\":\"https://cooooing.github.io\",\"age\":22}")
    @PostMapping("/bencode")
    public Result<String> bencode(@RequestBody Object obj) {
        return Result.success(BencodeUtils.encode(obj));
    }

    @Operation(summary = "bencode 解码")
    @DynamicParameter(example = "d4:nick8:Cooooing5:skilll6:Coding10:Basketb2alle4:blog26:https://cooooing.github.io3:agei22ee")
    @PostMapping("/bdecode")
    public Result<Object> bdecode(@RequestBody Object s) {
        return Result.success(BencodeUtils.decode(s.toString().getBytes()));
    }


}
