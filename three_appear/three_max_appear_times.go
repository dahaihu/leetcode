package threeappear

func maxThreeAppearTimes(values []string) []string {
	var a, b, c string
	var ac, bc, cc int
	for _, value := range values {
		switch value {
		case a:
			ac++
		case b:
			bc++
		case c:
			cc++
		default:
			switch {
			case ac == 0:
				a = value
				ac = 1
			case bc == 0:
				b = value
				bc = 1
			case cc == 0:
				c = value
				cc = 1
			default:
				ac--
				bc--
				cc--
			}
		}
	}
	if a == "" {
		return []string{}
	}
	if b == "" {
		return []string{a}
	}
	if c == "" {
		return []string{a, b}
	}
	return []string{a, b, c}
}
