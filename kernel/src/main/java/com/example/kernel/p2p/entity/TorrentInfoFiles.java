package com.example.kernel.p2p.entity;

import lombok.Data;
import lombok.experimental.Accessors;

import java.util.List;

@Data
@Accessors(chain = true)
public class TorrentInfoFiles {
    private List<String> path;
    private long length;
}
