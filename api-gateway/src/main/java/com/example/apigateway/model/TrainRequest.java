package com.example.apigateway.model;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Data
@Document(collection = "train_requests")
public class TrainRequest {
    @Id
    private String id; // Added for MongoDB primary key
    @Setter
    @Getter
    private String trainName;
    @Setter
    @Getter
    private String source;
    @Setter
    @Getter
    private String destination;
    @Setter
    @Getter
    private int numberOfSeats;
    @Setter
    @Getter
    private String status; // Added to track task status (e.g., PENDING, PROCESSING, COMPLETED)

}