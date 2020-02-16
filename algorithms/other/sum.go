package other

func sum(list []int) int  {
	if len(list) == 0 {
		return 0
	}

	return list[0] + sum(list[1:len(list)])
}