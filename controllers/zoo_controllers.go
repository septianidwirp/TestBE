package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "zoo-backend/models"
    "zoo-backend/services"
    "log" 
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
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create zoo"})
        log.Printf("CreateZoo failed: %v", err) 
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
        zoos = []models.Zoo{}
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

    
    err = c.Service.UpdateZoo(zoo)
    if err != nil {
        http.Error(w, "Failed to update zoo", http.StatusInternalServerError)
        log.Printf("UpdateZoo failed: %v", err) 
        return
    }

  
    response := map[string]string{
        "message": "Zoo updated successfully",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
    log.Printf("UpdateZoo succeeded: Updated zoo with ID %d", zoo.ID) // Logging
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
