package com.example.kernel.service;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.example.kernel.entity.po.GeneratePassword;
import com.example.kernel.mapper.GeneratePasswordMapper;
import org.springframework.stereotype.Service;

@Service
public class GeneratePasswordService extends ServiceImpl<GeneratePasswordMapper, GeneratePassword> {
}
