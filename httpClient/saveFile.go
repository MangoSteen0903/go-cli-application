package httpClient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func SaveFile(filename string, res *http.Response) (ContainerJson, error) {
	var pkgData ContainerJson
	file, err := os.Create(filename)
	if err != nil {
		return pkgData, err
	}
	w := bufio.NewWriter(file)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pkgData, err
	}
	err = json.Unmarshal(data, &pkgData)
	if err != nil {
		return pkgData, err
	}
	quotes := pkgData.Quotes
	tempList := []string{}
	for _, quote := range quotes {
		data := fmt.Sprintf(`{"id":%v, "quote":"%v", "author":"%v"},`, quote.Id, quote.Quote, quote.Author)
		tempList = append(tempList, data)
	}
	_, err = w.WriteString(fmt.Sprintf("[%v]", strings.Join(tempList, "")))

	if err != nil {
		return pkgData, err
	}
	w.Flush()
	return pkgData, nil
}
