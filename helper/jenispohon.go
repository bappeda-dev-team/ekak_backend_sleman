package helper

func GetJenisPohon(level int) string {
	switch level {
	case 4:
		return "Strategic"
	case 5:
		return "Tactical"
	case 6:
		return "Operational"
	default:
		return ""
	}
}
