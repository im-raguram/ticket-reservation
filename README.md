# ticket-reservation
Client - Server Ticket Reservation System coded using Golang and gRPC

To generate go code using proto file:

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/ticket.proto

How to Run:
1. Clone this repo
2. Open Postman and create a new gRPC request
3. Move to service definiton tab and import the proto file
4. Set the url as localhost:8080 and see all the imported functions
5. Move to message tab and click example message for a sample input and check the output