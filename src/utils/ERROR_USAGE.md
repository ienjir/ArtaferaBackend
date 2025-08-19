# Universal Error Messaging System

This document explains how to use the universal error messaging system implemented in `utils/errors.go`.

## Overview

The system provides:
- **Consistent error messages** across the entire application
- **Standardized response formats** for both errors and success responses
- **Easy-to-use helper functions** for common error scenarios
- **Centralized error management** for better maintainability

## Error Response Structure

All errors now follow this consistent format:

```json
{
  "error": "Error message",
  "code": 400,
  "details": "Optional additional details"
}
```

Success responses follow this format:

```json
{
  "message": "Optional success message",
  "data": { /* response data */ }
}
```

## Usage Examples

### 1. In Controllers

**Before:**
```go
if err := c.ShouldBindJSON(&json); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}
```

**After:**
```go
if err := c.ShouldBindJSON(&json); err != nil {
    utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
    return
}
```

### 2. Handling Service Errors

**Before:**
```go
if err != nil {
    c.JSON(err.StatusCode, gin.H{"error": err.Message})
    return
}
```

**After:**
```go
if err != nil {
    utils.RespondWithServiceError(c, err)
    return
}
```

### 3. Success Responses

**Before:**
```go
c.JSON(http.StatusOK, gin.H{"user": user})
```

**After:**
```go
utils.RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
// OR with a message:
utils.RespondWithSuccess(c, http.StatusOK, nil, "User successfully deleted")
```

## Available Error Constants

### Generic Errors
- `ErrInternalServer` - "Internal server error"
- `ErrInvalidJSON` - "Invalid JSON format"
- `ErrInvalidID` - "Invalid ID format"
- `ErrResourceNotFound` - "Resource not found"
- `ErrUnauthorized` - "Unauthorized access"
- `ErrForbidden` - "Access forbidden"
- `ErrBadRequest` - "Bad request"

### Field Validation
- `ErrFieldRequired` - " is required"
- `ErrFieldEmpty` - " cannot be empty"
- `ErrFieldInvalid` - " is invalid"
- `ErrFieldOutOfRange` - " is out of range"

### Authentication/Authorization
- `ErrInvalidCredentials` - "Invalid email or password"
- `ErrTokenExpired` - "Token has expired"
- `ErrTokenInvalid` - "Invalid token"
- `ErrRefreshTokenRequired` - "Refresh token is required"
- `ErrAdminRequired` - "Admin role required"
- `ErrOwnerAccess` - "You can only access your own resources"

### Business Logic
- `ErrPasswordInsecure` - "Password does not meet security requirements"
- `ErrEmailInvalid` - "Email format is invalid"
- `ErrPhoneInvalid` - "Phone number format is invalid"

## Helper Functions

### Error Constructors

Create service errors quickly:

```go
// For different HTTP status codes
err := utils.NewBadRequestError("Custom message")
err := utils.NewUnauthorizedError("Custom message")
err := utils.NewForbiddenError("Custom message")
err := utils.NewNotFoundError("Custom message")
err := utils.NewInternalServerError("Custom message")

// For common field validation errors
err := utils.NewFieldRequiredError("Email")      // "Email is required"
err := utils.NewFieldEmptyError("Password")      // "Password cannot be empty"
err := utils.NewFieldInvalidError("Status")      // "Status is invalid"
err := utils.NewFieldOutOfRangeError("Age", "between 18 and 100") // "Age must be between 18 and 100"

// For common business errors
err := utils.NewInvalidCredentialsError()
err := utils.NewAdminRequiredError()
err := utils.NewOwnerAccessError()
err := utils.NewPasswordInsecureError()
```

### Response Functions

```go
// Send error response
utils.RespondWithError(c, http.StatusBadRequest, "Custom error message")
utils.RespondWithError(c, http.StatusBadRequest, "Error message", "Additional details")

// Send service error response
utils.RespondWithServiceError(c, serviceError)

// Send success response
utils.RespondWithSuccess(c, http.StatusOK, data)
utils.RespondWithSuccess(c, http.StatusOK, data, "Success message")
```

## Migration Guide

### Step 1: Update Imports
Add the utils import to your files:
```go
import "github.com/ienjir/ArtaferaBackend/src/utils"
```

### Step 2: Replace Error Responses
Replace all `c.JSON(statusCode, gin.H{"error": message})` with appropriate utils functions.

### Step 3: Replace Success Responses
Replace `c.JSON(statusCode, gin.H{...})` with `utils.RespondWithSuccess(c, statusCode, data)`.

### Step 4: Update Validation Functions
The validation package has been updated to use the new error utilities. Ensure your validation calls are compatible.

## Benefits

1. **Consistency**: All errors follow the same format
2. **Maintainability**: Change error messages in one place
3. **Type Safety**: Reduced risk of typos in error messages
4. **Better UX**: Consistent error structure for frontend consumption
5. **Debugging**: Centralized error handling makes debugging easier
6. **API Standards**: Follows REST API best practices

## Best Practices

1. **Use appropriate HTTP status codes** - Don't use 500 for validation errors
2. **Provide meaningful messages** - Help users understand what went wrong
3. **Don't expose sensitive information** - Use generic messages for security errors
4. **Be consistent** - Always use the utils functions instead of manual JSON responses
5. **Add details sparingly** - Only include details when they provide value to the client