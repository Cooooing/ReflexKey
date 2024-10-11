package com.example.kernel.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.example.kernel.entity.po.HistoryRecord;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface HistoryRecordMapper extends BaseMapper<HistoryRecord> {
}
