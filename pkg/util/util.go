package util

// GetProgramVersion - View README.md for version history
func GetProgramVersion(timeUnix int) int {
	if timeUnix >= 1663810435 {
		return 1
	}
	return 0
}
