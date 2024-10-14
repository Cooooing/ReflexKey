package com.example.kernel.p2p.entity;

import com.alibaba.fastjson2.annotation.JSONField;
import lombok.Data;
import lombok.experimental.Accessors;

import java.util.List;

@Data
@Accessors(chain = true)
public class TorrentInfo {
    private String name;
    @JSONField(name = "piece length")
    private long pieceLength; // The length of each piece (usually to the power of 2, for example, 256 KB, 512 KB, etc.)
    private String pieces; // SHA-1 hash of each piece. 20 bytes per piece
    // There are single-file and multi-file. length exists as a single file and files exist as multiple files, but not both or neither
    private Long length;
    private List<TorrentInfoFiles> files;
}