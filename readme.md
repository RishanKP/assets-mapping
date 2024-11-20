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
## DATABASE DESIGN

### 1. User Collection (`users`)

The `users` collection stores information about the employees who will be assigned various assets.

#### Fields:
- `id` (Primary Key): Unique identifier for each user (ObjectId).
- `email`: The user's email address (string).
- `password`: The user's password (hashed).
- `firstName`: The user's first name (string).
- `lastName`: The user's last name (string).
- `contact`: The user's contact number (string).
- `gender`: The user's gender (string).
- `bloodGroup`: The user's blood group (string).
- `emergencyContact`: The user's emergency contact number (string).
- `createdAt`: The date and time when the user was created (Date).
- `updatedAt`: The date and time when the user information was last updated (Date).

### 2. Assets Collection (`assets`)

The `assets` collection stores information about various assets within the organization, such as laptops, furniture, and equipment.

#### Fields:
- `id` (Primary Key): Unique identifier for each asset (ObjectId).
- `name`: The name of the asset (string).
- `type`: The type of asset (e.g., "Laptop", "Monitor", etc.) (string).
- `createdAt`: The date and time when the asset was created (Date).
- `updatedAt`: The date and time when the asset information was last updated (Date).

### 3. Mapping Collection (`assets`)
The `mappings` collection holds the data that links a **user** to an **asset**. This is a crucial part of the system as it defines the many-to-many relationship between employees (users) and the assets they are assigned.

### Fields:

- `id` (Primary Key): A unique identifier for each mapping record (ObjectId).
- `userId`: The ObjectId of the user (employee) to whom the asset is assigned (ObjectId, references `users._id`).
- `assetId`: The ObjectId of the asset that is assigned to the user (ObjectId, references `assets._id`).
- `createdAt`: The date and time when the mapping was created (Date).
- `updatedAt`: The date and time when the mapping was last updated (Date).




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

