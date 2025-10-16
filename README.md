# Blog API

A RESTful API for a blog platform built with Go and Gin framework. This API provides user authentication, post management, and commenting functionality with JWT-based authentication and PostgreSQL database.

## 🚀 Features

- **User Authentication**: JWT-based user registration and login system
- **Post Management**: Create, read, and manage blog posts
- **Comment System**: Add and retrieve comments on posts
- **Protected Routes**: Secure endpoints with JWT middleware
- **PostgreSQL Database**: Robust data storage with proper relationships
- **RESTful API**: Clean and intuitive API endpoints

## 🛠 Tech Stack

- **Backend**: Go 1.25.3
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Environment Management**: godotenv

## 📋 Prerequisites

Before running this project, make sure you have the following installed:

- Go 1.25.3 or later
- PostgreSQL database
- Git

## ⚙️ Installation & Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/Fonate-Michael/Blog-API

   cd Blog-API
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   - Create a new PostgreSQL database named `blog-api`
   - Update the `.env` file with your database credentials

4. **Configure environment variables**
   ```env
   DB_NAME=blog-api
   DB_USER=your_db_user
   DB_PASS=your_db_password
   DB_SSL=disable
   ```

5. **Run database migrations**
   ```sql
   -- Users table
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(255) NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL
   );

   -- Posts table
   CREATE TABLE posts (
       id SERIAL PRIMARY KEY,
       user_id INTEGER REFERENCES users(id),
       title VARCHAR(255) NOT NULL,
       description TEXT NOT NULL
   );

   -- Comments table
   CREATE TABLE comments (
       id SERIAL PRIMARY KEY,
       user_id INTEGER REFERENCES users(id),
       post_id INTEGER REFERENCES posts(id),
       comment TEXT NOT NULL
   );

   -- Likes table
   CREATE TABLE likes (
       id SERIAL PRIMARY KEY,
       user_id INTEGER REFERENCES users(id),
       post_id INTEGER REFERENCES posts(id)
   );
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

The API will start running on `http://localhost:8000`

## 📚 API Endpoints

### Authentication Endpoints

#### Register User
```http
POST /register
```

**Request Body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "message": "User registered successfully"
}
```

#### Login User
```http
POST /login
```

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Public Endpoints

#### Get All Posts
```http
GET /posts
```

**Response:**
```json
{
  "posts": [
    {
      "id": 1,
      "user_id": 1,
      "title": "My First Blog Post",
      "description": "This is the content of my first blog post."
    }
  ]
}
```

#### Get Comments for a Post
```http
GET /posts/:id/comment
```

**Response:**
```json
{
  "comments": [
    {
      "id": 1,
      "user_id": 2,
      "post_id": 1,
      "comment": "Great post! Thanks for sharing."
    }
  ]
}
```

#### Health Check
```http
GET /health
```

**Response:**
```json
{
  "message": "OK"
}
```

### Protected Endpoints (Requires JWT)

Protected endpoints require an Authorization header with a Bearer token:

```
Authorization: Bearer <your_jwt_token>
```

#### Create New Post
```http
POST /posts
```

**Request Body:**
```json
{
  "title": "My New Post",
  "description": "Content of the new post"
}
```

**Response:**
```json
{
  "message": "Post added successfully"
}
```

#### Add Comment to Post
```http
POST /posts/:id/comment
```

**Request Body:**
```json
{
  "comment": "This is my comment on the post"
}
```

**Response:**
```json
{
  "message": "Comment added successfully"
}
```

## 🗄️ Database Schema

### Users Table
| Column   | Type    | Description          |
|----------|---------|---------------------|
| id       | SERIAL  | Primary key         |
| username | VARCHAR | User's username     |
| email    | VARCHAR | User's email        |
| password | VARCHAR | Hashed password     |

### Posts Table
| Column      | Type    | Description          |
|-------------|---------|---------------------|
| id          | SERIAL  | Primary key         |
| user_id     | INTEGER | Foreign key to users|
| title       | VARCHAR | Post title          |
| description | TEXT    | Post content        |

### Comments Table
| Column   | Type    | Description          |
|----------|---------|---------------------|
| id       | SERIAL  | Primary key         |
| user_id  | INTEGER | Foreign key to users|
| post_id  | INTEGER | Foreign key to posts|
| comment  | TEXT    | Comment content     |

### Likes Table
| Column   | Type    | Description          |
|----------|---------|---------------------|
| id       | SERIAL  | Primary key         |
| user_id  | INTEGER | Foreign key to users|
| post_id  | INTEGER | Foreign key to posts|

## 🔐 Authentication

This API uses JWT (JSON Web Tokens) for authentication:

1. **Registration**: Create a new user account via `/register`
2. **Login**: Authenticate and receive a JWT token via `/login`
3. **Authorization**: Include the token in the `Authorization` header as `Bearer <token>` for protected routes
4. **Token Expiration**: Tokens expire after 1 hour

The JWT secret key is currently hardcoded as `secret_key` in the middleware. In production, use a secure environment variable.

## 🏗️ Project Structure

```
blog-api/
├── controller/
│   ├── auth_controller.go     # User registration and login
│   ├── blog_contoller.go      # Post management
│   └── comments_controller.go # Comment management
├── db/
│   └── db.go                  # Database connection
├── middleware/
│   └── auth_middleware.go     # JWT authentication middleware
├── models/
│   └── models.go              # Data structures
├── routes/
│   └── routes.go              # API route definitions
├── .env                       # Environment variables
├── go.mod                     # Go module dependencies
├── go.sum                     # Dependency checksums
└── main.go                    # Application entry point
```

## 🔧 Usage Examples

### Using cURL

#### Register a new user:
```bash
curl -X POST http://localhost:8000/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "email": "test@example.com", "password": "password123"}'
```

#### Login:
```bash
curl -X POST http://localhost:8000/login \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "password": "password123"}'
```

#### Create a post:
```bash
curl -X POST http://localhost:8000/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title": "Hello World", "description": "My first post"}'
```

### Using JavaScript/Fetch

```javascript
// Login
const login = async () => {
  const response = await fetch('http://localhost:8000/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: 'test@example.com',
      password: 'password123'
    })
  });

  const data = await response.json();
  localStorage.setItem('token', data.token);
};

// Create post
const createPost = async (title, description) => {
  const token = localStorage.getItem('token');
  const response = await fetch('http://localhost:8000/posts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    },
    body: JSON.stringify({ title, description })
  });

  return response.json();
};
```

## 🚧 Current Limitations

- No password validation on registration
- No email verification system
- No rate limiting
- No input sanitization
- No error handling for duplicate emails/usernames
- JWT secret is hardcoded (should be in environment variable)
- No pagination for posts/comments
- No user profile management
- No post editing/deletion functionality

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Fonate Michael**
- GitHub: [@hehehe](https://github.com/Fonate-Michael)

---

Made with ❤️ using Go and Gin
