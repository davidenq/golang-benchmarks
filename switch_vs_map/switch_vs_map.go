package switch_vs_map


func withSwitch(text string) string {
	message := ""
	switch text {
	case " 1":
		message = "one"
	case " 2":
		message = "two"
	case " 3":
		message = "three"
	case " 4":
		message = "four"
	case " 5":
		message = "five"
	case " 6":
		message = "six"
	case " 7":
		message = "seven"
	case " 8":
		message = "eight"
	case " 9":
		message = "nine"
	case "10":
		message = "ten"
	case "11":
		message = "eleven"
	case "12":
		message = "twelve"
	case "13":
		message = "thirteen"
	case "14":
		message = "fourteen"
	case "15":
		message = "fifteen"
	case "16":
		message = "sixteen"
	case "17":
		message = "seventeen"
	case "18":
		message = "eighteen"
	case "19":
		message = "nineteen"
	case "20":
		message = "twenty"
	}
	return message
}

func withMap(text string) string {
	message := map[string]string{
		" 1": "one",
		" 2": "two",
		" 3": "three",
		" 4": "four",
		" 5": "five",
		" 6": "six",
		" 7": "seven",
		" 8": "eight",
		" 9": "nine",
		"10": "ten",
		"11": "eleven",
		"12": "twelve",
		"13": "thirteen",
		"14": "fourteen",
		"15": "fifteen",
		"16": "sixteen",
		"17": "seventeen",
		"18": "eighteen",
		"19": "nineteen",
		"20": "twenty",
	}
	return message[text]
}

