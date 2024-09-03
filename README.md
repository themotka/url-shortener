# URL shortener
## Руководство по запуску:
1. docker
  ```
   docker build -t prod:local .
  ```
2. docker-compose
  ```
   docker compose up 
  ```
 Флаг -d, отвечающий за запись в postgres, устанавливается в Dockerfile
 ```
 CMD ["/app", "-d"]
 ```
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
