# URL shortener
## Руководство по запуску:
1. docker-compose
  ```
  docker compose up --detach
  ```
2. go
  ```
   go run cmd/app/main.go -d
  ```
  -d - необязательный флаг, заставляет микросервис работать через Postgres, без него все сохраняется в map
## Handlers
1. POST http://localhost:8080/post
   body
   ```
   "data": "url.com"
   ```
   response:
   ```
   "ffo34"
   ```
   
3. GET http://localhost:8080/ffo34
   ```
   "url.com"
   ```
   GET http://localhost:8080/0000
   ```
   204 NO CONTENT
   ```
