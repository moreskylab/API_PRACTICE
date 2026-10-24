package controller

import (
	"net/http"
	"strconv"
	"sync"

	"backend-go/dto"
	"backend-go/model"

	"github.com/gin-gonic/gin"
)

type ObjectController struct {
	mu      sync.RWMutex // Protects the slice during parallel requests
	objects []model.ResponseObject
}

func NewObjectController() *ObjectController {
	return &ObjectController{
		objects: []model.ResponseObject{
			{ID: "1", Name: "Google Pixel 6 Pro", Data: map[string]interface{}{"color": "Cloudy White", "capacity": "128 GB"}},
			{ID: "2", Name: "Apple iPhone 12 Mini, 256GB, Blue", Data: nil},
			{ID: "3", Name: "Apple iPhone 12 Pro Max", Data: map[string]interface{}{"color": "Cloudy White", "capacity GB": float64(512)}},
			{ID: "4", Name: "Apple iPhone 11, 64GB", Data: map[string]interface{}{"price": 389.99, "color": "Purple"}},
			{ID: "5", Name: "Samsung Galaxy Z Fold2", Data: map[string]interface{}{"price": 689.99, "color": "Brown"}},
			{ID: "6", Name: "Apple AirPods", Data: map[string]interface{}{"generation": "3rd", "price": float64(120)}},
			{ID: "7", Name: "Apple MacBook Pro 16", Data: map[string]interface{}{"year": float64(2019), "price": 1849.99, "CPU model": "Intel Core i9", "Hard disk size": "1 TB"}},
			{ID: "8", Name: "Apple Watch Series 8", Data: map[string]interface{}{"Strap Colour": "Elderberry", "Case Size": "41mm"}},
			{ID: "9", Name: "Beats Studio3 Wireless", Data: map[string]interface{}{"Color": "Red", "Description": "High-performance wireless noise cancelling headphones"}},
			{ID: "10", Name: "Apple iPad Mini 5th Gen", Data: map[string]interface{}{"Capacity": "64 GB", "Screen size": 7.9}},
			{ID: "11", Name: "Apple iPad Mini 5th Gen", Data: map[string]interface{}{"Capacity": "254 GB", "Screen size": 7.9}},
			{ID: "12", Name: "Apple iPad Air", Data: map[string]interface{}{"Generation": "4th", "Price": "419.99", "Capacity": "64 GB"}},
			{ID: "13", Name: "Apple iPad Air", Data: map[string]interface{}{"Generation": "4th", "Price": "519.99", "Capacity": "256 GB"}},
		},
	}
}

func (oc *ObjectController) GetManyObjects(c *gin.Context) {
	oc.mu.RLock()
	defer oc.mu.RUnlock()
	c.JSON(http.StatusOK, oc.objects)
}

func (oc *ObjectController) GetOneObject(c *gin.Context) {
	id := c.Param("id")

	oc.mu.RLock()
	defer oc.mu.RUnlock()

	for _, item := range oc.objects {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"detail": "Object with id " + id + " not found"})
}

func (oc *ObjectController) PostOneObject(c *gin.Context) {
	var payload dto.ObjectCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	oc.mu.Lock()
	defer oc.mu.Unlock()

	maxID := 0
	for _, item := range oc.objects {
		if val, err := strconv.Atoi(item.ID); err == nil && val > maxID {
			maxID = val
		}
	}
	nextID := strconv.Itoa(maxID + 1)

	newObject := model.ResponseObject{
		ID:   nextID,
		Name: payload.Name,
		Data: payload.Data,
	}

	oc.objects = append(oc.objects, newObject)
	c.JSON(http.StatusCreated, newObject)
}

func (oc *ObjectController) PutOneObject(c *gin.Context) {
	id := c.Param("id")
	var payload dto.ObjectCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	oc.mu.Lock()
	defer oc.mu.Unlock()

	for i, item := range oc.objects {
		if item.ID == id {
			updatedObject := model.ResponseObject{
				ID:   id,
				Name: payload.Name,
				Data: payload.Data,
			}
			oc.objects[i] = updatedObject
			c.JSON(http.StatusOK, updatedObject)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"detail": "Object with id " + id + " not found"})
}

func (oc *ObjectController) PatchOneObject(c *gin.Context) {
	id := c.Param("id")
	var payload dto.ObjectUpdate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	oc.mu.Lock()
	defer oc.mu.Unlock()

	for i, item := range oc.objects {
		if item.ID == id {
			if payload.Name != nil {
				oc.objects[i].Name = *payload.Name
			}
			if payload.Data != nil {
				if oc.objects[i].Data == nil {
					oc.objects[i].Data = make(map[string]interface{})
				}
				for k, v := range payload.Data {
					oc.objects[i].Data[k] = v
				}
			}
			c.JSON(http.StatusOK, oc.objects[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"detail": "Object with id " + id + " not found"})
}

func (oc *ObjectController) DeleteOneObject(c *gin.Context) {
	id := c.Param("id")

	oc.mu.Lock()
	defer oc.mu.Unlock()

	for i, item := range oc.objects {
		if item.ID == id {
			oc.objects = append(oc.objects[:i], oc.objects[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"detail": "Object with id " + id + " not found"})
}
