package in.altgr.backend_java.dto;

import java.util.Map;

// Maps to ObjectCreate (name is mandatory)
public record ObjectCreate(
    String name,
    Map<String, Object> data
) {}
