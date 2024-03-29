package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Base(w http.ResponseWriter, r *http.Request) {

	images := make([]string, 7)

	for i := 0; i < 6; i++ {
		images = append(images, genImgBlock())
	}

	fmt.Fprint(w, cssHead)

	for _, image := range images {
		fmt.Fprint(w, image)
	}

	fmt.Fprint(w, htmlFooter)
}

func genImgBlock() string {
	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	svcName := os.Getenv("SVCNAME")
	svcPort := os.Getenv("SVCPORT")
	return `<div class="responsive">
		<div class="gallery">
  <a target="_blank" href="img_5terre.jpg">
    <img src="` + `http://` + svcName + `:` + svcPort + `/image?` + timeStr + `"alt="Cinque Terre" width="600" height="400">
  </a>
  <div class="desc">` + getExcuse() + `</div>
  </div>
</div>`
}
