package com.example.kernel.entity;

import com.baomidou.mybatisplus.annotation.TableName;

import java.util.Date;

/**
 * Todo
 **/
@TableName("user")
public class Test {
    private Integer id;
    private String name;
    private Date createdAt;
}
