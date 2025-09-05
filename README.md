# Sample-Task-4
# PROJECT STRUCTURE 


appversal
├── test_client/

│   └── test.go  # Your test client code here

├── cron/

│   └── jobs.go

├── health-check/

│   └── health.go

├── proto/

│   ├── report.proto

│   ├── report.pb.go

│   └── report_grpc.pb.go

├── server/

│   └── server.go

├── main.go

├── go.mod

├── go.sum


## Prerequisites

- Go 1.20+  
- Protobuf compiler (`protoc`)
- Go plugins for protobuf:
  - `google.golang.org/protobuf/cmd/protoc-gen-go`
  - `google.golang.org/grpc/cmd/protoc-gen-go-grpc`
- [robfig/cron/v3](https://github.com/robfig/cron)

## Setup

1. **Install dependencies**

go mod tidy 


2. **Build protobuf files**
cd proto
protoc --go_out=. --go-grpc_out=. report.proto
cd ..

text

3. **Run the service**
go run main.go

Console output will log generated reports, and cron job activity.

4. **Health Check**

cd health-check in different terminal 
go run health.go

Console output will generate Health status 

5. **test_client** 
cd test_client
go run test.go 

Console output will show the end points results for the user.

## Automated Cron Job

- The cron job in `cron/jobs.go` triggers every 10 seconds and calls `GenerateReport` for several predefined user IDs.
- Check the console logs to see periodic report generation output.

## Project Extensibility

- Reports are stored in memory by default (see `server/server.go`).
- Modular design: proto, server logic, cron jobs, and health check are all easily extensible in their own directories.
