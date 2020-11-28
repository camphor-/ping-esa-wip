package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/hori-ryota/esa-go/esa"
)

const teamName = "camphor"

var (
	dryrun = flag.Bool("dry-run", false, "skip posting comment")
)

func main() {
	flag.Parse()

	token := os.Getenv("ESA_API_TOKEN")
	if token == "" {
		fmt.Println("Environment variable ESA_API_TOKEN is not set.")
		os.Exit(1)
		return
	}

	aMonthBefore := time.Now().Add(-30 * 24 * time.Hour)

	client := esa.NewClient(token, teamName)

	page := uint(0)
	for {

		nonUpdatedPosts, err := client.ListPosts(context.Background(), esa.ListPostsParam{
			Q: "updated:<" + aMonthBefore.Format("2006-01-02") + " wip:true -in:Users -in:Archived -in:ポエム",
		}, page, 100)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		for _, post := range nonUpdatedPosts.Posts {
			if err := postComment(client, post); err != nil {
				fmt.Println(err)
				os.Exit(1)
				return
			}
		}

		if nonUpdatedPosts.NextPage == nil {
			break
		}
		page = *nonUpdatedPosts.NextPage
	}

}

func postComment(client esa.Client, post esa.Post) error {
	if dryrun != nil && *dryrun {
		fmt.Printf("skip posting comment: #%d %s\n", post.Number, post.Name)
		return nil
	}

	bodyTemplate := `
@%s WIPのまま記事が1ヶ月以上放置されています！

記事を更新してShip Itしましょう！
`
	body := fmt.Sprintf(bodyTemplate, post.CreatedBy.ScreenName)
	botUser := "esa_bot"
	_, err := client.CreateComment(context.Background(), post.Number, esa.CreateCommentParam{
		BodyMD: body,
		User:   &botUser,
	})
	if err != nil {
		return fmt.Errorf("create comment: %w", err)
	}
	return nil
}
