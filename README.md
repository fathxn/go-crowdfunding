
# Go Crowdfunding REST API Documentation
## Introduction
**go-crowdfunding** is a backend project for crowdfunding web application where users as a backer can fund or donate to support the user's project. The backend of this project is build with Golang and Gin.
### Base URL
All API endpoints can be accessed via the following base URL:
`https://127.0.0.1:8080/v1`
## Endpoints
### 1. Auth Endpoints
#### 1.1 Register Endpoint
`POST` `/users`
```http
POST /users HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json
{ 
	"name": "John Doe",
	"occupation": "Backend Developer", 
	"email": "john@doe.com",
	"password": "password123"
}
```
#### 1.2 Login/Session Endpoint
`POST` `/sessions`
```http
POST /sessions HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json
{
	"email": "john@doe.com",
	"password": "password123"
}
```
#### 1.3 Email Check Endpoint
`POST` `/email_check`
```http
POST /email_check HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json
{
	"email": "john@doe.com"
}
```
#### 1.4 Upload Avatar Endpoint
`POST` `/avatars`
```http
POST /avatars HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
Content-Type: multipart/form-data

--boundary--
Content-Disposition: form-data; name="file"; filename="profile.jpg"
Content-Type: image/jpeg
(binary image data)
--boundary--
```
#### 1.5 Fetch User Endpoint
`GET` `/users/fetch`
```http
GET /users/fetch HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
```

### 2. Campaign Endpoints
#### 2.1 Get All Campaigns Endpoint
`GET` `/campaigns`
```http
GET /campaigns HTTP/1.1
Host: 127.0.0.1:8080
```
#### 2.2 Get Campaign Detail Endpoint
`GET` `/campaigns/:id`
```http
GET /campaigns/1 HTTP/1.1
Host: 127.0.0.1:8080
```
#### 2.3 Create Campaign Endpoint
`POST` `/campaigns`
```http
POST /campaigns HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
Content-Type: application/json
{
	"name": "Campaign Title",
	"short_description": "Campaign Short Description",
	"description": "Campaign Very Long Description",
	"goal_amount": "1000000",
	"perks": "some,perks,separate,with,comma",
}
```
#### 2.4 Edit Campaign Endpoint
`PUT` `/campaigns/:id`
```http
PUT /campaigns/1 HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
Content-Type: application/json
{
	"name": "Campaign Title",
	"short_description": "Campaign Short Description",
	"description": "Campaign Very Long Description",
	"goal_amount": "1000000",
	"perks": "some,perks,separate,with,comma",
}
```
#### 2.5 Upload Campaign Images Endpoint
`PUT` `/campaigns/:id`
```http
PUT /campaigns/1 HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
Content-Type: multipart/form-data

--boundary--
Content-Disposition: form-data; name="file"; filename="profile.jpg"
Content-Type: image/jpeg
(binary image data)
--boundary--
```
### 3. Transaction Endpoints
#### 3.1 Get Transaction List From Campaign By ID
`GET` `/campaigns/:id/transactions`
```http
GET /campaigns/1/transactions HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
```
#### 3.2 Get Transaction List From User By ID
`GET` `/transactions`
```http
GET /transactions HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
```
#### 3.3 Create Transaction
`POST` `/transactions`
```http
POST /transactions HTTP/1.1
Host: 127.0.0.1:8080
Authorization: Bearer <your-token>
Content-Type: application/json
{
	"amount": 1000000,
	"campaign_id": 1
}
```

## Error Codes

| Code | Message                  | Description                             |
|------|--------------------------|-----------------------------------------|
| 400  | Bad Request              | The request is invalid.                 |
| 401  | Unauthorized             | Authentication failed.                  |
| 403  | Forbidden                | You do not have permission to access this resource. |
| 404  | Not Found                | Resource not found.                     |
| 415  | Unsupported Media Type   | The media type is not supported.        |
| 422  | Unprocessable Entity     | The request was well-formed but could not be followed due to semantic errors. |
| 500  | Internal Server Error    | An error occurred on the server.        |
| 503  | Service Unavailable      | The server is currently unavailable.    |

## Contact
If you have any questions or issues regarding this API, please contact me at fathanfirdaus20@gmail.com.