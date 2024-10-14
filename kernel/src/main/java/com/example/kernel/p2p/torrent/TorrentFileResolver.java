package com.example.kernel.p2p.torrent;

import com.alibaba.fastjson2.JSON;
import com.example.kernel.common.exceptionAdvice.DefinedException;
import com.example.kernel.common.util.BencodeUtils;
import com.example.kernel.common.util.FileUtils;
import com.example.kernel.p2p.entity.TorrentInfo;
import com.example.kernel.p2p.entity.Torrent;
import com.example.kernel.p2p.entity.TorrentInfoFiles;
import lombok.Data;
import lombok.extern.slf4j.Slf4j;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;

@Slf4j
@Data
public class TorrentFileResolver {

    private String directoryName;
    private Path path;
    private List<Path> paths;
    private Torrent torrent;
    private boolean isSingleFile = false;
    private byte[] torrentBytes;
    private int infoStart;
    private int infoEnd;

    private TorrentFileResolver() {
    }

    /**
     * 读取torrent文件
     */
    public TorrentFileResolver(String path) throws IOException {
        this.path = Paths.get(path).toAbsolutePath().normalize();
        this.torrentBytes = Files.readAllBytes(this.path);
        this.infoStart = 0;
        this.infoEnd = 0;
    }

    /**
     * 生成torrent文件
     */
    public TorrentFileResolver(List<String> paths) {
        this.paths = paths.stream().map(i -> Paths.get(i).toAbsolutePath().normalize()).toList();
        this.torrentBytes = encode();
        this.infoStart = 0;
        this.infoEnd = 0;
        this.torrent = new Torrent();
        this.isSingleFile = this.paths.size() == 1;
    }

    public Torrent read() {
        if (torrent != null) {
            return torrent;
        }
        String jsonString = JSON.toJSONString(decode(torrentBytes));
        torrent = JSON.parseObject(jsonString, Torrent.class);
        log.info("torrent info:{}", jsonString);
        byte[] infoByte = new byte[infoEnd - infoStart + 1];
        System.arraycopy(torrentBytes, infoStart, infoByte, 0, infoEnd - infoStart + 1);
        isSingleFile = torrent.getTorrentInfo().getFiles() == null;
        torrent.setInfoBencode(infoByte);
        torrent.setPath(path);
        torrent.hash();
        return torrent;
    }

    // Todo 生成torrent文件
    public void write(String path) throws IOException {
        Path normalizePath = Paths.get(path.endsWith(".torrent") ? path : path + ".torrent").toAbsolutePath().normalize();
        FileUtils.createFile(normalizePath);
        torrent.setPath(normalizePath);
        TorrentInfo torrentInfo = new TorrentInfo();
        if (isSingleFile){

        }else {

        }
    }

    private byte[] CalculatePieces() {
//        torrent.getin

        return null;
    }

    private Object decode(byte[] s) {
        Object o = null;
        if (s == null || s.length == 0) {
            return o;
        }
        try {
            o = decodeObject(s, 0)[0];
        } catch (Exception e) {
            throw new DefinedException(e);
        }
        return o;
    }

    private Object[] decodeObject(byte[] s, int index) {
        byte b = s[index];
        if (b == 'i') {
            // 整数类型
            index++;
            int start = index;
            while (s[index] != 'e') {
                index++;
            }
            long value = Long.parseLong(new String(s, start, index - start));
            index++;
            return new Object[]{value, index};
        } else if ('0' <= b && b <= '9') {
            // 字符串类型
            int start = index;
            while (s[index] != ':') {
                index++;
            }
            int length = Integer.parseInt(new String(s, start, index - start));
            index++;
            String str = new String(s, index, length);
            index += length;
            return new Object[]{str, index};
        } else if (b == 'd') {
            // 字典类型
            index++;
            Map<String, Object> map = new LinkedHashMap<>();
            while (s[index] != 'e') {
                Object[] keyResult = decodeObject(s, index);
                String key = (String) keyResult[0];
                boolean record = key.equals("info");
                if (record) {
                    infoStart = index + 6;
                }
                index = (int) keyResult[1];

                Object[] valueResult = decodeObject(s, index);
                Object value = valueResult[0];
                index = (int) valueResult[1];
                if (record) {
                    infoEnd = index - 1;
                }
                map.put(key, value);
            }
            index++;
            return new Object[]{map, index};
        } else if (b == 'l') {
            // 列表类型
            index++;
            List<Object> list = new ArrayList<>();
            while (s[index] != 'e') {
                Object[] result = decodeObject(s, index);
                list.add(result[0]);
                index = (int) result[1];
            }
            index++;
            return new Object[]{list, index};
        } else {
            // 忽略无效的字符
            index++;
            return decodeObject(s, index);
        }
    }

    // Todo 将torrent进行编码
    private byte[] encode() {
        HashMap<String, Object> torrent = new HashMap<>();
        torrent.put("announce", this.torrent.getAnnounce());
        torrent.put("announce-list", this.torrent.getAnnounce() == null ? null : this.torrent.getAnnounceList());

        HashMap<String, Object> info = new HashMap<>();
        info.put("name", this.torrent.getTorrentInfo().getName());
        info.put("piece length", this.torrent.getTorrentInfo().getPieceLength());
        info.put("pieces", this.torrent.getTorrentInfo().getPieces());
        if (this.isSingleFile) {
            info.put("length", this.torrent.getTorrentInfo().getLength());
        } else {
            List<HashMap<String, Object>> files = new ArrayList<>();
            for (TorrentInfoFiles file : this.torrent.getTorrentInfo().getFiles()) {
                HashMap<String, Object> fileInfo = new HashMap<>();
                fileInfo.put("length", file.getLength());
                fileInfo.put("path", file.getPath());
                files.add(fileInfo);
            }
            info.put("files", files);
        }
        torrent.put("info", info);

        torrent.put("comment", this.torrent.getComment() == null ? null : this.torrent.getComment());
        torrent.put("created by", this.torrent.getCreatedBy() == null ? null : this.torrent.getCreatedBy());
        torrent.put("creation date", this.torrent.getCreationDate() == null ? null : this.torrent.getCreationDate());
        torrent.put("encoding", this.torrent.getEncoding() == null ? null : this.torrent.getEncoding());
        torrent.put("publisher", this.torrent.getPublisher() == null ? null : this.torrent.getPublisher());
        torrent.put("publisher-url", this.torrent.getPublisherUrl() == null ? null : this.torrent.getPublisherUrl());
        torrent.put("url-list", this.torrent.getUrlList() == null ? null : this.torrent.getUrlList());
        return BencodeUtils.encode(torrent).getBytes();
    }

}
