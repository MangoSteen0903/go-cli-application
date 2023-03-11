package htmlClient

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
