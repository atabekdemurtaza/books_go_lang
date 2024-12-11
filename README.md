## Installation

# Prerequisites

* Go 1.20 or later installed. You can download it from Go's official website.
* SQLite (pre-installed on most systems).

# Steps

1. Clone the repository.
```
git clone https://github.com/atabekdemurtaza/books_go_lang.git
cd book-management-api
```

2. Install dependencies
```
go mod tidy
```

3. Run the app
```
go run main.go
```

4. Or use Make commands to run
```
    make help
```

# API Endpoints
### Base URL

#### Routes
| Method | Endpoint   | Description      |
|--------|------------|------------------|
| GET    | /books     | Get all books    |
| GET    | /books/:id | Get a book by id |
| POST   | /books     | Post a book      |
| DELETE | /books/:id | Delete a book    |

1. Get all books

```
curl http://localhost:8080/books
```
  
2. Create a new book

```
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"title":"Book name", "author":"Author"}'
```


