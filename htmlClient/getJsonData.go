package htmlClient

type QuotesJson struct {
	Id     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}
type ContainerJson struct {
	Quotes []QuotesJson `json:"quotes"`
	Total  int          `json:"total"`
	Skip   int          `json:"skip"`
	Limit  int          `json:"limit"`
}

func GetJsonData(url string, filename string) error {
	var err error
	res, err := GetURL(url)
	if err != nil {
		return err
	}
	fileLocation := "./file/" + filename
	_, err = SaveFile(fileLocation, res)
	if err != nil {
		return err
	}
	return nil
}
