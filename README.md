
<p align="center">
  <a href="" rel="noopener">
 <img src="https://djeqr6to3dedg.cloudfront.net/repo-logos/library/golang/live/logo.png" alt="Project logo"></a>
</p]>
<h3 align="center">GOLANG REST API</h3>


# REST_API Package.

This package contains the REST API implementation for the Golang project. We have dockerised application and pushed the image to Docker hub via github action and deployed to self hosted EC2 instance as well (github action ci/cd)

## Endpoints

- `GET /events`: Get all events
- `GET /events/:id`: Get a specific event
- `POST /events`: Create a new event
- `PUT /events/:id`: Update an event
- `DELETE /events/:id`: Delete an event
- `POST /events/:id/register`: Register for an event
- `DELETE /events/:id/register`: Cancel event registration
- `POST /signup`: User signup
- `POST /login`: User login

## Authentication

- Routes are protected using JWT authentication
- User signup and login functionality included
- Authorization implemented for few routes where only the valid loggedin user who created event only allowed to update and delete events

## Usage

1. Initialize the database by running `db.InitDB()` in the [main](cci:1:///Users/vikasmalviya/REST_API/main.go:10:0-19:1) package.
2. Register the routes using `routes.RegisteredRoutes(server)` in the [main](cci:1:///Users/vikasmalviya/REST_API/main.go:10:0-19:1) package.

## Running the Project

Locally Run
1. Clone the repository:
   ```bash
   git clone git@github.com:vikky225/GolangRestAPI.git
   ```

2.Navigate to the project directory
```bash
cd REST_API
```

3.Install dependencies
```go mod download```

4. Run the application
```go run .```

## Dockerized Environment to Run
To run your application in a Dockerized environment based on the provided Dockerfile, follow these instructions:

- Ensure Docker is installed on your system.

- Create a directory containing your Go application code and the Dockerfile.

- Save the provided Dockerfile in the same directory as your application code.

- Open a terminal and navigate to the directory containing your code and Dockerfile.

- Build the Docker image using the following command:

``` docker build -t my-go-app .```
This command will build the Docker image based on the instructions in the Dockerfile. Replace my-go-app with a suitable name for your Docker image.

- Run a Docker container based on the image you just built:

``` docker run -p 8080:8080 my-go-app```
- This command will start a container running your Go application, which will be accessible at http://localhost:8080.

These instructions assume you have your Go application code ready in the same directory as the Dockerfile. Let me know if you need further clarification or assistance!

## Git hub action for building and pushing image to docker hub than deploying it to EC2 instance
{{EC2 instacnce IP address}}:8080/events should work and give you results
or you can directly run as we have pushed image to docker hub so you have to run in your local 
docker run -p 8080:8080 vikky225/go-app:latest and than run localhost:8080/events and you can play around with endpoints to create and get events etc


you can access site as below 
http://ec2-3-26-144-138.ap-southeast-2.compute.amazonaws.com:8080/events 

## Dependencies

- `github.com/gin-gonic/gin`: Gin web framework
- `github.com/mattn/go-sqlite3`: SQLite database driver
- `github.com/golang-jwt/jwt/v5`: JWT library

## Author

- Vikas Malviya

Feel free to reach out for any questions or contributions related to the REST_API package!
