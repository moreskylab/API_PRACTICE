from fastapi import FastAPI, HTTPException, status, Response
from fastapi.middleware.cors import CORSMiddleware
from typing import Optional, Any
from pydantic import BaseModel, Field

app = FastAPI()

# Define the domains allowed to access your API
origins = [
    "http://172.30.97.48:8888",
    "http://localhost:80",
    "https://es-d-97718351720261025-01a16668-6a61-7815-b0a1-c8323162d51b.codepen.dev" # Your production domain
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,            # Allows specific domains to connect
    # allow_origins=["*"],            # Use this to allow ANY domain (not recommended for production)
    allow_credentials=True,
    allow_methods=["*"],        
          # Allows all HTTP methods (GET, POST, etc.)
    allow_headers=["*"],              # Allows all headers
)


objects = [
    {"id":"1","name":"Google Pixel 6 Pro","data":{"color":"Cloudy White","capacity":"128 GB"}},
    {"id":"2","name":"Apple iPhone 12 Mini, 256GB, Blue","data":None},
    {"id":"3","name":"Apple iPhone 12 Pro Max","data":{"color":"Cloudy White","capacity GB":512}},
    {"id":"4","name":"Apple iPhone 11, 64GB","data":{"price":389.99,"color":"Purple"}},
    {"id":"5","name":"Samsung Galaxy Z Fold2","data":{"price":689.99,"color":"Brown"}},
    {"id":"6","name":"Apple AirPods","data":{"generation":"3rd","price":120}},
    {"id":"7","name":"Apple MacBook Pro 16","data":{"year":2019,"price":1849.99,"CPU model":"Intel Core i9","Hard disk size":"1 TB"}},
    {"id":"8","name":"Apple Watch Series 8","data":{"Strap Colour":"Elderberry","Case Size":"41mm"}},
    {"id":"9","name":"Beats Studio3 Wireless","data":{"Color":"Red","Description":"High-performance wireless noise cancelling headphones"}},
    {"id":"10","name":"Apple iPad Mini 5th Gen","data":{"Capacity":"64 GB","Screen size":7.9}},
    {"id":"11","name":"Apple iPad Mini 5th Gen","data":{"Capacity":"254 GB","Screen size":7.9}},
    {"id":"12","name":"Apple iPad Air","data":{"Generation":"4th","Price":"419.99","Capacity":"64 GB"}},
    {"id":"13","name":"Apple iPad Air","data":{"Generation":"4th","Price":"519.99","Capacity":"256 GB"}}
]

@app.get("/objects")
def get_many_objects():
    return objects

@app.get("/objects/{id}")
def get_one_object(id: str):
    obj = next((item for item in objects if item["id"] == id), None)
    
    if obj is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, 
            detail=f"Object with id {id} not found"
        )
    return obj


class ObjectCreate(BaseModel):
    name: str
    data: dict[str, Any] = None

@app.post("/objects", status_code=status.HTTP_201_CREATED)
def post_one_object(payload: ObjectCreate):
    # Calculate the next ID based on your existing string IDs
    next_id = str(max([int(item["id"]) for item in objects]) + 1) if objects else "1"
    
    # Construct the new item
    new_object = {
        "id": next_id,
        "name": payload.name,
        "data": payload.data
    }
    
    objects.append(new_object)
    return new_object

@app.put("/objects/{id}", status_code=status.HTTP_200_OK)
def put_one_object(id: str, payload: ObjectCreate):
    # Find the object index in the list
    for index, item in enumerate(objects):
        if item["id"] == id:
            # Replace the object entirely (keeping the original ID)
            updated_object = {
                "id": id,
                "name": payload.name,
                "data": payload.data
            }
            objects[index] = updated_object
            return updated_object
            
    # If the ID wasn't found, raise a 404 error
    raise HTTPException(
        status_code=status.HTTP_404_NOT_FOUND,
        detail=f"Object with id {id} not found"
    )


class ObjectUpdate(BaseModel):
    name: str | None = None
    data: dict[str, Any] = None

@app.patch("/objects/{id}", status_code=status.HTTP_200_OK)
def patch_one_object(id: str, payload: ObjectUpdate):
    # Find the object in your list
    obj = next((item for item in objects if item["id"] == id), None)
    
    if obj is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Object with id {id} not found"
        )
    
    # Extract only the fields explicitly sent by the user
    update_data = payload.model_dump(exclude_unset=True)
    
    if "name" in update_data:
        obj["name"] = update_data["name"]
        
    if "data" in update_data:
        # If the existing object data is None, initialize it as a dictionary
        if obj["data"] is None:
            obj["data"] = {}
            
        # Perform a partial merge on the nested 'data' dictionary
        if isinstance(update_data["data"], dict):
            obj["data"].update(update_data["data"])
        else:
            obj["data"] = update_data["data"]

    return obj


@app.delete("/objects/{id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_one_object(id: str):
    # Find the object in your list
    for index, item in enumerate(objects):
        if item["id"] == id:
            # Remove the object from the in-memory array and return it
            deleted_object = objects.pop(index)
            return Response(status_code=status.HTTP_204_NO_CONTENT)
            
    # If the ID wasn't found, raise a 404 error
    raise HTTPException(
        status_code=status.HTTP_404_NOT_FOUND,
        detail=f"Object with id {id} not found"
    )