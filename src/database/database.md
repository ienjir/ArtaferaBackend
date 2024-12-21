# Database

- User
- Roles
- Texts
- Text Translations
- Art
- Art Translation
- Languages
- Pictures

## User

- UserID
- Firstname
- Lastname
- Email
- Phone (optional)
- Street
- Number
- City
- Postal code
- Password
- Last access
- Created at
- Updated at

| Key | Name       | Type         | Nullable | More           | 
|:---:|------------|--------------|----------|----------------|
| PK  | UserID     | Int          | No       | Auto increment |
|     | Firstname  | Varchar(255) | No       |                |
|     | Lastname   | Varchar(255) | No       |                |
|     | Email      | Varchar(255) | No       | Unique         |
|     | Phone      | Varchar(20)  | Yes      | Unique         |
|     | Street     | Varchar(255) | Yes      |                |
|     | Number     | Smallint     | Yes      |                |
|     | City       | Varchar(255) | Yes      |                |
|     | PostalCode | Varchar(32)  | Yes      |                |
|     | Password   | Char(40)     | No       |                |
|     | LastAccess | Time         | Yes      | Auto update    |
|     | CreatedAt  | Time         | Yes      | Auto update    |
|     | UpdatedAt  | Time         | Yes      | Auto update    |

## Roles

- RoleID
- Role
- Created at
- Updated at

| Key | Name      | Type        | Nullable | More           |
|:---:|-----------|-------------|----------|----------------|
| PK  | RoleID    | Int         | No       | Auto increment |
|     | Role      | Varchar(50) | No       |                |
|     | CreatedAt | Time        | Yes      | Auto update    |
|     | UpdatedAt | Time        | Yes      | Auto update    |

## Text

- TextID
- Created at
- Updated at

| Key | Name      | Type | Nullable | More           |
|:---:|-----------|------|----------|----------------|
| PK  | TextID    | Int  | No       | Auto increment |
|     | CreatedAt | Time | Yes      | Auto update    |
|     | UpdatedAt | Time | Yes      | Auto update    |

## Text Translations

- TranslationID
- TextID
- LanguageID
- Text
- Created at
- Updated at

| Key | Name          | Type | Nullable | More           |
|:---:|---------------|------|----------|----------------|
| PK  | TranslationID | Int  | No       | Auto Increment |
| FK  | TextID        | Int  | No       |                |  
| FK  | LanguageID    | Int  | No       |                |
|     | Text          | Text | No       |                |
|     | CreatedAt     | Time | Yes      | Auto update    |
|     | UpdatedAt     | Time | Yes      | Auto update    |

## Languages

- LanguageID
- Language name
- Language code
- Created At
- Updated At

| Key | Name         | Type        | Nullable | More           |
|:---:|--------------|-------------|----------|----------------|
| PK  | LanguageID   | Int         | No       | Auto Increment |
|     | LanguageName | Varchar(50) | No       | Unique         |
|     | LanguageCode | Char(3)     | No       | Unique         |
|     | CreatedAt    | Time        | Yes      | Auto update    |
|     | UpdatedAt    | Time        | Yes      | Auto update    |

## Art

- ArtID
- Price
- Currency
- Creation year
- Dimensions
- Created at
- Updated at

| Key | Name          | Type        | Nullable | More           |
|:---:|---------------|-------------|----------|----------------|
| PK  | ArtID         | Int         | No       | Auto Increment |
|     | Price         | Int         | No       |                |
|     | Currency      | Char(3)     | No       |                |
|     | Creation year | Char(4)     | No       |                |
|     | Dimensions    | Varchar(40) | Yes      |                |
|     | CreatedAt     | Time        | Yes      | Auto update    |
|     | UpdatedAt     | Time        | Yes      | Auto update    |

## Art translations

- TranslationID
- ArtID
- LanguageID
- Title
- Description
- Created at
- Updated at

| Key | Name          | Type         | Nullable | More           |
|:---:|---------------|--------------|----------|----------------|
| PK  | TranslationID | Int          | No       | Auto Increment |
| FK  | ArtID         | Int          | No       |                |
| FK  | LanguageID    | Int          | No       |                |
|     | Title         | Varchar(255) | No       | Unique         |
|     | Description   | Text         | No       |                |
|     | CreatedAt     | Time         | Yes      | Auto update    |
|     | UpdatedAt     | Time         | Yes      | Auto update    |

## Pictures

- PictureID
- PictureLink
- Created at
- Updated at

| Key | Name        | Type         | Nullable | More           |
|:---:|-------------|--------------|----------|----------------|
| PK  | PictureID   | Int          | No       | Auto Increment |
|     | PicutreLink | Varchar(255) | No       |                |
|     | CreatedAt   | Time         | Yes      | Auto update    |
|     | UpdatedAt   | Time         | Yes      | Auto update    |

