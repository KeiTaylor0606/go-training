package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gen2brain/beeep"
)

//全ての構造体の共通インターフェース
type Warning interface {
	Show(message string)
}

//コンソールにメッセージを通知する構造体
type ConsoleWarning struct{}

func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]:%s\n", os.Args[0], message)
}

//デスクトップにメッセージを通知する構造体
type DesktopWarning struct{}

func (d DesktopWarning) Show(message string) {
	beeep.Alert(os.Args[0], message, "")
}

//Slackにメッセージを通知する構造体
type SlackWarning struct {
	URL     string
	Channel string
}

type slackMessage struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func (s SlackWarning) Show(message string) {
	params, _ := json.Marshal(slackMessage{
		Text:      message,
		Username:  os.Args[0],
		IconEmoji: ":robot_face:",
		Channel:   s.Channel,
	})

	resp, err := http.PostForm(
		s.URL,
		url.Values{"payload": {string(params)}},
	)
	if err == nil {
		io.ReadAll(resp.Body)
		defer resp.Body.Close()
	}
}

func main() {
	//Show()メソッドを持つインスタンスはなんでも入れられる変数
	var warn Warning

	//コンソールに出力
	warn = &ConsoleWarning{}
	warn.Show("Hello World to console")

	//デスクトップに出力
	warn = &DesktopWarning{}
	warn.Show("Hello World to desktop")

	//Slackに出力
	warn = &SlackWarning{
		URL:     os.Getenv("SLACK_URL"),
		Channel: "#general",
	}
	warn.Show("Hello World to Slack")
}
