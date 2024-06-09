# Bima TRI API

This API provides an interface to interact with the Bima TRI service.

## Table of Contents

- [Features](#features)
- [Endpoints](#endpoints)
- [Swagger Documentation](#swagger-documentation)
- [Contributing](#contributing)

## Features

- **Authentication:** Login and OTP-based authentication for secure access.
- **User Profile:** Retrieve detailed user profile information.
- **Promo Information:** Get the latest promos and offers.
- **Transaction History:** View and filter transaction history by category and status.

## Endpoints

### Authentication
- `POST /api/v1/auth/login`: Request an OTP for login.
- `POST /api/v1/auth/login-otp`: Login with the received OTP.

### User
- `POST /api/v1/profile`: Retrieve user profile information.

### Promo
- `POST /api/v1/promo`: Retrieve promo information.

### Transaction
- `POST /api/v1/transaction`: Retrieve transaction history with optional filters (category, status).

## Swagger Documentation
The API is fully documented using Swagger.  Access the interactive documentation after starting the server:

URL: http://bima.thxrhmn.cloud/swagger/index.html

## Contributing
Contributions are welcome! Please feel free to submit issues, pull requests, or suggestions.