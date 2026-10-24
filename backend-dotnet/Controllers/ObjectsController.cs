using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System.Text.Json.Serialization;

namespace YourApiNamespace.Controllers;

[ApiController]
[Route("[controller]")] // Maps automatically to /objects
public class ObjectsController : ControllerBase
{
    // Static in-memory data store shared across request instances
    private static readonly List<PhoneObject> Objects = new()
    {
        new("1", "Google Pixel 6 Pro", new Dictionary<string, object> { { "color", "Cloudy White" }, { "capacity", "128 GB" } }),
        new("2", "Apple iPhone 12 Mini, 256GB, Blue", null),
        new("3", "Apple iPhone 12 Pro Max", new Dictionary<string, object> { { "color", "Cloudy White" }, { "capacity GB", 512 } }),
        new("4", "Apple iPhone 11, 64GB", new Dictionary<string, object> { { "price", 389.99 }, { "color", "Purple" } }),
        new("5", "Samsung Galaxy Z Fold2", new Dictionary<string, object> { { "price", 689.99 }, { "color", "Brown" } }),
        new("6", "Apple AirPods", new Dictionary<string, object> { { "generation", "3rd" }, { "price", 120 } }),
        new("7", "Apple MacBook Pro 16", new Dictionary<string, object> { { "year", 2019 }, { "price", 1849.99 }, { "CPU model", "Intel Core i9" }, { "Hard disk size", "1 TB" } }),
        new("8", "Apple Watch Series 8", new Dictionary<string, object> { { "Strap Colour", "Elderberry" }, { "Case Size", "41mm" } }),
        new("9", "Beats Studio3 Wireless", new Dictionary<string, object> { { "Color", "Red" }, { "Description", "High-performance wireless noise cancelling headphones" } }),
        new("10", "Apple iPad Mini 5th Gen", new Dictionary<string, object> { { "Capacity", "64 GB" }, { "Screen size", 7.9 } }),
        new("11", "Apple iPad Mini 5th Gen", new Dictionary<string, object> { { "Capacity", "254 GB" }, { "Screen size", 7.9 } }),
        new("12", "Apple iPad Air", null) // Handles cut-off example gracefully
    };

    // Helper method to emulate FastAPI's HTTPException 404 response structure
    private ObjectResult ObjectNotFound(string id) =>
        Problem(detail: $"Object with id {id} not found", statusCode: StatusCodes.Status404NotFound);

    // GET /objects
    [HttpGet]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(IEnumerable<PhoneObject>))]
    public IActionResult GetManyObjects()
    {
        return Ok(Objects);
    }

    // GET /objects/{id}
    [HttpGet("{id}")]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(PhoneObject))]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    public IActionResult GetOneObject(string id)
    {
        var obj = Objects.FirstOrDefault(o => o.Id == id);
        return obj is not null ? Ok(obj) : ObjectNotFound(id);
    }

    // POST /objects
    [HttpPost]
    [ProducesResponseType(StatusCodes.Status201Created, Type = typeof(PhoneObject))]
    public IActionResult PostOneObject([FromBody] ObjectCreate payload)
    {
        var nextId = Objects.Any()
            ? (Objects.Max(o => int.TryParse(o.Id, out var val) ? val : 0) + 1).ToString()
            : "1";

        var newObj = new PhoneObject
        {
            Id = nextId,
            Name = payload.Name,
            Data = payload.Data is not null ? new Dictionary<string, object>(payload.Data) : null
        };

        Objects.Add(newObj);
        return CreatedAtAction(nameof(GetOneObject), new { id = nextId }, newObj);
    }

    // PUT /objects/{id}
    [HttpPut("{id}")]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(PhoneObject))]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    public IActionResult PutOneObject(string id, [FromBody] ObjectCreate payload)
    {
        var index = Objects.FindIndex(o => o.Id == id);
        if (index == -1) return ObjectNotFound(id);

        var updatedObj = new PhoneObject
        {
            Id = id,
            Name = payload.Name,
            Data = payload.Data is not null ? new Dictionary<string, object>(payload.Data) : null
        };

        Objects[index] = updatedObj;
        return Ok(updatedObj);
    }

    // PATCH /objects/{id}
    [HttpPatch("{id}")]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(PhoneObject))]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    public IActionResult PatchOneObject(string id, [FromBody] ObjectUpdate payload)
    {
        var obj = Objects.FirstOrDefault(o => o.Id == id);
        if (obj is null) return ObjectNotFound(id);

        if (payload.Name is not null)
        {
            obj.Name = payload.Name;
        }

        if (payload.Data is not null)
        {
            obj.Data ??= new Dictionary<string, object>();
            foreach (var kvp in payload.Data)
            {
                obj.Data[kvp.Key] = kvp.Value;
            }
        }

        return Ok(obj);
    }

    // DELETE /objects/{id}
    [HttpDelete("{id}")]
    [ProducesResponseType(StatusCodes.Status204NoContent)]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    public IActionResult DeleteOneObject(string id)
    {
        var index = Objects.FindIndex(o => o.Id == id);
        if (index == -1) return ObjectNotFound(id);

        Objects.RemoveAt(index);
        return NoContent();
    }
}

// Models & Request Body Objects (DTOs)
public class PhoneObject
{
    public PhoneObject() { }

    public PhoneObject(string id, string name, Dictionary<string, object>? data)
    {
        Id = id;
        Name = name;
        Data = data;
    }

    [JsonPropertyName("id")] public string Id { get; set; } = string.Empty;
    [JsonPropertyName("name")] public string Name { get; set; } = string.Empty;
    [JsonPropertyName("data")] public Dictionary<string, object>? Data { get; set; }
}

public record ObjectCreate(
    [property: JsonPropertyName("name")] string Name,
    [property: JsonPropertyName("data")] Dictionary<string, object>? Data = null
);

public record ObjectUpdate(
    [property: JsonPropertyName("name")] string? Name = null,
    [property: JsonPropertyName("data")] Dictionary<string, object>? Data = null
);
