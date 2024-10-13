package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "zoo-backend/models"
    "zoo-backend/services"
    "log" 
    "fmt" 
)

type ZooController struct {
    Service *services.ZooService
}

func (c *ZooController) CreateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo

    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input format"})
        log.Printf("CreateZoo failed: Invalid input format. Error: %v", err) 
        return
    }

    if zoo.Name == "" || zoo.Class == "" || zoo.Legs <= 0 {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid data provided"})
        log.Printf("CreateZoo failed: Invalid data provided. Received: %+v", zoo) 
        return
    }

    id, err := c.Service.CreateZoo(zoo)
    if err != nil {
        if err.Error() == fmt.Sprintf("zoo with ID '%d' already exists", zoo.ID) {
            w.WriteHeader(http.StatusConflict) 
            json.NewEncoder(w).Encode(map[string]string{"error": "Zoo with this ID already exists"})
            log.Printf("CreateZoo conflict: Zoo with ID '%d' already exists", zoo.ID) 
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create zoo"})
            log.Printf("CreateZoo failed: %v", err) 
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Successfully created zoo",
        "id":      id,
    })
    log.Printf("CreateZoo succeeded: Zoo created with ID %d", id) 
}


func (c *ZooController) GetAllZoos(w http.ResponseWriter, r *http.Request) {
    zoos, err := c.Service.GetAllZoos()
    if err != nil {
        http.Error(w, "Failed to get zoos", http.StatusInternalServerError)
        log.Printf("GetAllZoos failed: %v", err) 
        return
    }

    if len(zoos) == 0 {
        http.Error(w, "No zoos found", http.StatusNotFound)
        log.Printf("GetAllZoos failed: No zoos found") 
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoos)

    log.Printf("GetAllZoos succeeded: Returned %d zoos with status %d", len(zoos), http.StatusOK)
}


func (c *ZooController) GetZooByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    zoo, err := c.Service.GetZooByID(id)
    if err != nil {
        http.Error(w, "Zoo not found", http.StatusNotFound)
        log.Printf("GetZooByID failed: Zoo with ID %d not found", id) 
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(zoo)
    log.Printf("GetZooByID succeeded: Returned zoo with ID %d", id) 
}

func (c *ZooController) UpdateZoo(w http.ResponseWriter, r *http.Request) {
    var zoo models.Zoo

    err := json.NewDecoder(r.Body).Decode(&zoo)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        log.Printf("UpdateZoo failed: Invalid input format. Error: %v", err) 
        return
    }

    if zoo.ID == 0 {
        http.Error(w, "Missing Zoo ID", http.StatusBadRequest)
        log.Printf("UpdateZoo failed: Missing Zoo ID") 
        return
    }

    updated, err := c.Service.UpsertZoo(zoo)
    if err != nil {
        http.Error(w, "Failed to update or create zoo", http.StatusInternalServerError)
        log.Printf("UpsertZoo failed: %v", err) 
        return
    }

    var message string
    if updated {
        message = fmt.Sprintf("Zoo with ID %d was successfully updated", zoo.ID)
    } else {
        message = fmt.Sprintf("Zoo with ID %d not found and successfully created", zoo.ID)
    }

    response := map[string]string{
        "message": message,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
    log.Printf("UpsertZoo succeeded: %s", message) 
}



func (c *ZooController) DeleteZoo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        log.Printf("DeleteZoo failed: Invalid ID format. Status: %d", http.StatusBadRequest)
        return
    }

    err = c.Service.DeleteZoo(id)
    if err != nil {
        http.Error(w, "Zoo not found", http.StatusNotFound)
        log.Printf("DeleteZoo failed: Zoo with ID %d not found. Status: %d", id, http.StatusNotFound)
        return
    }

    response := map[string]string{
        "message": "Zoo deleted successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)

    log.Printf("DeleteZoo succeeded: Zoo with ID %d deleted successfully. Status: %d", id, http.StatusOK)
}
