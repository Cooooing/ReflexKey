package com.example.kernel.util;

import lombok.extern.slf4j.Slf4j;
import org.springframework.util.CollectionUtils;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.UUID;
import java.util.concurrent.ThreadLocalRandom;
import java.util.stream.Stream;

@Slf4j
public final class RandomUtils {

    private static final List<Character> lowerCaseChars = List.of('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z');
    private static final List<Character> upperCaseChars = List.of('A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z');
    private static final List<Character> numbers = List.of('0', '1', '2', '3', '4', '5', '6', '7', '8', '9');
    private static final List<Character> specialChars = List.of('!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', ']', '{', '}', '|', '\\', ';', ':', '\'', '\"', ',', '<', '.', '>', '/', '?');

    private static final List<Character> shortUuidChars = Stream.of(lowerCaseChars, upperCaseChars, numbers).flatMap(List::stream).toList();
    private static final int shortUuidCharsLengthHex = Integer.parseInt(Integer.toHexString(shortUuidChars.size()), 16);

    public static String generateUuid() {
        String ticket = UUID.randomUUID().toString();
        return ticket.replaceAll("-", "");
    }

    public static String generateShortUuid() {
        StringBuilder shortBuffer = new StringBuilder();
        String uuid = RandomUtils.generateUuid();
        for (int i = 0; i < 8; i++) {
            String str = uuid.substring(i * 4, i * 4 + 4);
            int x = Integer.parseInt(str, 16);
            shortBuffer.append(shortUuidChars.get(x % shortUuidCharsLengthHex));
        }
        return shortBuffer.toString();
    }

    public static String generatePassword(int length, boolean hasLowerCase, boolean hasUpperCase, boolean hasNumber, boolean hasSpecialChar, List<Character> includeChars, List<Character> excludeChars) {
        StringBuilder password = new StringBuilder();
        Stream<Character> stream = Stream.empty();
        stream = hasLowerCase ? Stream.concat(stream, lowerCaseChars.stream()) : stream;
        stream = hasUpperCase ? Stream.concat(stream, upperCaseChars.stream()) : stream;
        stream = hasNumber ? Stream.concat(stream, numbers.stream()) : stream;
        stream = hasSpecialChar ? Stream.concat(stream, specialChars.stream()) : stream;
        stream = CollectionUtils.isEmpty(includeChars) ? stream : Stream.concat(stream, includeChars.stream());
        stream = CollectionUtils.isEmpty(excludeChars) ? stream : stream.filter(c -> !excludeChars.contains(c));
        List<Character> characters = new ArrayList<>(stream.distinct().toList());
        if (length < 1 || CollectionUtils.isEmpty(characters)) {
            return password.toString();
        }
        Collections.shuffle(characters);
        for (int i = 0; i < length; i++) {
            int randomIndex = ThreadLocalRandom.current().nextInt(characters.size());
            password.append(characters.get(randomIndex));
        }
        return password.toString();
    }
}
