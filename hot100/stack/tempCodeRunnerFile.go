r{}
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		nums.WriteByte(s[*ptr])
	}
	return nums.String()