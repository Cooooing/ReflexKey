package com.example.kernel.p2p.entity;

import com.alibaba.fastjson2.annotation.JSONField;
import com.example.kernel.common.entity.base.Constant;
import com.example.kernel.common.util.DigestUtils;
import com.example.kernel.p2p.torrent.TorrentFileResolver;
import lombok.Data;
import lombok.experimental.Accessors;

import java.util.List;

@Data
@Accessors(chain = true)
public class Torrent {

    private String announce;
    @JSONField(name = "announce-list")
    private List<List<String>> announceList; // optional spare Tracker
    private Info info;

    @Data
    @Accessors(chain = true)
    public class Info {
        private String name;
        @JSONField(name = "piece length")
        private long pieceLength;
        private String pieces; // SHA-1 hash of each piece. 20 bytes per piece
        // There are single-file and multi-file. length exists as a single file and files exist as multiple files, but not both or neither
        private long length;
        private List<Files> files;
    }

    @Data
    @Accessors(chain = true)
    public class Files {
        private List<String> path;
        private long length;
    }

    // The following properties are optional
    private String comment;
    @JSONField(name = "created by")
    private String createdBy;
    @JSONField(name = "creation date")
    private Long creationDate;
    private String encoding;

    private String publisher;
    @JSONField(name = "publisher-url")
    private String publisherUrl;
    @JSONField(name = "url-list")
    private List<String> urlList;


    @JSONField(serialize = false)
    private String path;
    @JSONField(serialize = false)
    private byte[] infoBencode;
    @JSONField(serialize = false)
    private String infoHash;


    public Torrent() {
        this.createdBy = Constant.NAME + " " + Constant.VERSION;
        this.creationDate = System.currentTimeMillis() / 1000L;
        this.encoding = "UTF-8";
    }

    public static Torrent buildWithPath(String path) throws Exception {
        TorrentFileResolver resolver = TorrentFileResolver.init(path);
        Torrent torrent = resolver.read().setPath(path);
        torrent.hash();
        return torrent;
    }

    private void hash() {
        if (infoHash == null) {
            this.infoHash = DigestUtils.toSHA1(this.infoBencode);
        }
    }

    public long fileBlockNum() {
        return this.info.pieceLength / 20L;
    }

}