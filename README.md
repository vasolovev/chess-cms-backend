# ChessCMS
 
docker run --name mongodb -d -p 27017:27017 mongo
protoc --go_out=paths=import:. -I. ChessCMS/protobuf/chess.proto