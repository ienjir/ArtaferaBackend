package models

type Argon2idHash struct {
	time    uint32 // time represents the number of passed over the specified memory.
	memory  uint32 // cpu memory to be used.
	threads uint8  // threads for parallelism aspect of the algorithm.
	keyLen  uint32 // keyLen of the generate hash key.
	saltLen uint32 // saltLen the length of the salt used.
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
