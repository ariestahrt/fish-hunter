package twitter

import (
	"bytes"
	"encoding/json"
	"fish-hunter/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/dghubble/oauth1"
)

type Media struct {
	MediaIDS []string `json:"media_ids"`
}

type TweetData struct {
	Text  string `json:"text"`
	Media Media  `json:"media,omitempty"`
}

func UploadMedia(mediaPath string) (map[string]interface{}, error) {
	endpoint := "https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image"

	config := oauth1.NewConfig(util.GetConfig("TWITTER_CONSUMER_KEY"), util.GetConfig("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(util.GetConfig("TWITTER_ACCESS_TOKEN"), util.GetConfig("TWITTER_SECRET_TOKEN"))
	httpClient := config.Client(oauth1.NoContext, token)

	// create body form
	payload := &bytes.Buffer{}
	form := multipart.NewWriter(payload)

	// create media paramater
	fw, err := form.CreateFormFile("media", filepath.Base(mediaPath))
	if err != nil {
		return nil, err
	}

	// open file
	opened, err := os.Open(mediaPath)
	if err != nil {
		return nil, err
	}
	defer opened.Close()

	// copy to form
	_, err = io.Copy(fw, opened)
	if err != nil {
		return nil, err
	}

	// close form
	form.Close()

	resp, err := httpClient.Post(endpoint, form.FormDataContentType(), payload)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Convert body to map json
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	return result, nil
}

func Tweet(t TweetData) error {
	endpoint := "https://api.twitter.com/2/tweets"

	config := oauth1.NewConfig(util.GetConfig("TWITTER_CONSUMER_KEY"), util.GetConfig("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(util.GetConfig("TWITTER_ACCESS_TOKEN"), util.GetConfig("TWITTER_SECRET_TOKEN"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Convert tweetdata to json
	payloadb, _ := json.Marshal(t)
	payloads := string(payloadb)
	payloadIO := strings.NewReader(payloads)

	req, err := httpClient.Post(endpoint, "application/json", payloadIO)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	return nil
}