# Use OpenJDK as the base image
FROM eclipse-temurin:17-jdk

# Set the working directory inside the container
WORKDIR /app

# Copy the JAR file from the target directory
COPY target/demo-0.0.1-SNAPSHOT.jar api-gateway.jar

# Expose the port the app runs on
EXPOSE 8080

# Run the application
ENTRYPOINT ["java", "-jar", "api-gateway.jar"]
