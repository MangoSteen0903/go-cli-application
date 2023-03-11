package httpClient

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

type PkgRegisterResult struct {
	Id int `json:"id"`
}
