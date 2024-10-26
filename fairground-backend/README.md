# Backend

Make sure you have Docker installed and running.

```bash
docker compose up
```
This will start the go backend + the postgres database and also initialize the database with the schema.

# API Documentation

### 1. Register
- **Method:** `POST`
- **URL:** `http://localhost:3000/api/v1/auth/register`
- **Request Body:**
    ```json
    {
        "first_name": "Student",
        "last_name": "Student",
        "email": "student@gmail.com",
        "password": "123456",
        "role": "student"
    }
    ```
- **Description:** This API registers a new user with the provided details.

### 2. Login
- **Method:** `POST`
- **URL:** `http://localhost:3000/api/v1/auth/login`
- **Request Body:**
    ```json
    {
        "email": "parent@gmail.com",
        "password": "123456"
    }
    ```
- **Description:** Authenticates a user and returns a token for subsequent requests.

### 3. List All Users
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/users`
- **Auth:** Bearer Token
- **Description:** Retrieves a list of all registered users.

### 4. List User
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/users/4`
- **Auth:** Bearer Token
- **Description:** Fetches details of a specific user by ID.

### 5. Get Me
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/user/me`
- **Auth:** Bearer Token
- **Description:** Retrieves the details of the authenticated user.

### 6. List All Tokens
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/tokens`
- **Auth:** Bearer Token
- **Description:** Lists all tokens available in the system.

### 7. List My Tokens
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/tokens/me`
- **Auth:** Bearer Token
- **Description:** Retrieves tokens associated with the authenticated user.

### 8. Purchase Token
- **Method:** `POST`
- **URL:** `http://localhost:3000/api/v1/tokens/purchase`
- **Request Body:**
    ```json
    {
        "amount": 500
    }
    ```
- **Auth:** Bearer Token
- **Description:** Allows a user to purchase tokens.

### 9. List All Students
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/students`
- **Auth:** Bearer Token
- **Description:** Retrieves a list of all students.

### 10. List All Parents
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/parents`
- **Auth:** Bearer Token
- **Description:** Fetches a list of all parents.

### 11. List My Students
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/parents/students`
- **Auth:** Bearer Token
- **Description:** Lists students associated with the authenticated parent.

### 12. Delete My Student
- **Method:** `DELETE`
- **URL:** `http://localhost:3000/api/v1/parents/student/17`
- **Auth:** Bearer Token
- **Description:** Deletes a student associated with the authenticated parent.

### 13. Add a Student to Parent
- **Method:** `POST`
- **URL:** `http://localhost:3000/api/v1/parents/student`
- **Request Body:**
    ```json
    {
        "email": "9876@gmail.com"
    }
    ```
- **Auth:** Bearer Token
- **Description:** Adds a student to the authenticated parent's list.

### 14. Get My Student
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/parents/student/{studentId}`
- **Auth:** Bearer Token
- **Description:** Retrieves details of a specific student associated with the authenticated parent.

### 15. Transfer Token
- **Method:** `POST`
- **URL:** `http://localhost:3000/api/v1/tokens/transfer`
- **Request Body:**
    ```json
    {
        "to_user_id": 31,
        "tokens": 1
    }
    ```
- **Auth:** Bearer Token
- **Description:** Transfers tokens from the authenticated user to another user.

### 16. Get My Transactions
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/transactions/me`
- **Auth:** Bearer Token
- **Description:** Retrieves transaction history for the authenticated user.

### 17. Get Students Parents
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/students/parents`
- **Auth:** Bearer Token
- **Description:** Lists the parents of the authenticated user’s students.

### 18. Add Parent to Student
- **Method:** `POST`
- **URL:** `http://localhost:3000/api/v1/students/parent`
- **Request Body:**
    ```json
    {
        "email": "parent1@gmail.com"
    }
    ```
- **Auth:** Bearer Token
- **Description:** Adds a parent to a student.

### 19. Delete Parent From Student
- **Method:** `DELETE`
- **URL:** `http://localhost:3000/api/v1/students/parent/3`
- **Auth:** Bearer Token
- **Description:** Removes a parent from a student's profile.

### 20. Get All Raffles
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/raffles`
- **Auth:** Bearer Token
- **Description:** Retrieves a list of all available raffles.

### 21. List My Raffle Tickets
- **Method:** `GET`
- **URL:** `http://localhost:3000/api/v1/raffle/tickets/me`
- **Auth:** Bearer Token
- **Description:** Lists raffle tickets owned by the authenticated user.

### 22. Purchase Raffle Ticket
- **Method:** POST
- **URL:** `http://localhost:3000/api/v1/raffle/ticket/purchase/{raffleId}`
- **Auth:** Bearer Token
- **Description:** Allows a user to purchase raffle tickets.

## Notes
- Ensure that you replace `{studentId}` and other placeholders with actual values when making requests.
- Use the provided Bearer Tokens in the Authorization header for secure endpoints.

