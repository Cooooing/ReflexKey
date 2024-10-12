package com.example.kernel.common.service;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.example.kernel.common.entity.po.HistoryRecord;
import com.example.kernel.common.mapper.HistoryRecordMapper;
import org.springframework.stereotype.Service;

@Service
public class HistoryRecordService extends ServiceImpl<HistoryRecordMapper, HistoryRecord> {
}
