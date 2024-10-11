package com.example.kernel.service;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.example.kernel.entity.po.HistoryRecord;
import com.example.kernel.mapper.HistoryRecordMapper;
import org.springframework.stereotype.Service;

@Service
public class HistoryRecordService extends ServiceImpl<HistoryRecordMapper, HistoryRecord> {
}
