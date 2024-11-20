# Assets Mapping

## How to Run the App

1. **Prerequisites:**
   - Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).
   - Ensure that you have MongoDB installed and running. You can find installation instructions on the [MongoDB website](https://www.mongodb.com/try/download/community). You can use local instance or create an account in MongoDB Atlas, a cloud based MongoDB service.

2. **Clone and run the application:**
   ```bash
   git clone https://github.com/RishanKP/assets-mapping.git
   cd assets-mapping
   ```
   - Setup your .env file
    ```
    (an env file is provided with test configurations in the repo)

    DB_USER=dbUser
    DB_PASS=dbPass
    DB_CLUSTER=dbCluster
    DB_NAME=dbName
    PORT=8080
    JWT_TOKEN=myToken
    ```
   - Install dependencies and run the application
   ```bash
   go mod tidy
   go run main.go
    ```
## API DOCUMENTATION

(all APIs, except the login API requires authentication)

| Method | Endpoint          | Description                                                     |  
|--------|-------------------|-----------------------------------------------------------------|
| POST   | /login            | Login with credentials.                                         | 
| GET    | /dashboard        | Dashboard api to get mapping count based on employee            | 
| POST   | /employee         | Create a new employee.                                          | 
| PUT    | /employee         | Update employee details.                                        | 
| GET    | /employee         | Get list of employees.                                          |  
| GET    | /employee/{id}    | Get employee by id                                              | 
| DELETE | /employee/{id}    | Delete employee by id                                           | 
| POST   | /asset            | Create a new asset.                                             | 
| PUT    | /asset            | Update asset details.                                           | 
| GET    | /asset            | Get list of assets.                                             |  
| GET    | /asset/{id}       | Get asset by id                                                 | 
| DELETE | /asset/{id}       | Delete asset by id                                              | 
| POST   | /mapping          | Create a new asset mapping.                                             | 
| GET    | /mapping/employee/{id}       | Get mappings by employee id                                                 | 
| DELETE | /asset/{id}       | Delete mapping by id                                              | 

## SAMPLE REQUESTS
You can download the postman collection for sample requests using the link below:

[Download postman collection](./requests.json)

