package static

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeVideo(c *gin.Context) {
	filepath := path.Join("static/videos", path.Clean(c.Param("filepath")))

	file, err := os.Open(filepath)
	if err != nil || !strings.HasSuffix(".mp4", filepath) {
		c.String(http.StatusNotFound, "Invalid file name")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil || stat.IsDir() {
		c.String(http.StatusInternalServerError, "Invalid file")
		return
	}

	http.ServeContent(c.Writer, c.Request, stat.Name(), stat.ModTime(), file)
}
