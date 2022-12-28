package api

import(
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		key := "AIzaSyDE3X4RTUY97rVZU2H3dkdU6MB2P6dqGVM"
		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(key))
		if err != nil {
			logrus.Fatalf("Error creating new Youtube service: %v", err)
		}
		call := yts.Videos.List([]string{"id", "snippet"}).Chart("mostPopular").MaxResults(3)
		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		return c.JSON(fasthttp.StatusOK, res)
	}
}
