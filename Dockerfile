#base image golang 1.23.5
FROM golang:1.23.5
#set current working directory inside container
WORKDIR /app
#copy go mod and sum files
COPY go.mod ./
#copy go.sum  
COPY go.sum ./
#download all dependencies 
RUN go mod download
#copy the source code
COPY . .
# build the binary
RUN go build -o main ./cmd/api
#expose port  to the outside
EXPOSE 8080
#make the executable file
RUN chamod +x main
#command to run 
CMD ["./main"]
