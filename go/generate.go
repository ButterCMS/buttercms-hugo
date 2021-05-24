package main

import (
	"fmt"
	"os"
	"strconv"

	md "github.com/JohannesKaufmann/html-to-markdown"
	ButterCMS "github.com/buttercms/buttercms-go"
)

func ensureDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		print(err)
		err := os.MkdirAll(dirName, os.ModePerm)
		if err == nil || os.IsExist(err) {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func writeToFile(post ButterCMS.Post, md string) {
	path := fmt.Sprintf("./content/posts/%s.md", post.Slug)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(md)
	}
	file.Close()
}

func main() {

	if err := ensureDir("./content/posts/"); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}

	ButterCMS.SetAuthToken("YOUR_AUTH_TOKEN")

	page := 1

	for {
		params := map[string]string{"page": strconv.Itoa(page)}

		postResp, err := ButterCMS.GetPosts(params)
		if err != nil {
			break
		}

		page += 1

		if len(postResp.PostList) == 0 {
			break
		}

		for index, post := range postResp.PostList {
			fmt.Printf("%d %s\n", index, post.Title)
			converter := md.NewConverter("", true, nil)
			markdown, err := converter.ConvertString(post.Body)
			if err != nil {
				continue
			}

			fmt.Printf(post.FeaturedImage)
			if post.FeaturedImage != "" {
				mdOutput := fmt.Sprintf("---\ntitle: \"%s\"\ndate: %s\ndraft: false\nfeatured_image: %s\n---\n%s", post.Title, post.Published, post.FeaturedImage, markdown)
				writeToFile(post, mdOutput)

			} else {
				mdOutput := fmt.Sprintf("---\ntitle: \"%s\"\ndate: %s\ndraft: false\n---\n%s", post.Title, post.Published, markdown)
				writeToFile(post, mdOutput)
			}

		}
	}

}
