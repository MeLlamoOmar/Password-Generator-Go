# Password Generator and Manager

A simple command-line password generator and manager built with Go that allows users to create, store, and manage passwords securely using SQLite database.

## Features

- Generate random passwords with customizable options:
  - Adjustable length
  - Include/exclude numbers
  - Include/exclude special symbols
- Store passwords with custom labels
- View all stored passwords
- View specific passwords by ID
- Persistent storage using SQLite

## Requirements

- Go 1.x
- SQLite3

## Installation

1. Clone the repository:

```bash
git clone https://github.com/MeLlamoOmar/Password-Generator-Go.git
cd goWeb
```

2. Install dependencies:

```bash
go mod tidy
```

## Usage

1. Run the application:

```bash
go run main.go
```

2. Use the interactive menu to:
   - Generate new passwords
   - Save existing passwords
   - View saved passwords
   - View specific passwords by ID

## Project Structure

```
goWeb/
├── main.go         # Main application entry point
├── util/
│   └── util.go     # Utility functions for password generation
├── model/          # Data models
├── service/        # Business logic layer
├── store/          # Data access layer
└── password.db     # SQLite database file
```

## Database Schema

The application uses a simple SQLite database with the following schema:

```sql
CREATE TABLE passwords(
    id integer primary key autoincrement,
    label text not null,
    password text not null,
    created_at text not null
)
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
