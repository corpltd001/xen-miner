package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

type Argon2ID struct {
	time, memory, keyLen uint32
	thread               uint8
	salt                 []byte

	format string
}

func GetDefaultArgon2ID() Argon2ID {
	return Argon2ID{
		time:   defaultTimeDifficulty,
		memory: defaultMemoryDifficulty,
		keyLen: defaultKeyLength,
		thread: defaultThread,
		salt:   defaultSalt,
	}
}

func (id *Argon2ID) Hash(password []byte) []byte {
	return argon2.IDKey(password, id.salt, id.time, id.memory, id.thread, id.keyLen)
}

func (id *Argon2ID) HashStr(password []byte) string {
	return base64.RawStdEncoding.EncodeToString(id.Hash(password))
}

func (id *Argon2ID) Signature(password []byte) string {
	return fmt.Sprintf(id.GetFormat(), id.HashStr(password))
}

func (id *Argon2ID) Valid() bool {
	return id.memory >= uint32(id.thread*8)
}

func (id *Argon2ID) GetFormat() string {
	if id.format == "" {
		id.format = fmt.Sprintf(
			`$argon2id$v=19$m=%d,t=%d,p=%d$%s$%%s`,
			id.memory,
			id.time,
			id.thread,
			base64.RawStdEncoding.EncodeToString(id.salt),
		)
	}

	return id.format
}
