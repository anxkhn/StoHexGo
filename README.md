# StoHexGo: Learning Go and Hexagonal Architecture through Stock Portfolio Management



## Project Overview

StoHexGo is a hands-on learning project designed to explore the fundamentals of Go programming and the Hexagonal Architecture pattern. It tackles the practical problem of stock portfolio management, providing an engaging context to apply these concepts.

## Problem Statement

The core challenge is to build a system that manages a portfolio of stocks, enabling buying and selling operations. The system must adhere to the First In, First Out (FIFO) principle when selling shares, ensuring accurate profit calculations.

## Constraints

1.  **FIFO Selling Rule:** When selling shares, prioritize the oldest (first bought) shares.
2.  **Stock Ownership**: Prevent selling more shares than currently owned.
3.  **Transaction Order**: Process transactions chronologically, disallowing sales before purchases.
4.  **Transaction Details**: Each transaction includes its type ("buy" or "sell"), stock identifier, price, and quantity.

## Solution Highlights

*   **Go Language:** Leverages Go's simplicity, concurrency, and performance for efficient backend development.
*   **Hexagonal Architecture:** Organizes the codebase into loosely coupled components (ports and adapters), promoting adaptability and testability.
*   **GORM (SQLite):** Employs GORM for seamless database interactions, simplifying data persistence.
*   **RESTful API:** Provides endpoints to process transactions and retrieve portfolio balance.

## Learning Objectives

*   **Go Fundamentals:** Gain hands-on experience with Go syntax, data structures, error handling, and concurrency.
*   **Hexagonal Architecture:** Understand the principles and benefits of this architectural pattern, fostering clean code and maintainability.
*   **Practical Application:** Apply theoretical knowledge to a real-world problem, solidifying understanding.

## Getting Started

### Prerequisites

-   **Go 2.22+**
-   **SQLite Database** (managed via GORM)

### Installation and Setup

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/anxkhn/StoHexGo.git)
    cd StoHexGo
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Set up the database:**
    *   GORM will automatically handle database migrations when the application starts.

4.  **Run the application:**

    ```bash
    go run main.go
    ```

5.  **Test the API:**
    *   Use tools like Postman, cURL, or any REST client to test the API endpoints.

---

## API Endpoints

### 1. `/` (GET)

Checks if the server is running.

**Response:**

```json
{
    "msg": "🚀 Server is running!"
}
```


### 2. `/transaction` (POST)

Processes a list of transactions, calculates the net profit, and returns it.

**Request Body:**

```json
{
    "transactions": [
        { "type": "buy", "stockId": "AAPL", "price": 100, "quantity": 10 },
        { "type": "buy", "stockId": "AAPL", "price": 150, "quantity": 15 },
        { "type": "sell", "stockId": "AAPL", "price": 120, "quantity": 10 },
        { "type": "buy", "stockId": "AAPL", "price": 110, "quantity": 10 },
        { "type": "sell", "stockId": "AAPL", "price": 130, "quantity": 25 },
        { "type": "buy", "stockId": "AAPL", "price": 140, "quantity": 20 },
        { "type": "sell", "stockId": "AAPL", "price": 150, "quantity": 15 },
        { "type": "buy", "stockId": "AAPL", "price": 135, "quantity": 10 },
        { "type": "sell", "stockId": "AAPL", "price": 160, "quantity": 15 }
    ]
}
```

**Response:**

```json
{
    "profit": 600
}
```

### 3. `/balance` (GET)

Retrieves the current inventory of stocks in the portfolio.

**Response:**

```json
{
    "AAPL": 20
}
```

## Project Structure

Here's an overview of the project structure:

```
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── handlers
│   │   │   └── http_handler.go
│   │   └── repositories
│   │       └── sqlite_repository.go
│   ├── config
│   │   └── config.go
│   └── core
│       ├── domain
│       │   └── transactions.go
│       ├── ports
│       │   ├── repositories.go
│       │   └── services.go
│       └── services
│           └── portfolio_services.go
├── portfolio.db
└── README.md
```

### Explanation of Hexagonal Architecture in Go

**Hexagonal Architecture**, also known as the Ports and Adapters pattern, aims to create a flexible and maintainable system by separating the core business logic from external systems (like databases, web servers, etc.). Here’s how it is implemented in StoHexGo:

- **`cmd/main.go`**: The entry point of the application. It initializes the server and other necessary components.

- **`internal/adapters`**: Contains the adapters for communication with the external world.
  - **`handlers`**: Implements HTTP handlers that process incoming requests and send responses. For example, `http_handler.go` manages REST API interactions.
  - **`repositories`**: Handles data persistence and retrieval. `sqlite_repository.go` is responsible for interacting with the SQLite database using GORM.

- **`internal/config`**: Manages configuration settings of the application. `config.go` provides configuration loading and management functionality.

- **`internal/core`**: Contains the core business logic and the application's domain model.
  - **`domain`**: Defines the core domain models and entities. `transactions.go` includes definitions related to stock transactions.
  - **`ports`**: Defines interfaces for interacting with the core logic. `repositories.go` and `services.go` specify the contracts for repository and service layers.
  - **`services`**: Implements the core business logic. `portfolio_services.go` includes the logic for managing stock portfolios and calculating profits.

- **`portfolio.db`**: The SQLite database file used for data persistence.


---

## Conclusion

This project demonstrates how to implement a stock portfolio management system that adheres to FIFO rules using Go, GORM, and Hexagonal Architecture. The API allows you to manage transactions and calculate profits while maintaining a clean and scalable codebase.

