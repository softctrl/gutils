package scgotils

func IsIpAddress(__ip string) bool {
	return IP_ADDRESS_PATTERN.MatchString(__ip)
}
