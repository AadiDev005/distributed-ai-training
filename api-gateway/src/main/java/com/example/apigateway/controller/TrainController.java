package com.example.apigateway.controller;

import com.example.apigateway.model.TrainRequest;
import com.example.apigateway.repository.TrainRequestRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.SendResult;
import org.springframework.web.bind.annotation.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.validation.Valid;
import java.util.concurrent.CompletableFuture;

@RestController
@RequestMapping("/api/v1/trains")
public class TrainController {

    private static final Logger logger = LoggerFactory.getLogger(TrainController.class);
    private static final String TOPIC = "train-tasks";  // Matches docker-compose.yml and application.properties

    @Autowired
    private TrainRequestRepository repository;

    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    @PostMapping("/add")
    public ResponseEntity<String> addTrain(@Valid @RequestBody TrainRequest request) {
        logger.info("Received request to add train: {}", request.getTrainName());
        try {
            if (request.getTrainName() == null || request.getTrainName().isEmpty()) {
                throw new IllegalArgumentException("Train name cannot be null or empty");
            }

            logger.debug("Setting status to PENDING for train: {}", request.getTrainName());
            request.setStatus("PENDING");

            logger.debug("Attempting to save train request to MongoDB: {}", request);
            TrainRequest savedRequest = repository.save(request);
            logger.info("Successfully saved train request to MongoDB with ID: {}", savedRequest.getId());

            String message = String.format("Train: %s, Source: %s, Destination: %s, Seats: %d, ID: %s",
                    request.getTrainName(), request.getSource(), request.getDestination(),
                    request.getNumberOfSeats(), savedRequest.getId());
            logger.debug("Attempting to send message to Kafka topic '{}': {}", TOPIC, message);
            CompletableFuture<SendResult<String, String>> future = kafkaTemplate.send(TOPIC, message);
            future.whenComplete((result, ex) -> {
                if (ex == null) {
                    logger.info("Train task sent to Kafka: {} (Partition: {}, Offset: {})",
                            message, result.getRecordMetadata().partition(), result.getRecordMetadata().offset());
                } else {
                    logger.error("Failed to send message to Kafka topic '{}': {}", TOPIC, ex.getMessage());
                }
            });
            logger.debug("Kafka send initiated for message: {}", message);

            return ResponseEntity.ok("Train added successfully: " + request.getTrainName());
        } catch (Exception e) {
            logger.error("Failed to process train request '{}': {}", request.getTrainName(), e.getMessage(), e);
            return ResponseEntity.status(500).body("Failed to add train: " + e.getMessage());
        }
    }

    @GetMapping("/status")
    public ResponseEntity<String> getTrainStatus() {
        logger.info("Status check request received");
        // Placeholder; could query MongoDB for real status later
        return ResponseEntity.ok("Train is operational and running smoothly.");
    }
}