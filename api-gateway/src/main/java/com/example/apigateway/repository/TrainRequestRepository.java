package com.example.apigateway.repository;

import com.example.apigateway.model.TrainRequest;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface TrainRequestRepository extends MongoRepository<TrainRequest, String> {
}