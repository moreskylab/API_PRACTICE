package in.altgr.backend_java.controller;

import in.altgr.backend_java.dto.ObjectCreate;
import in.altgr.backend_java.dto.ObjectUpdate;
import in.altgr.backend_java.model.ResponseObject;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.server.ResponseStatusException;

import java.util.*;
import java.util.concurrent.CopyOnWriteArrayList;

@RestController
@RequestMapping("/objects")
public class ObjectController {

    // Thread-safe in-memory repository mimicking your global Python list
    private final List<ResponseObject> objects = new CopyOnWriteArrayList<>(Arrays.asList(
        new ResponseObject("1", "Google Pixel 6 Pro", Map.of("color", "Cloudy White", "capacity", "128 GB")),
        new ResponseObject("2", "Apple iPhone 12 Mini, 256GB, Blue", null),
        new ResponseObject("3", "Apple iPhone 12 Pro Max", Map.of("color", "Cloudy White", "capacity GB", 512)),
        new ResponseObject("4", "Apple iPhone 11, 64GB", Map.of("price", 389.99, "color", "Purple")),
        new ResponseObject("5", "Samsung Galaxy Z Fold2", Map.of("price", 689.99, "color", "Brown")),
        new ResponseObject("6", "Apple AirPods", Map.of("generation", "3rd", "price", 120)),
        new ResponseObject("7", "Apple MacBook Pro 16", Map.of("year", 2019, "price", 1849.99, "CPU model", "Intel Core i9", "Hard disk size", "1 TB")),
        new ResponseObject("8", "Apple Watch Series 8", Map.of("Strap Colour", "Elderberry", "Case Size", "41mm")),
        new ResponseObject("9", "Beats Studio3 Wireless", Map.of("Color", "Red", "Description", "High-performance wireless noise cancelling headphones")),
        new ResponseObject("10", "Apple iPad Mini 5th Gen", Map.of("Capacity", "64 GB", "Screen size", 7.9)),
        new ResponseObject("11", "Apple iPad Mini 5th Gen", Map.of("Capacity", "254 GB", "Screen size", 7.9)),
        new ResponseObject("12", "Apple iPad Air", Map.of("Generation", "4th", "Price", "419.99", "Capacity", "64 GB")),
        new ResponseObject("13", "Apple iPad Air", Map.of("Generation", "4th", "Price", "519.99", "Capacity", "256 GB"))
    ));

    @GetMapping
    public List<ResponseObject> getManyObjects() {
        return objects;
    }

    // @GetMapping("/{id}")
    // public ResponseObject getOneObject(@PathVariable String id) {
    //     return objects.stream()
    //             .filter(item -> item.id().equals(id))
    //             .findFirst()
    //             .orElseThrow(() -> new ResponseStatusException(
    //                     HttpStatus.NOT_FOUND, "Object with id " + id + notFoundMessage(id)));
    // }

    @GetMapping("/{id}")
    public ResponseObject getOneObject(@PathVariable String id) {
        return objects.stream()
                .filter(item -> item.id().equals(id))
                .findFirst()
                .orElseThrow(() -> new ResponseStatusException(
                        HttpStatus.NOT_FOUND, "Object with id " + id + " not found"));
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public ResponseObject postOneObject(@RequestBody ObjectCreate payload) {
        // Find maximum integer ID to calculate the next string ID
        int maxId = objects.stream()
                .mapToInt(item -> Integer.parseInt(item.id()))
                .max()
                .orElse(0);
        String nextId = String.valueOf(maxId + 1);

        ResponseObject newObject = new ResponseObject(nextId, payload.name(), payload.data());
        objects.add(newObject);
        return newObject;
    }

    @PutMapping("/{id}")
    public ResponseObject putOneObject(@PathVariable String id, @RequestBody ObjectCreate payload) {
        for (int i = 0; i < objects.size(); i++) {
            if (objects.get(i).id().equals(id)) {
                ResponseObject updatedObject = new ResponseObject(id, payload.name(), payload.data());
                objects.set(i, updatedObject);
                return updatedObject;
            }
        }
        throw new ResponseStatusException(HttpStatus.NOT_FOUND, "Object with id " + id + notFoundMessage(id));
    }

    @PatchMapping("/{id}")
    public ResponseObject patchOneObject(@PathVariable String id, @RequestBody ObjectUpdate payload) {
        for (int i = 0; i < objects.size(); i++) {
            ResponseObject current = objects.get(i);
            if (current.id().equals(id)) {
                String updatedName = payload.name() != null ? payload.name() : current.name();
                
                // Partial patch merge matching the python payload.model_dump(exclude_unset=True) context
                Map<String, Object> updatedData = current.data() != null ? new LinkedHashMap<>(current.data()) : new LinkedHashMap<>();
                if (payload.data() != null) {
                    updatedData.putAll(payload.data());
                } else if (current.data() == null && payload.data() == null) {
                    updatedData = null; // Keeps structure equivalent if it was explicitly absent
                }

                ResponseObject patchedObject = new ResponseObject(id, updatedName, updatedData);
                objects.set(i, patchedObject);
                return patchedObject;
            }
        }
        throw new ResponseStatusException(HttpStatus.NOT_FOUND, "Object with id " + id + notFoundMessage(id));
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteOneObject(@PathVariable String id) {
        boolean removed = objects.removeIf(item -> item.id().equals(id));
        if (!removed) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "Object with id " + id + notFoundMessage(id));
        }
    }

    private String notFoundMessage(String id) {
        return " not found";
    }
}
