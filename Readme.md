# Wallet API

This is a simple RESTful API built with Go for managing user wallets and transactions. It allows users to create accounts, create wallets, check their balance, transfer funds between wallets, and view their transaction history.

## Functionality

- **User Management:**
  - `POST /users`: Creates a new user. Requires `name` and `email` in the request body.
- **Wallet Management:**
  - `POST /wallets`: Creates a new wallet for an existing user. Requires `user_id` and `initial_balance` in the request body.
  - `GET /wallets/{wallet_id}/balance`: Retrieves the balance of a specific wallet.
- **Transaction Management:**
  - `POST /wallets/transfer`: Transfers funds from one wallet to another. Requires `from_wallet_id`, `to_wallet_id`, and `amount` in the request body.
  - `GET /wallets/{wallet_id}/transactions`: Retrieves the transaction history for a specific wallet.

## How the Project Works

1.  **Request Handling:** The `main.go` file sets up an HTTP server using the `gorilla/mux` router. The `routes/routes.go` file defines the API endpoints and maps them to specific handler functions in the `controllers` directory.
2.  **Controller Logic:** The controller functions (`controllers/user.go` and `controllers/transaction.go`) handle the business logic for each API endpoint. They:
    - Decode the JSON request body.
    - Perform validation of input data.
    - Interact with the database using GORM.
    - Construct JSON responses with appropriate HTTP status codes.
3.  **Database Interaction:** The `database/postgres.go` file establishes a connection to a PostgreSQL database using the credentials defined in the `.env` file. It uses GORM to perform database operations (create, read, update) on the `users`, `wallets`, and `transactions` tables, which are automatically created or migrated based on the models defined in the `models` directory.
4.  **Data Models:** The `models` directory contains Go structs (`User`, `Wallet`, `Transaction`) that represent the database tables and are used for data serialization and deserialization. GORM uses these models to interact with the database.
5.  **Utilities:** The `utils/response.go` file provides a helper function (`RespondJSON`) for sending consistent JSON responses.
6.  **Configuration:** The `config/config.go` file (though not explicitly shown in the provided code) typically handles loading environment variables from a `.env` file using the `github.com/joho/godotenv` library.

## How to Run the Project

1.  **Prerequisites:**

    - **Go Installation:** Ensure you have Go (version 1.16 or later) installed on your system. You can download it from [https://go.dev/dl/](https://go.dev/dl/).
    - **PostgreSQL Installation:** You need a running PostgreSQL database. You can download and install it from [https://www.postgresql.org/download/](https://www.postgresql.org/download/).
    - **PostgreSQL User and Database:** Create a PostgreSQL user (`wallet_user`) and a database (`wallet_db`) with the password (`wallet_password`) as specified in the `.env` file. You can use a tool like `psql` or pgAdmin for this.
    - **Postman or Similar Tool:** A REST client like Postman ([https://www.postman.com/downloads/](https://www.postman.com/downloads/)) is recommended for testing the API endpoints.

2.  **Clone the Repository:**

    ```bash
    git clone <repository_url>
    cd wallet-api
    ```

    (Replace `<repository_url>` with the actual URL of your Git repository)

3.  **Set Up Environment Variables:**

    - Create a `.env` file in the project root directory with the following content:
      ```dotenv
      DB_HOST=localhost
      DB_PORT=5432
      DB_USER=wallet_user
      DB_PASSWORD=wallet_password
      DB_NAME=wallet_db
      ```
      **Note:** The default PostgreSQL port is `5432`. Ensure your PostgreSQL server is running on this port or update the `DB_PORT` in the `.env` file accordingly.

4.  **Install Dependencies:**

    ```bash
    go mod tidy
    ```

    This command will download all the necessary Go dependencies, including `gorm.io/gorm`, `gorm.io/driver/postgres`, `github.com/joho/godotenv`, and `github.com/gorilla/mux`.

5.  **Run the Application:**
    ```bash
    go run main.go
    ```
    You should see the message `Server running on http://localhost:8080` in your terminal, indicating that the API server has started successfully.

## Dependencies Used

- **`github.com/joho/godotenv`:** For loading environment variables from a `.env` file.
- **`gorm.io/gorm`:** An ORM library for Go, simplifying database interactions.
- **`gorm.io/driver/postgres`:** The PostgreSQL driver for GORM.
- **`github.com/gorilla/mux`:** A powerful HTTP router for Go, used for handling API endpoints and path parameters.

## Testing the API with Postman

You can use Postman to send requests to the API endpoints. Ensure your Go server is running.

**Example Requests:**

- **Create User (POST `http://localhost:8080/users`):**

  ```json
  {
    "name": "Devansh",
    "email": "Devansh.apply@gmail.com"
  }
  ```

- **Create Wallet (POST `http://localhost:8080/wallets`):**

  ```json
  {
    "user_id": 1,
    "initial_balance": 10000.0
  }
  ```

- **Get Wallet Balance (GET `http://localhost:8080/wallets/{wallet_id}/balance`):**
  Replace `{wallet_id}` with the ID of the wallet you want to check (e.g., `http://localhost:8080/wallets/1/balance`).

- **Transfer Funds (POST `http://localhost:8080/wallets/transfer`):**

  ```json
  {
    "from_wallet_id": 1, //sender's wallet ID
    "to_wallet_id": 2, // receiver's wallet ID
    "amount": 25.0
  }
  ```

- **Get Wallet Transactions (GET `http://localhost:8080/wallets/{wallet_id}/transactions`):**
  Replace `{wallet_id}` with the ID of the wallet whose transactions you want to view (e.g., `http://localhost:8080/wallets/1/transactions`).

## THANKS FOR THE OPPORTUNITY
