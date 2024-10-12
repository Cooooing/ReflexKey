package com.example.kernel.common.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.example.kernel.common.entity.po.HistoryRecord;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface HistoryRecordMapper extends BaseMapper<HistoryRecord> {
}
