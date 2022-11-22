# ChessCMS
 
docker run --name mongodb -d -p 27017:27017 mongo
protoc -I=./protobuf --go_out=./protobuf ./protobuf/chess.proto