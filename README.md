# gRPC <> Kafka <> Docker

This project is a demo that shows how to connect gRPC GO modules to a Kafka broker running in Docker.

It also goes one step ahead of trying to make a comparision between the performance of REST(HTTP) and gRPC(TCP).

### Folder structure and what they do:
##### client
* Defines a ```grpc_client.go``` which handles all the gRPC communication to server
* Defines a ```http_client.go``` which handles all the HTTP communication to server
 
##### clientstubs
*  Defines a `http - json` stub and a `grpc - protobuf` stub

##### kafka_testing
* Send and receive str messages into and from a Kafka broker hosted by docker.

##### performance_analysis
* Contains a `grpc` and a `http` analyzer. These are utility files to send messages to Kafka using the desired 
protocol and retrieve those messages using the smae protocol.
* They also print out a short summary of operation in progress and some time analysis on how long it took.

##### server
* Contains a `http_server` (Kafka Backend)
* Contains a `in_memory_server` (No Kafka backend)
* Contains a `grpc_server` (Kafka Backend)
* These three utilities are easily exposed using a `main.go` which accepts flags.
```bash
    > pwd
        ~/gRPC-hands-on/server
      
    > go build -o server main.go
    > ./server --backend=[grpc_kafka|memory|http_kafka]
``` 

##### user
* Contains the protocol buffer definition for the model: `User`
* Contains a `JSON` equivalent definition for HTTP  
To convert the `user.proto` to a .pb.go file  
```bash
    > pwd
        ~/gRPC-hands-on
        
    > go get -u github.com/golang/protobuf/protoc-gen-go
    > protoc -I user user/user.proto --go_out=plugins=grpc:user
```

##### user
* General utils and Kafka common vars



### Test Runs
#### Start Kafka Broker
```bash
    > pwd
        ~/gRPC-hands-on/server
            
    > docker-compose up
        # Starts zookeeper(2181) and kafka(9092) on respective ports    
```

#### gRPC analysis
* Start the gRPC server
```bash
    > pwd
        ~/gRPC-hands-on/server
        
    > ./server --backend=grpc_kafka
        Starting gRPC server
```

* Start the gRPC analyzer which intern invokes the gRPC client
```bash
    > pwd
        ~/gRPC-hands-on/performance_analysis
        
    > go build -o grpc_performance_analyzer grpc_performance_analyzer.go
    > ./grpc_performance_analyzer
``` 

###### Results
```bash
I0801 17:25:08.597790   77864 grpc_performance_analyzer.go:51] ---------- Starting gRPC Kafka performance test ----------
I0801 17:25:08.597845   77864 grpc_performance_analyzer.go:52]  >>>>> Test to create 100 users <<<<<
I0801 17:26:49.523614   77864 grpc_performance_analyzer.go:22] .. 100 users created
I0801 17:26:49.523621   77864 grpc_performance_analyzer.go:27] Time taken to create 100 records via gRPC<>Kafka: 1m40.925185697s
I0801 17:26:49.523654   77864 grpc_performance_analyzer.go:54]  >>>>> Test to get 100 users <<<<<
I0801 17:26:58.831950   77864 grpc_client.go:39] Total records retrieved: 223
I0801 17:26:58.831982   77864 grpc_performance_analyzer.go:37] Time taken to read 100 records via gRPC<>Kafka: 4.308257619s

```

#### HTTP analysis
* Start the HTTP server
```bash
    > pwd
        ~/gRPC-hands-on/server
        
    > ./server --backend=http_kafka
        Starting HTTP server
```

* Start the HTTP analyzer which intern invokes the HTTP client
```bash
    > pwd
        ~/gRPC-hands-on/performance_analysis
        
    > go build -o http_performance_analyzer http_performance_analyzer.go
    > ./http_performance_analyzer
``` 

###### Results
```bash
I0801 17:16:47.430956   77864 http_performance_analyzer.go:39] ---------- Starting HTTP Kafka performance test ----------
I0801 17:16:47.430959   77864 http_performance_analyzer.go:40]  >>>>> Test to create 100 users <<<<<
I0801 17:16:47.430962   77695 http_performance_analyzer.go:18] .. 100 users created
I0801 17:16:47.430974   77695 http_performance_analyzer.go:23] Time taken to create 100 records via HTTP<>Kafka: 1m40.8536612s
I0801 17:16:47.431005   77695 http_performance_analyzer.go:42]  >>>>> Test to get 100 users <<<<<
I0801 17:16:57.584743   77695 http_client.go:22] Retrieved 100 records from server
I0801 17:16:57.584761   77695 http_performance_analyzer.go:31] Time taken to read 100 records via HTTP<>Kafka: 10.153690406s
```
