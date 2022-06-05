package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"regexp"
	"strconv"
	"time"
)

type CompactMs struct {
	// Id primary key, 主键
	Id int64
	// Title title in English, 英文标题
	Title string
	// Regions available regions, 可观看地区
	Regions []string
	// Subtitles languages of subtitles, 字幕语言
	SubtitlesSubtitles []string
	// Audios languages of audios, 音频语言
	Audios []string
}

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	var body string
	url := "https://www.netflix.com/watch/80162147?trackId=255875003"
	err := chromedp.Run(ctx, Fetch(&body, url, "memclid", "834f9567-31ce-4d1b-b6c8-2d13c77d36b6", "netflix-sans-normal-3-loaded", "true", "netflix-sans-bold-3-loaded", "true", "nfvdid", "BQFmAAEBEFpGlQ8mO56ZXeVYlhy-9JpgqWrCSW2GpVi1tkzG318-WHijqCOGwQ3AjAOdis8S3j7-Lcpa8JjdWN8YaKgnAGB2IzpKntWj4CT-em82HAoyrxEs4EFHhaICU_0AQ7YHIBVd4BwIJCJc816aj7GqRm7J", "flwssn", "97401bbb-4e2e-4bc0-b3e1-7869037dca2d", "SecureNetflixId", "v%3D2%26mac%3DAQEAEQABABQtbkO_7AsjHQWH-CefHoUtKQftcqHJ_LI.%26dt%3D1654325530854", "NetflixId", "ct%3DBQAOAAEBEFhPg3BfNaJRD1MwyQkHhSGB4JOEumwoAEWHOjNE2lJm-faCu7YOM8x2zHpxoEFFEU3EyYYZ5PZ4LhYa_qDXL44DBxsdjQdUX5IyfKhNoFzhf41FyGDXRle0naxDpTMKZvEz6fPf4_iefZcDDjK24LB90dgbY-fHeEQ_tbm37i4Qfz2brILlv7_ngps8vfognuMqObnJl3BGNoo37_-vuCz9iG1KUAkChmdpAD4E2lTVCBwvRFiYL5qJtO_x_KGGShY-3TBUIEXxv2k-CQ6PsL6f3jzeEJ3cAWS9GZ4m-qRT-somyEeOCCkDSfSZ-bcnfogfKJXd3NsuMJiEuzQNQlyOZ1MZyWpXhCk-_jF5BTRSYlgSPCD-9TcoE4hG0B1AY2GUi8WCClrOu8G56RV1ISmNu4rjUuEIejcHv2w4IbAeCg4sP8zCrla3D9z9JYYXSyPweISFA6APOthhW5t4XC4PsgFX3Ur8Ugw1n74MzvoNiZZuPIG3z6Vh5lB9TwXRGLN8TV9olqUrzHn66d4xji64bxkjRG3uiqLN-ZM48p1Ilf-WzJFEMtel8v5nR6Ii3oJ6vevEbUPsvja9-7H_AqHnHcvSomoE-AEt8e9BbJNmRIHP90csq6fRID8pvXNsoxJ4Rty8otLfN-2r_qwunv0nfA..%26bt%3Ddbl%26ch%3DAQEAEAABABSXfjcYREZ_2a7xHkkMbCuj0i6rZVgU2c8.%26v%3D2%26mac%3DAQEAEAABABQBeYdQ147RDc80nZ5I641rzeRIRX2wNZU.", "profilesNewSession", "0", "OptanonConsent", "=isIABGlobal=false&datestamp=Sat+Jun+04+2022+16%3A09%3A48+GMT%2B0800+(%E4%B8%AD%E5%9B%BD%E6%A0%87%E5%87%86%E6%97%B6%E9%97%B4)&version=6.6.0&consentId=f4030888-327b-4a22-b11d-815486235f4d&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1&hosts=H12%3A1%2CH13%3A1%2CH51%3A1%2CH45%3A1%2CH46%3A1%2CH48%3A1%2CH49%3A1&AwaitingReconsent=false"))
	CheckErrQuitIfNotNil(err)
	rs, err := ParseDetail(body)
	CheckErrQuitIfNotNil(err)
	fmt.Println(rs)

}

// ParseList parses a list of movie subject metadata
func ParseList(body []byte) (rs []*CompactMs, err error) {
	panic("unimplemented")
}

// ParseDetail parses one movie subject metadata
func ParseDetail(body string) (rs *CompactMs, err error) {
	//正则匹配目标
	id := regexp.MustCompile(`videoId":(\w.*?),`)
	title := regexp.MustCompile(`title">(\w.*?)</h1><div`)
	region := regexp.MustCompile(`locale":("\w.*?),`)
	audio := regexp.MustCompile(`audio">(\w.*?)<!-- -->`)
	subtitle := regexp.MustCompile(`subtitle">(\w.*?)<!-- -->`)
	rs = new(CompactMs)
	Id, _ := strconv.Atoi(id.FindStringSubmatch(body)[1])
	rs.Id = int64(Id)
	rs.Title = title.FindStringSubmatch(body)[1]
	audios := audio.FindAllStringSubmatch(body, -1)
	subtitles := subtitle.FindAllStringSubmatch(body, -1)
	regions := region.FindAllStringSubmatch(body, -1)
	for i, _ := range regions {
		rs.Regions = append(rs.Regions, regions[i][1])
	}
	for i := 0; i < len(subtitles); i++ {
		rs.SubtitlesSubtitles = append(rs.SubtitlesSubtitles, subtitles[i][1])
	}
	for i := 0; i < len(audios)-1; i++ {
		rs.Audios = append(rs.Audios, audios[i][1])
	}
	return rs, nil

}

// Fetch requests HTTP body by API UR
func Fetch(body *string, fullUrl string, cookies ...string) chromedp.Tasks {
	//创建一个chrome任务
	return chromedp.Tasks{
		//ActionFunc是一个适配器，允许使用普通函数作为操作。
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 设置Cookie存活时间
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// 添加Cookie到chrome
			for i := 0; i < len(cookies); i += 2 {
				//SetCookie使用给定的cookie数据设置一个cookie； 如果存在，可能会覆盖等效的cookie。
				network.SetCookie(cookies[i], cookies[i+1]).
					// 设置cookie到期时间
					WithExpires(&expr).
					// 设置httponly,防止XSS攻击
					WithHTTPOnly(true).
					//Do根据提供的上下文执行Network.setCookie。
					Do(ctx)

			}
			return nil
		}),
		// 跳转指定的url地址
		chromedp.Navigate(fullUrl),
		chromedp.OuterHTML(`document.querySelector("body")`, body, chromedp.ByJSPath),
	}
}

func CheckErrQuitIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
