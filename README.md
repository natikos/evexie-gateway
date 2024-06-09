# Evexie Gateway

Evexie Gateway is the central API gateway for the Evexie app, managing secure and efficient communication between frontend clients and backend microservices.

### Start RabbitMQ for communication between gateway and microservices

docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
