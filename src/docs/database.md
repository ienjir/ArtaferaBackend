# Database

## Tables Overview

- User
- Roles
- UserRoles
- Texts
- Translations
- Languages
- Art
- ArtPictures
- Pictures
- Orders
- OrderDetails
- Payments
- Currencies

## User

| Key | Name       | Type         | Nullable | More           |
|:---:|------------|--------------|----------|----------------|
| PK  | UserID     | Int          | No       | Auto increment |
|     | Firstname  | Varchar(255) | No       |                |
|     | Lastname   | Varchar(255) | No       |                |
|     | Email      | Varchar(255) | No       | Unique         |
|     | Phone      | Varchar(20)  | Yes      |                |
|     | Address1   | Varchar(255) | Yes      |                |
|     | Address2   | Varchar(255) | Yes      |                |
|     | City       | Varchar(255) | Yes      |                |
|     | PostalCode | Varchar(32)  | Yes      |                |
|     | Password   | Binary(60)   | No       | Hashed         |
|     | LastAccess | Datetime     | Yes      | Auto update    |
|     | CreatedAt  | Datetime     | Yes      | Auto update    |
|     | UpdatedAt  | Datetime     | Yes      | Auto update    |
|     | IsDeleted  | Boolean      | Yes      | Default false  |

## Roles

| Key | Name      | Type        | Nullable | More           |
|:---:|-----------|-------------|----------|----------------|
| PK  | RoleID    | Int         | No       | Auto increment |
|     | Role      | Varchar(50) | No       |                |
|     | CreatedAt | Datetime    | Yes      | Auto update    |
|     | UpdatedAt | Datetime    | Yes      | Auto update    |

## UserRoles

| Key | Name       | Type     | Nullable | More           |
|:---:|------------|----------|----------|----------------|
| PK  | UserRoleID | Int      | No       | Auto increment |
| FK  | UserID     | Int      | No       |                |
| FK  | RoleID     | Int      | No       |                |
|     | CreatedAt  | Datetime | Yes      | Auto update    |
|     | UpdatedAt  | Datetime | Yes      | Auto update    |

## Texts

| Key | Name      | Type     | Nullable | More           |
|:---:|-----------|----------|----------|----------------|
| PK  | TextID    | Int      | No       | Auto increment |
|     | CreatedAt | Datetime | Yes      | Auto update    |
|     | UpdatedAt | Datetime | Yes      | Auto update    |

## Translations

| Key | Name          | Type        | Nullable | More              |
|:---:|---------------|-------------|----------|-------------------|
| PK  | TranslationID | Int         | No       | Auto Increment    |
| FK  | EntityID      | Int         | No       | Links Text or Art |
| FK  | LanguageID    | Int         | No       |                   |
|     | Context       | Varchar(50) | No       | Describes field   |
|     | Text          | Text        | No       |                   |
|     | CreatedAt     | Datetime    | Yes      | Auto update       |
|     | UpdatedAt     | Datetime    | Yes      | Auto update       |

## Languages

| Key | Name         | Type        | Nullable | More           |
|:---:|--------------|-------------|----------|----------------|
| PK  | LanguageID   | Int         | No       | Auto Increment |
|     | LanguageName | Varchar(50) | No       | Unique         |
|     | LanguageCode | Char(3)     | No       | Unique         |
|     | CreatedAt    | Datetime    | Yes      | Auto update    |
|     | UpdatedAt    | Datetime    | Yes      | Auto update    |

## Art

| Key | Name         | Type         | Nullable | More                |
|:---:|--------------|--------------|----------|---------------------|
| PK  | ArtID        | Int          | No       | Auto Increment      |
|     | Price        | Int          | No       |                     |
| FK  | CurrencyID   | Int          | No       | References Currency |
|     | CreationYear | Char(4)      | No       |                     |
|     | Width        | Decimal(8,2) | Yes      |                     |
|     | Height       | Decimal(8,2) | Yes      |                     |
|     | Depth        | Decimal(8,2) | Yes      |                     |
|     | CreatedAt    | Datetime     | Yes      | Auto update         |
|     | UpdatedAt    | Datetime     | Yes      | Auto update         |

## ArtPictures

| Key | Name         | Type     | Nullable | More           |
|:---:|--------------|----------|----------|----------------|
| PK  | ArtPictureID | Int      | No       | Auto Increment |
| FK  | ArtID        | Int      | No       |                |
| FK  | PictureID    | Int      | No       |                |
|     | Priority     | Int      | Yes      | Sorting Order  |
|     | CreatedAt    | Datetime | Yes      | Auto update    |
|     | UpdatedAt    | Datetime | Yes      | Auto update    |

## Pictures

| Key | Name        | Type         | Nullable | More           |
|:---:|-------------|--------------|----------|----------------|
| PK  | PictureID   | Int          | No       | Auto Increment |
|     | PictureLink | Varchar(255) | No       |                |
|     | CreatedAt   | Datetime     | Yes      | Auto update    |
|     | UpdatedAt   | Datetime     | Yes      | Auto update    |

## Orders

| Key | Name       | Type          | Nullable | More           |
|:---:|------------|---------------|----------|----------------|
| PK  | OrderID    | Int           | No       | Auto increment |
| FK  | UserID     | Int           | No       |                |
|     | OrderDate  | Datetime      | No       |                |
|     | TotalPrice | Decimal(10,2) | No       |                |
|     | Status     | Varchar(50)   | No       |                |
|     | CreatedAt  | Datetime      | Yes      | Auto update    |
|     | UpdatedAt  | Datetime      | Yes      | Auto update    |

## OrderDetails

| Key | Name          | Type          | Nullable | More           |
|:---:|---------------|---------------|----------|----------------|
| PK  | OrderDetailID | Int           | No       | Auto Increment |
| FK  | OrderID       | Int           | No       |                |
| FK  | ArtID         | Int           | No       |                |
|     | Quantity      | Int           | No       |                |
|     | Price         | Decimal(10,2) | No       |                |
|     | CreatedAt     | Datetime      | Yes      | Auto update    |
|     | UpdatedAt     | Datetime      | Yes      | Auto update    |

## Payments

| Key | Name          | Type          | Nullable | More           |
|:---:|---------------|---------------|----------|----------------|
| PK  | PaymentID     | Int           | No       | Auto increment |
| FK  | OrderID       | Int           | No       |                |
|     | PaymentDate   | Datetime      | No       |                |
|     | Amount        | Decimal(10,2) | No       |                |
|     | PaymentMethod | Varchar(50)   | No       |                |
|     | Status        | Varchar(50)   | No       |                |
|     | CreatedAt     | Datetime      | Yes      | Auto update    |
|     | UpdatedAt     | Datetime      | Yes      | Auto update    |

## Currencies

| Key | Name         | Type        | Nullable | More           |
|:---:|--------------|-------------|----------|----------------|
| PK  | CurrencyID   | Int         | No       | Auto Increment |
|     | CurrencyCode | Char(3)     | No       | Unique         |
|     | CurrencyName | Varchar(50) | No       |                |
|     | CreatedAt    | Datetime    | Yes      | Auto update    |
|     | UpdatedAt    | Datetime    | Yes      | Auto update    |