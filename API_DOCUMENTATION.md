# Local Friendly Marketplace API Documentation

## Overview
This is the REST API documentation for the Local Friendly Marketplace backend. All endpoints are prefixed with `/api`.

**Base URL:** `http://localhost:3000/api`

## Authentication
Most protected endpoints require JWT authentication via the `Authorization` header:
```
Authorization: Bearer <jwt_token>
```

JWT tokens are obtained from login or register endpoints and are valid for 24 hours.

## Response Format
All responses follow a consistent JSON structure:

### Success Response
```json
{
  "success": true,
  "message": "Operation description",
  "data": {}
}
```

### Error Response
```json
{
  "success": false,
  "error": "Error message"
}
```

---

## Authentication Endpoints

### Register User
**POST** `/auth/register`

Create a new user account.

#### Request
```json
{
  "name": "Alya Putri",
  "email": "alya@example.com",
  "password": "Secret123!",
  "phone": "+628123456789",
  "profileImageUrl": "https://example.com/profiles/alya.jpg",
  "roles": ["buyer"]
}
```

#### Response (201 Created)
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": "664a1f2c7d9a2c0012345678",
      "name": "Alya Putri",
      "email": "alya@example.com",
      "phone": "+628123456789",
      "profileImageUrl": "https://example.com/profiles/alya.jpg",
      "roles": ["buyer"],
      "marketplaceId": null,
      "sellerId": null,
      "createdAt": "2026-05-19T10:00:00Z",
      "updatedAt": "2026-05-19T10:00:00Z",
      "lastSyncedAt": null,
      "isSynced": false
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

### Login User
**POST** `/auth/login`

Authenticate user and obtain JWT token.

#### Request
```json
{
  "email": "alya@example.com",
  "password": "Secret123!"
}
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": "664a1f2c7d9a2c0012345678",
      "name": "Alya Putri",
      "email": "alya@example.com",
      "phone": "+628123456789",
      "profileImageUrl": "https://example.com/profiles/alya.jpg",
      "roles": ["buyer"],
      "marketplaceId": null,
      "sellerId": null,
      "createdAt": "2026-05-19T10:00:00Z",
      "updatedAt": "2026-05-19T10:00:00Z",
      "lastSyncedAt": null,
      "isSynced": false
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

## User Endpoints

### Get All Users
**GET** `/users` [Protected]

Retrieve all users. Admin only.

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": "664a1f2c7d9a2c0012345678",
      "name": "Alya Putri",
      "email": "alya@example.com",
      "phone": "+628123456789",
      "profileImageUrl": "https://example.com/profiles/alya.jpg",
      "roles": ["buyer"],
      "createdAt": "2026-05-19T10:00:00Z",
      "updatedAt": "2026-05-19T10:00:00Z"
    }
  ]
}
```

---

### Get User by ID
**GET** `/users/{id}` [Protected]

Retrieve specific user information.

#### Parameters
- `id` (path, required): User ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": "664a1f2c7d9a2c0012345678",
    "name": "Alya Putri",
    "email": "alya@example.com",
    "phone": "+628123456789",
    "profileImageUrl": "https://example.com/profiles/alya.jpg",
    "roles": ["buyer"],
    "createdAt": "2026-05-19T10:00:00Z",
    "updatedAt": "2026-05-19T10:00:00Z"
  }
}
```

---

### Update User
**PUT** `/users/{id}` [Protected]

Update user information.

#### Parameters
- `id` (path, required): User ID

#### Request
```json
{
  "name": "Alya Putri Updated",
  "email": "alya.updated@example.com",
  "phone": "+628123456789"
}
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": "664a1f2c7d9a2c0012345678",
    "name": "Alya Putri Updated",
    "email": "alya.updated@example.com",
    "phone": "+628123456789",
    "roles": ["buyer"],
    "updatedAt": "2026-05-19T11:00:00Z"
  }
}
```

---

### Delete User
**DELETE** `/users/{id}` [Protected]

Remove a user account.

#### Parameters
- `id` (path, required): User ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

---

## Product Endpoints

### Get All Products
**GET** `/products` [Public]

Retrieve all available products.

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": "664a1f2c7d9a2c0012345679",
      "sellerId": "664a1f2c7d9a2c0012345678",
      "name": "Kopi Premium Arabika",
      "description": "Biji kopi arabika pilihan dari Aceh",
      "price": 150000,
      "quantity": 50,
      "category": "Kopi",
      "images": ["https://example.com/product1.jpg"],
      "sku": "KOPI-001",
      "weight": 500,
      "unit": "gram",
      "isAvailable": true,
      "isLocalOnly": true,
      "createdAt": "2026-05-19T10:00:00Z",
      "updatedAt": "2026-05-19T10:00:00Z"
    }
  ]
}
```

