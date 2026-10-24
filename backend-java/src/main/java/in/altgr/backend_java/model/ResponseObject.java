package in.altgr.backend_java.model;

import java.util.Map;

// The base model matching your objects structure
public record ResponseObject(
    String id,
    String name,
    Map<String, Object> data
) {}
