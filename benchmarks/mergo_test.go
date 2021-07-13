package benchmarks

import (
	"github.com/germanoeich/nirn/discord"
	"github.com/imdario/mergo"
	"testing"
)

var mockUser = discord.User{
	ID:            1,
	Username:      "ASdasdasdasd",
	Discriminator: 1,
	Avatar:        "asdasda12304u1djni",
	Bot:           true,
	Locale:        "en",
	Verified:      true,
	Email:         "test@test.com",
	Flags:         0,
	PremiumType:   0,
	PublicFlags:   0,
}

var mockUserUpdate = discord.User{
	ID:            2,
	Username:      "ASdasdasdasd",
	Discriminator: 1452,
	Avatar:        "asdasda12304u1djni",
}

func BenchmarkMergo(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		userCopy := mockUser
		mergo.Merge(&userCopy, mockUserUpdate, mergo.WithOverride)
	}
}

func BenchmarkManual(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		userCopy := mockUser
		if mockUserUpdate.ID != 0 {
			userCopy.ID = mockUserUpdate.ID
		}
		if mockUserUpdate.Discriminator != 0 {
			userCopy.Discriminator = mockUserUpdate.Discriminator
		}
		if mockUserUpdate.Username != "" {
			userCopy.Username = mockUserUpdate.Username
		}
		if mockUserUpdate.Avatar != "" {
			userCopy.ID = mockUserUpdate.ID
		}
	}
}