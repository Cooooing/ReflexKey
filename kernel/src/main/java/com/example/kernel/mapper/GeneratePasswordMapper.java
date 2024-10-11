package com.example.kernel.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.example.kernel.entity.po.GeneratePassword;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface GeneratePasswordMapper extends BaseMapper<GeneratePassword> {
}
