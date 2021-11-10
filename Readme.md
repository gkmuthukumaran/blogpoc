###source

1. git clone https://github.com/gkmuthukumaran/blogpoc
2. cd blogpoc
3. go run main.go
4. copy the server url
5. install postman ping the url mention below for example http://localhost:8080
6. Get JWT toke for access all api without token you can't call the api, its enabled for security reasons
      =>URL :localhost:8080/login 
      => Method : post
      =>form-data => username : bloguser password: p@ss1234
      => copy the token url and use for call the api get blogs and insert blogs

HOW TO USE TOKEN FOR CALL THE APIS Examble below in pass as header

        Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY1ODIzODMsIm5hbWUiOiJibG9ndXNlciJ9.1hIbuWQdw8OOpR80VhIG6XbBzdvsGEAgXPbs7MFcLFY
        Content-Type:application/json


### Docker
1. git clone https://github.com/gkmuthukumaran/blogpoc
2. cd blogpoc
3. docker build --tag blogpoc .
4. docker run -d -p 80:8080 blogpoc .

### Unit test
go run main.go
go test ./...
