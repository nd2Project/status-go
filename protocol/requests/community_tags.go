package requests

var TagsEmojies map[string]string

func init() {
	TagsEmojies = make(map[string]string)
	TagsEmojies["Activism"] = "โ"
	TagsEmojies["Art"] = "๐จ"
	TagsEmojies["Blockchain"] = "๐"
	TagsEmojies["Books & blogs"] = "๐"
	TagsEmojies["Career"] = "๐ผ"
	TagsEmojies["Collaboration"] = "๐ค"
	TagsEmojies["Commerce"] = "๐"
	TagsEmojies["Culture"] = "๐"
	TagsEmojies["DAO"] = "๐"
	TagsEmojies["DeFi"] = "๐"
	TagsEmojies["Design"] = "๐งฉ"
	TagsEmojies["DIY"] = "๐จ"
	TagsEmojies["Environment"] = "๐ฟ"
	TagsEmojies["Education"] = "๐"
	TagsEmojies["Entertainment"] = "๐ฟ"
	TagsEmojies["Ethereum"] = "ฮ"
	TagsEmojies["Event"] = "๐"
	TagsEmojies["Fantasy"] = "๐งโโ๏ธ"
	TagsEmojies["Fashion"] = "๐งฆ"
	TagsEmojies["Food"] = "๐ถ"
	TagsEmojies["Gaming"] = "๐ฎ"
	TagsEmojies["Global"] = "๐"
	TagsEmojies["Health"] = "๐ง "
	TagsEmojies["Hobby"] = "๐"
	TagsEmojies["Innovation"] = "๐งช"
	TagsEmojies["Language"] = "๐"
	TagsEmojies["Lifestyle"] = "โจ"
	TagsEmojies["Local"] = "๐"
	TagsEmojies["Love"] = "โค๏ธ"
	TagsEmojies["Markets"] = "๐"
	TagsEmojies["Movies & TV"] = "๐"
	TagsEmojies["Music"] = "๐ถ"
	TagsEmojies["News"] = "๐"
	TagsEmojies["NFT"] = "๐ผ"
	TagsEmojies["Non-profit"] = "๐"
	TagsEmojies["NSFW"] = "๐"
	TagsEmojies["Org"] = "๐ข"
	TagsEmojies["Pets"] = "๐ถ"
	TagsEmojies["Play"] = "๐ฒ"
	TagsEmojies["Podcast"] = "๐๏ธ"
	TagsEmojies["Politics"] = "๐ณ๏ธ"
	TagsEmojies["Product"] = "๐ฑ"
	TagsEmojies["Psyche"] = "๐"
	TagsEmojies["Privacy"] = "๐ป"
	TagsEmojies["Security"] = "๐"
	TagsEmojies["Social"] = "โ"
	TagsEmojies["Software dev"] = "๐ฉโ๐ป"
	TagsEmojies["Sports"] = "โฝ๏ธ"
	TagsEmojies["Tech"] = "๐ฑ"
	TagsEmojies["Travel"] = "๐บ"
	TagsEmojies["Vehicles"] = "๐"
	TagsEmojies["Web3"] = "๐"
}

func ValidateTags(input []string) bool {
	for _, t := range input {
		_, ok := TagsEmojies[t]
		if !ok {
			return false
		}
	}

	if len(unique(input)) != len(input) {
		// Contains duplicates. Shouldn't have happened
		return false
	}

	return true
}

func RemoveUnknownAndDeduplicateTags(input []string) []string {
	var result []string
	for _, t := range input {
		_, ok := TagsEmojies[t]
		if ok {
			result = append(result, t)
		}
	}
	return unique(result)
}

func unique(slice []string) []string {
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
