package flowermanga

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/breno5g/mangad/internal/helpers"
)

func GetMangaChapters(url string) {
	res, err := http.Get(url)
	helpers.CheckError(err)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	helpers.CheckError(err)

	lines, err := helpers.FindPattern(string(body), helpers.FlowerMangaChapterPattern)
	helpers.CheckError(err)

	lines, err = helpers.FindPattern(strings.Join(lines, ","), helpers.UrlPattern)
	helpers.CheckError(err)

	urlArray := helpers.ReverseArray(lines)

	for _, line := range urlArray {
		fmt.Println(line)
	}
}
