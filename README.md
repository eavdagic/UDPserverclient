# UDPserverclient
UDP client/server

## Set Up
#### Client
cd to client directory
run ```go build```
#### Server
cd to server directory
run ```go build```

## To Run It
After building client and server:
#### Using start.sh
Run the ./start.sh script to run both the client and server at same time without needing to open a second terminal
start.sh accepts two flags, -p and -a
  * -p is port (ex: 8000)
  * -a is host:port (ex: 127.0.0.1:8001)
#### Running client/server individually
##### Server
cd to server directory and run server
to run with default flags ```./server```
./server accepts -address flag (ex: 127.0.0.1:8000)
to run with address flag ```./server -address="127.0.0.1:8000"```
##### Client
cd to client directory and run client
to run with default flags ```./client```
./client accepts -address and -port flag
to run with flags ```./client -address="127.0.0.1:8000" -port="8001"

