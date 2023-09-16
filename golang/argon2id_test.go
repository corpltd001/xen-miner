package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var (
	argon2idDefault = GetDefaultArgon2ID()

	signatureNotContainXEN11Password = []byte(`997609015e4cc106fe2a5971f0c2d3bf2b471fe21e3f3989761a451392a306c88295b6a1cfccd9dacba568ad6abf04cbd375b4c049908993e5b6a26c49437e7e`)
	signatureNotContainXEN11Hash     = argon2idDefault.Hash(signatureNotContainXEN11Password)
	signatureNotContainXEN11         = argon2idDefault.Signature(signatureNotContainXEN11Password)

	signatureContainXEN11Password = []byte(`bbc2d7801619bb2f96fee0fe18024f34ca050ff4d49f85054d321f79a36691687e8735225603782aabb360c803003e12f26fe2bed7ae83ca1d1c90bd7bca3d94`)
	signatureContainXEN11Hash     = argon2idDefault.Hash(signatureContainXEN11Password)
	signatureContainXEN11         = argon2idDefault.Signature(signatureContainXEN11Password)
)

const (
	signatureNotContainXEN11Expected = `$argon2id$v=19$m=8,t=1,p=1$WEVOMTAwODIwMjJYRU4$cwLU76JXyYqx08SWYEwzNogoolZoL9BnXTmaXcpyZQWsjniK7/KRoAGzxTQgTsxRWoqeLyR2YqZ9k3lpz4Lx9g`
	signatureContainXEN11Expected    = `$argon2id$v=19$m=8,t=1,p=1$WEVOMTAwODIwMjJYRU4$cwLU76JXyYqx08SWYEwzNogoolZoL9BnXTmaXcpyZQWsjniK7/KRoAGzxTQgTsxRWoqeLyR2YqZ9k3lpz4Lx9g`
)

func TestMiner_Hash(t *testing.T) {
	assert.Equal(t, signatureNotContainXEN11Expected, signatureNotContainXEN11)

	assert.False(t, strings.Contains(signatureNotContainXEN11, defaultTargetXEN11))
}

func BenchmarkFullSignatureContain(b *testing.B) {
	fromIndex := len(signatureNotContainXEN11) - 86
	for i := 0; i < b.N; i++ {
		strings.Contains(signatureNotContainXEN11[fromIndex:], defaultTargetXEN11)
	}
}

func BenchmarkSignatureContain(b *testing.B) {
	hashStr := argon2idDefault.HashStr(signatureNotContainXEN11Password)
	for i := 0; i < b.N; i++ {
		strings.Contains(hashStr, defaultTargetXEN11)
	}
}