---

### Get Product by ID
**GET** `/products/{id}` [Public]

Retrieve specific product information.

#### Parameters
- `id` (path, required): Product ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Product retrieved successfully",
  "data": {
    "id": "664a1f2c7d9a2c0012345679",
    "sellerId": "664a1f2c7d9a2c0012345678",
    "name": "Kopi Premium Arabika",
    "description": "Biji kopi arabika pilihan dari Aceh",
    "price": 150000,
    "quantity": 50,
    "category": "Kopi",
    "images": ["https://example.com/product1.jpg"],
    "sku": "KOPI-001",
    "weight": 500,
    "unit": "gram",
    "isAvailable": true,
    "isLocalOnly": true,
    "createdAt": "2026-05-19T10:00:00Z",
    "updatedAt": "2026-05-19T10:00:00Z"
  }
}
```

---

### Search Products
**GET** `/products/search` [Public]

Search products by query.

#### Parameters
- `query` (query, required): Search query string

#### Example
```
GET /api/products/search?query=kopi
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Products found",
  "data": [
    {
      "id": "664a1f2c7d9a2c0012345679",
      "name": "Kopi Premium Arabika",
      "price": 150000,
      ...
    }
  ]
}
```

---

### Create Product
**POST** `/products` [Protected - Seller/Admin]

Add a new product.

#### Request
```json
{
  "sellerId": "664a1f2c7d9a2c0012345678",
  "name": "Kopi Premium Arabika",
  "description": "Biji kopi arabika pilihan dari Aceh",
  "price": 150000,
  "quantity": 50,
  "category": "Kopi",
  "images": ["https://example.com/product1.jpg"],
  "sku": "KOPI-001",
  "weight": 500,
  "unit": "gram",
  "isAvailable": true,
  "isLocalOnly": true
}
```

#### Response (201 Created)
```json
{
  "success": true,
  "message": "Product created successfully",
  "data": {
    "id": "664a1f2c7d9a2c0012345679",
    "sellerId": "664a1f2c7d9a2c0012345678",
    "name": "Kopi Premium Arabika",
    ...
  }
}
```

---

### Update Product
**PUT** `/products/{id}` [Protected - Seller/Admin]

Update product information.

#### Parameters
- `id` (path, required): Product ID

#### Request
```json
{
  "name": "Kopi Premium Arabika Updated",
  "description": "Updated description",
  "price": 160000,
  "quantity": 45
}
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Product updated successfully",
  "data": {
    "id": "664a1f2c7d9a2c0012345679",
    "name": "Kopi Premium Arabika Updated",
    ...
  }
}
```

---

### Delete Product
**DELETE** `/products/{id}` [Protected - Seller/Admin]

Remove a product.

#### Parameters
- `id` (path, required): Product ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Product deleted successfully"
}
```

---

## Order Endpoints

### Create Order
**POST** `/orders` [Protected]

Create a new order with items.

#### Request
```json
{
  "userId": "664a1f2c7d9a2c0012345678",
  "items": [
    {
      "productId": "664a1f2c7d9a2c0012345679",
      "quantity": 2,
      "unitPrice": 150000,
      "subtotal": 300000
    }
  ],
  "subtotal": 300000,
  "tax": 30000,
  "shippingCost": 25000,
  "total": 355000,
  "notes": "Kirim pagi hari saja"
}
```

#### Response (201 Created)
```json
{
  "success": true,
  "message": "Order created successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567a",
    "userId": "664a1f2c7d9a2c0012345678",
    "status": "pending",
    "items": [
      {
        "productId": "664a1f2c7d9a2c0012345679",
        "quantity": 2,
        "unitPrice": 150000,
        "subtotal": 300000
      }
    ],
    "subtotal": 300000,
    "tax": 30000,
    "shippingCost": 25000,
    "total": 355000,
    "createdAt": "2026-05-19T10:00:00Z"
  }
}
```

---

### Get Order by ID
**GET** `/orders/{id}` [Protected]

Retrieve order details.

