package threeappear

func maxThreeAppearTimes(values []string) []string {
	var a, b, c string
	var ac, bc, cc int
	for _, value := range values {
		switch value {
		case a:
			ac++
			continue
		case b:
			bc++
			continue
		case c:
			cc++
			continue
		}
		switch {
		case ac == 0:
			ac++
			a = value
		case bc == 0:
			bc++
			b = value
		case cc == 0:
			cc++
			c = value
		default:
			ac--
			bc--
			cc--
		}
	}
	if ac != 0 && bc != 0 && cc != 0 {
		return []string{a, b, c}
	}
	if ac != 0 && bc != 0 {
		return []string{a, b}
	}
	if ac != 0 {
		return []string{a}
	}
	return []string{}
}
