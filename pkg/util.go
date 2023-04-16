package drip

func GetIdlVersion(slot uint64) int {
	// First upgrade.
	// 2VZGDEZvBE8CUE88dPW2Bhs8inRoCf5ckGhpxTmRfzcxUSvgCRkSUujB9Br4xBqW1p9YrB8CoGgwTE3ZpM9wwR2W
	if slot >= uint64(0x90ba3c8) {
		return 1
	}

	// Initial deploy.
	// 5TJHBphZYdsoTTbu4uLWXBXH1UW574AMqz25TbjB4maHYYVQiswgSQhfRbUpySW24a9nqqQeQLX8QeHCi2s4abT3 
	if slot >= uint64(0x8ae175b) && slot < uint64(0x90ba3c8) {
		return 0
	}

	return 0
}