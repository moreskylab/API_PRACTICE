package in.altgr.backend_java.dto;

import java.util.Map;

// Maps to ObjectUpdate (fields can be missing/null)
public record ObjectUpdate(
    String name,
    Map<String, Object> data
) {}