#### Parameters
- `id` (path, required): Order ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Order retrieved successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567a",
    "userId": "664a1f2c7d9a2c0012345678",
    "status": "pending",
    "items": [...],
    "subtotal": 300000,
    "tax": 30000,
    "shippingCost": 25000,
    "total": 355000,
    "createdAt": "2026-05-19T10:00:00Z"
  }
}
```

---

### Get Orders by User
**GET** `/orders/buyer` [Protected]

Retrieve orders for a specific user.

#### Parameters
- `userId` (query, required): User ID

#### Example
```
GET /api/orders/buyer?userId=664a1f2c7d9a2c0012345678
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Orders retrieved successfully",
  "data": [
    {
      "id": "664a1f2c7d9a2c001234567a",
      "userId": "664a1f2c7d9a2c0012345678",
      "status": "pending",
      ...
    }
  ]
}
```

---

### Update Order Status
**PUT** `/orders/{id}/status` [Protected]

Update order status.

#### Parameters
- `id` (path, required): Order ID
- `status` (query, required): New status (pending, confirmed, processing, shipped, delivered, cancelled, refunded)

#### Example
```
PUT /api/orders/664a1f2c7d9a2c001234567a/status?status=confirmed
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Order status updated successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567a",
    "status": "confirmed",
    ...
  }
}
```

---

### Cancel Order
**DELETE** `/orders/{id}` [Protected]

Cancel an existing order.

#### Parameters
- `id` (path, required): Order ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Order cancelled successfully"
}
```

---

## Marketplace Endpoints

### Get All Marketplaces
**GET** `/marketplaces` [Public]

Retrieve all available marketplaces.

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Marketplaces retrieved successfully",
  "data": [
    {
      "id": "664a1f2c7d9a2c001234567b",
      "name": "Pasar Seni Bandung",
      "description": "Marketplace untuk kesenian lokal",
      "logoUrl": "https://example.com/logo.jpg",
      "websiteUrl": "https://pasarseni.local",
      "contactEmail": "info@pasarseni.local",
      "contactPhone": "+628123456789",
      "address": "Jl. Gatot Subroto No. 123",
      "isActive": true,
      "createdAt": "2026-05-19T10:00:00Z"
    }
  ]
}
```

---

### Get Marketplace by ID
**GET** `/marketplaces/{id}` [Public]

Retrieve specific marketplace.

#### Parameters
- `id` (path, required): Marketplace ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Marketplace retrieved successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567b",
    "name": "Pasar Seni Bandung",
    ...
  }
}
```

---

### Create Marketplace
**POST** `/marketplaces` [Protected - Seller/Admin]

Add a new marketplace.

#### Request
```json
{
  "name": "Pasar Seni Bandung",
  "description": "Marketplace untuk kesenian lokal",
  "logoUrl": "https://example.com/logo.jpg",
  "websiteUrl": "https://pasarseni.local",
  "contactEmail": "info@pasarseni.local",
  "contactPhone": "+628123456789",
  "address": "Jl. Gatot Subroto No. 123",
  "isActive": true
}
```

#### Response (201 Created)
```json
{
  "success": true,
  "message": "Marketplace created successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567b",
    "name": "Pasar Seni Bandung",
    ...
  }
}
```

---

### Update Marketplace
**PUT** `/marketplaces/{id}` [Protected - Seller/Admin]

Update marketplace information.

#### Parameters
- `id` (path, required): Marketplace ID

#### Request
```json
{
  "name": "Pasar Seni Bandung Updated",
  "isActive": true
}
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Marketplace updated successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567b",
    "name": "Pasar Seni Bandung Updated",
    ...
  }
}
```

---

### Delete Marketplace
**DELETE** `/marketplaces/{id}` [Protected - Seller/Admin]

Remove a marketplace.

#### Parameters
- `id` (path, required): Marketplace ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Marketplace deleted successfully"
}
```

---

## Seller Endpoints

### Get All Sellers
**GET** `/sellers` [Public]

Retrieve all sellers.

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Sellers retrieved successfully",
  "data": [
    {
      "id": "664a1f2c7d9a2c001234567c",
      "userId": "664a1f2c7d9a2c0012345678",
      "shopName": "Kopi Aceh Premium",
      "shopDescription": "Penjual kopi premium dari Aceh",
      "shopImageUrl": "https://example.com/shop.jpg",
      "shopAddress": "Jl. Merdeka 45",
      "shopPhone": "+628123456789",
      "categories": ["Kopi", "Cokelat"],
      "rating": 4.5,
      "totalReviews": 150,
      "totalProducts": 45,
      "isVerified": true,
      "isActive": true,
      "isOnline": true,
      "createdAt": "2026-05-19T10:00:00Z"
    }
  ]
}
```

---

### Get Seller by ID
**GET** `/sellers/{id}` [Public]

Retrieve specific seller.

