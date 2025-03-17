package com.example.apigateway.controller;

import com.example.apigateway.model.TrainRequest;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.http.ResponseEntity;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@RestController
@RequestMapping("/api/v1/trains")
public class TrainController {

    private static final Logger logger = LoggerFactory.getLogger(TrainController.class);

    @PostMapping("/add")
    public ResponseEntity<String> addTrain(@RequestBody TrainRequest request) {
        logger.info("Received request to add train: {}", request.getTrainName());
        // Simulate sending the request to the Task Scheduler (Golang) or saving it to the database
        return ResponseEntity.ok("Train added successfully: " + request.getTrainName());
    }

    @GetMapping("/status")
    public ResponseEntity<String> getTrainStatus() {
        logger.info("Status check request received");
        // Simulate status response
        return ResponseEntity.ok("Train is operational and running smoothly.");
    }
}
