package utils

func GetService(id int) string {
	switch id {
	case 1:
		return "Netflix"
	case 2:
		return "Disney"
	case 3:
		return "HBO MAX"
	case 4:
		return "Star+"
	case 5:
		return "Spotify"
	case 6:
		return "Prime Video"
	case 7:
		return "Crunchyroll"
	default:
		return ""
	}
}
