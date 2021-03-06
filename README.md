Todo Demo App

## Run 
Pull the code into your local Go workspace. $ git clone https://github.com/satishbhoi/TodoList.git

Open console in the project folder and run below command.

go build ./ToDoList


## Signup

Retrieve user credentials from the body and validate against database.
For invalid email or password, `send 400 - Bad Request` response.
For valid email and password, save user in database and send `201 - Created` response.  

Request  

```sh
curl \
  -X POST \
  http://localhost:9000/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"satish@gmail.com","password":"satish@123"}'
```
Response  

```json
 { "id": "58465b4ea6fe886d3215c6df", "email": "satish@gmail.com", "password": "satish@123" }
```


## Login
User login  

Retrieve user credentials from the body and validate against database. Each subsequent request must include the Authorization header.
Method: `POST`  
Path: `/login`  

Request  

```sh
curl \
  -X POST \
  http://localhost:9000/login \
  -H "Content-Type: application/json" \
-d '{"email":"satish@gmail.com","password":"satish@123"}'

```
Response  


```json
{ "id": "5c6af5b3a309fe6c5b677bb1",
 "email": "satish@gmail.com",
 "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTA3NzMxMTUsImlkIjoiNWM2YWY1YjNhMzA5ZmU2YzViNjc3YmIxIn0.eqki3kLJyWQ4qVbKHf7Q1GTfMjxqyzj7-DwO1KXl94g" }
```



## Add Task

Method: `POST` 
Path: `/addTodo`

Request

```sh
curl \
  -X POST \
  http://localhost:9000/addTodo \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"

json

{ "TaskName":"pLAY FOOTBALL", "TaskDesc":"ON A MATCH FOOTBALL", "Notes":"PICKING UP FOODS", "Status":0 }
```


Response
```

{ "id": "5c6afca9a309fe7aa57b30ea", "UserId": "5c6af5b3a309fe6c5b677bb1", "TaskName": "pLAY FOOTBALL", "TaskDesc": "ON A MATCH FOOTBALL", "Notes": "PICKING UP FOODS", "TaskDate": "2019-02-19T00:12:49.800254781+05:30", "Status": 0 }
```



## Get Todo List 
```
Method:`POST`
Path:`getTodo`

 sh
  curl \ 
-X POST \
http://localhost:9000/getTodo \
 -H "Authorization: Bearer" eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
```
```
Response: 
[ { "id": "5c6afca9a309fe7aa57b30ea",
 "UserId": "5c6af5b3a309fe6c5b677bb1", 
 "TaskName": "pLAY FOOTBALL",
  "TaskDesc": "ON A MATCH FOOTBALL",
   "Notes": "PICKING UP FOODS", 
   "TaskDate": "2019-02-19T00:12:49.8+05:30",
    "Status": 0 } ]
```



## Update Todo
Request:
```
 sh 
 curl \ 
-X POST \
http://localhost:6000/updateTodo/5c6afca9a309fe7aa57b30ea  \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"

Json

{ "TaskName":"PLAY CRICKET", "TaskDesc":"ON A CRICKET MATCH", }
```

Response:
```

{ "id": "5c6afca9a309fe7aa57b30ea", "UserId": "5c6af5b3a309fe6c5b677bb1", "TaskName": "PLAY CRICKET", "TaskDesc": "ON A CRICKET MATCH", "Notes": "PICKING UP FOODS", "TaskDate": "2019-02-19T00:12:49.8+05:30", "Status": 0 }
```

## Status Update
```
Request: 
sh 
curl \
-X POST \
http://localhost:6000/setStatus/5c6afca9a309fe7aa57b30ea \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
```

Response:

```
Task completed
```

