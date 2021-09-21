package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sharpvik/log-go/v2"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	postsFolder = "posts"
	sliceSize   = 50
)

func init() {
	log.SetLevel(log.LevelDebug)
	makeSurePostsFolderExists(postsFolder)
}

func main() {
	e := router()
	e.Start(":http")
}

/* ROUTER */

func router() (e *echo.Echo) {
	e = echo.New()

	e.Use(
		middleware.CORS(),
	)

	e.POST("/post", post)
	e.GET("/fetch/:n", fetch)
	e.Static("/", "dist")

	return
}

func post(c echo.Context) error {
	log.Debug("new post")
	postText, _ := io.ReadAll(c.Request().Body)
	sanitized := bluemonday.UGCPolicy().Sanitize(
		strings.TrimSpace(string(postText)))
	if len(sanitized) == 0 {
		return c.String(http.StatusBadRequest, "won't publish empty post")
	}
	if err := storePost(sanitized); err != nil {
		return c.String(http.StatusInternalServerError, "failed to save post")
	}
	return c.String(http.StatusOK, "post accepted")
}

func fetch(c echo.Context) error {
	log.Debug("new fetch")
	n, err := strconv.ParseInt(c.Param("n"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "weird slice argument format")
	}

	dir, err := os.ReadDir(postsFolder)
	if err != nil {
		return c.String(http.StatusInternalServerError,
			"failed to get the list of posts")
	}

	totalPosts := int64(len(dir))
	to := max(totalPosts-n*sliceSize, 0)
	from := max(totalPosts-(n+1)*sliceSize, 0)
	sliced := dir[from:to]
	posts := make([]string, len(sliced))

	for i, post := range sliced {
		file, err := os.Open(path.Join(postsFolder, post.Name()))
		if err != nil {
			return c.String(http.StatusInternalServerError,
				"failed to read one of the posts")
		}

		text, err := io.ReadAll(file)
		if err != nil {
			return c.String(http.StatusInternalServerError,
				"failed to read one of the posts")
		}

		posts[i] = string(text)
	}

	return c.JSON(http.StatusOK, posts)
}

/* UTIL */

func doesDirExists(dir string) bool {
	file, err := os.Open(dir)

	if err != nil {
		return false
	}
	defer func() { _ = file.Close() }()

	stat, err := file.Stat()
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func makeSurePostsFolderExists(path string) {
	if doesDirExists(path) {
		return
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		panic(fmt.Errorf("%s: %s", err, path))
	}
}

func freshName() string {
	return path.Join(postsFolder, strconv.FormatInt(time.Now().Unix(), 10)+".txt")
}

func storePost(text string) (err error) {
	file, err := os.Create(freshName())
	if err != nil {
		return
	}
	_, err = file.WriteString(text)
	return
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