#### Parameters
- `id` (path, required): Seller ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Seller retrieved successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567c",
    "userId": "664a1f2c7d9a2c0012345678",
    "shopName": "Kopi Aceh Premium",
    ...
  }
}
```

---

### Get Nearest Stores
**GET** `/sellers/nearest` [Public]

Retrieve nearby stores based on coordinates.

#### Parameters
- `lat` (query, required): Latitude
- `lon` (query, required): Longitude
- `limit` (query, optional): Number of results (default: 10)

#### Example
```
GET /api/sellers/nearest?lat=-6.2088&lon=106.8456&limit=5
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Nearest stores retrieved successfully",
  "data": [
    {
      "id": "664a1f2c7d9a2c001234567c",
      "shopName": "Kopi Aceh Premium",
      "rating": 4.5,
      ...
    }
  ]
}
```

---

### Create Seller
**POST** `/sellers` [Protected - Buyer/Seller/Admin]

Register as a seller.

#### Request
```json
{
  "userId": "664a1f2c7d9a2c0012345678",
  "shopName": "Kopi Aceh Premium",
  "shopDescription": "Penjual kopi premium dari Aceh",
  "shopImageUrl": "https://example.com/shop.jpg",
  "shopAddress": "Jl. Merdeka 45",
  "shopPhone": "+628123456789",
  "categories": ["Kopi", "Cokelat"],
  "isActive": true
}
```

#### Response (201 Created)
```json
{
  "success": true,
  "message": "Seller created successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567c",
    "userId": "664a1f2c7d9a2c0012345678",
    "shopName": "Kopi Aceh Premium",
    ...
  }
}
```

---

### Update Seller
**PUT** `/sellers/{id}` [Protected - Seller/Admin]

Update seller information.

#### Parameters
- `id` (path, required): Seller ID

#### Request
```json
{
  "shopName": "Kopi Aceh Premium Updated",
  "rating": 4.6,
  "isActive": true
}
```

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Seller updated successfully",
  "data": {
    "id": "664a1f2c7d9a2c001234567c",
    "shopName": "Kopi Aceh Premium Updated",
    ...
  }
}
```

---

### Delete Seller
**DELETE** `/sellers/{id}` [Protected - Seller/Admin]

Remove a seller.

#### Parameters
- `id` (path, required): Seller ID

#### Response (200 OK)
```json
{
  "success": true,
  "message": "Seller deleted successfully"
}
```

---

## System Endpoints

### Health Check
**GET** `/health` [Public]

Check API health status.

#### Response (200 OK)
```json
{
  "status": "ok"
}
```

---

## Error Codes

| Status | Message                     | Description                               |
| ------ | --------------------------- | ----------------------------------------- |
| 400    | Invalid request body        | Malformed JSON or missing required fields |
| 401    | Missing authorization token | Authentication required but not provided  |
| 401    | Invalid token               | Token is invalid or expired               |
| 403    | Insufficient permissions    | User role doesn't have access             |
| 404    | Not found                   | Resource doesn't exist                    |
| 500    | Internal server error       | Server error occurred                     |

---

## cURL Examples

### Register
```bash
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alya Putri",
    "email": "alya@example.com",
    "password": "Secret123!",
    "roles": ["buyer"]
  }'
```

### Login
```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alya@example.com",
    "password": "Secret123!"
  }'
```

### Get Products
```bash
curl -X GET http://localhost:3000/api/products
```

### Create Product (with token)
```bash
curl -X POST http://localhost:3000/api/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "sellerId": "664a1f2c7d9a2c0012345678",
    "name": "Kopi Premium",
    "description": "Premium coffee",
    "price": 150000,
    "quantity": 50,
    "category": "Kopi"
  }'
```

### Search Products
```bash
curl -X GET "http://localhost:3000/api/products/search?query=kopi"
```

### Get Nearest Sellers
```bash
curl -X GET "http://localhost:3000/api/sellers/nearest?lat=-6.2088&lon=106.8456&limit=5"
```

---

## Notes

1. **JWT Token**: Valid for 24 hours from creation
2. **Password**: Minimum requirements should be enforced on client side
3. **Roles**: Available roles are `buyer`, `seller`, `admin`
4. **Timestamps**: All timestamps are in ISO 8601 format (UTC)
5. **Pagination**: Not yet implemented; all list endpoints return all records
6. **CORS**: Configure as needed for frontend integration

---

## Swagger/OpenAPI

The API documentation is also available in OpenAPI 3.0 format in `swagger.yaml` file.

To view the Swagger UI, use:
- [Swagger Editor](https://editor.swagger.io/)
- [ReDoc](https://redoc.ly/)
- Or integrate with go-swagger as documented at [goswagger.io](https://goswagger.io/)

### Generate Go code from Swagger:
```bash
swagger generate server -f swagger.yaml -t api -n pade
```

### Generate client:
```bash
swagger generate client -f swagger.yaml -t api -n pade
```
