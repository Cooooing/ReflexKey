package com.example.kernel.p2p.entity;

import com.alibaba.fastjson2.annotation.JSONField;
import com.example.kernel.common.entity.base.Constant;
import com.example.kernel.common.util.DigestUtils;
import lombok.Data;
import lombok.experimental.Accessors;

import java.nio.charset.StandardCharsets;
import java.nio.file.Path;
import java.util.List;

@Data
@Accessors(chain = true)
public class Torrent {

    private String announce;
    @JSONField(name = "announce-list")
    private List<List<String>> announceList; // optional spare Tracker
    private TorrentInfo torrentInfo;

    // The following properties are optional
    private String comment;
    @JSONField(name = "created by")
    private String createdBy = Constant.NAME + " " + Constant.VERSION;
    @JSONField(name = "creation date")
    private Long creationDate = System.currentTimeMillis() / 1000L;
    private String encoding = StandardCharsets.UTF_8.name();

    private String publisher;
    @JSONField(name = "publisher-url")
    private String publisherUrl;
    @JSONField(name = "url-list")
    private List<String> urlList;


    @JSONField(serialize = false)
    private Path path;
    @JSONField(serialize = false)
    private byte[] infoBencode;
    @JSONField(serialize = false)
    private String infoHash;

    public void hash() {
        if (infoHash == null) {
            this.infoHash = DigestUtils.toHexString(DigestUtils.toSHA1(this.infoBencode));
        }
    }

    public long fileBlockNum() {
        return this.torrentInfo.getPieceLength() / 20L;
    }

}