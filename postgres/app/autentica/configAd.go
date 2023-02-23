package autentica

import auth "github.com/korylprince/go-ad-auth/v3"

func GetConfig() *auth.Config {
	adConfig = &auth.Config{
		Server:   "santacruz.local",
		Port:     389,
		BaseDN:   "dc=santacruz,dc=local",
		Security: auth.SecurityInsecureStartTLS,
	}
	return adConfig
}
