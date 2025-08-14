# E-commerce API

A RESTful API built with Go for managing products and users in an e-commerce application.

## 🚀 Features

- **Product Management**: CRUD operations for products
- **User Management**: User registration and retrieval
- **MySQL Database**: Persistent data storage
- **RESTful API**: Clean HTTP endpoints
- **CORS Support**: Cross-origin resource sharing enabled

## 📋 Prerequisites

- Go 1.24.3 or higher
- MySQL database
- Git

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <your-repository-url>
   cd GO
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=your_username
   DB_PASS=your_password
   DB_NAME=your_database_name
   ```

4. **Set up MySQL database**
   
   Create a MySQL database and update the `.env` file with your database credentials.

## 🏃‍♂️ Running the Application

1. **Start the server**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:3000`

2. **Verify the server is running**
   ```bash
   curl http://localhost:3000/products
   ```

## 📚 API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| GET | `/product/{productID}` | Get product by ID |
| POST | `/create-products` | Create a new product |

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | Get all users |
| GET | `/user/` | Get user by ID |
| POST | `/create-user` | Create a new user |

## 📝 API Examples

### Get All Products
```bash
curl http://localhost:3000/products
```

### Get Product by ID
```bash
curl http://localhost:3000/product/1
```

### Create Product
```bash
curl -X POST http://localhost:3000/create-products \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Product",
    "description": "Product description",
    "price": 29.99,
    "image_url": "http://example.com/image.jpg"
  }'
```

### Get All Users
```bash
curl http://localhost:3000/users
```

### Create User
```bash
curl -X POST http://localhost:3000/create-user \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

## 🏗️ Project Structure

```
GO/
├── cmd/
│   └── serve.go          # Server initialization
├── controllers/
│   └── user_controller.go # User controller logic
├── db/
│   └── db.go            # Database connection
├── global_router/
│   └── global_router.go # Global middleware (CORS)
├── models/
│   └── user.go          # User data model
├── routes/
│   └── routes.go        # Route definitions
├── util/
│   ├── debug.go         # Debug utilities
│   └── send_data.go     # Response utilities
├── main.go              # Application entry point
├── go.mod               # Go module file
└── go.sum               # Go module checksums
```

## 🔧 Configuration

The application uses environment variables for database configuration. Make sure to set up your `.env` file with the following variables:

- `DB_HOST`: MySQL host (default: localhost)
- `DB_PORT`: MySQL port (default: 3306)
- `DB_USER`: MySQL username
- `DB_PASS`: MySQL password
- `DB_NAME`: MySQL database name

## 🧪 Testing

To test the API endpoints, you can use tools like:
- cURL (command line)
- Postman
- Insomnia
- Any HTTP client

## 📦 Dependencies

- `github.com/go-sql-driver/mysql`: MySQL driver for Go
- `github.com/joho/godotenv`: Environment variable management
- `filippo.io/edwards25519`: Cryptographic library

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

If you encounter any issues or have questions, please open an issue on the GitHub repository.

---

**Note**: This is a development project. Make sure to implement proper security measures, input validation, and error handling before using in production. 