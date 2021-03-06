```
docker-compose up -d --build
```
```
docker exec -it server /bin/bash
```
В докере
```
/micro login
```
Login: admin

Password: micro

Запускаем сервис albato
```
/micro run albato
```
Запускаем сервис faker
```
/micro run faker
```
Смотрим
```
http://localhost:8099/faker/FakeAddress
```
```
http://localhost:8099/albato/CustomUserRequest?msg=%22awd%22&msg2=%22123%22
```