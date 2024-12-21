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
| 

## Languages

- LanguageID
- Language name
- Language code

## Art

- ArtID
- Price
- Creation date
- Dimensions
- Created at
- Updated at

## Art translations

- TranslationID
- ArtID
- LanguageID
- Title
- Description
- Created at
- Updated at

## Pictures

- PictureID
- PictureLink
- Created at
- Updated at