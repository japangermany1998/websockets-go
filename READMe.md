# Mục tiêu
- Khởi chạy 1 server chat sử dụng websocket golang
- Khởi chạy 4 client trong server chat, client1 và client2 trong room1, client3 và client4 trong room2

# Các bước chạy  
## 1. Chạy server
```go
go run main.go client.go server
```

## 2. Chạy các client:
```go
go run main.go client.go client1
go run main.go client.go client2
go run main.go client.go client3
go run main.go client.go client4
```

## 3. Kết quả
![image](https://github.com/japangermany1998/websockets-go/assets/81943808/a419ebe1-a277-4a8f-81a9-62efcdf55bcd)

