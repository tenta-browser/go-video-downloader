package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type EchoMskIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewEchoMskIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &EchoMskIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.)?echo\.msk\.ru/sounds/(?P<id>\d+)`
	return ret
}

func (self *EchoMskIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *EchoMskIE) Key() string {
	return "EchoMsk"
}

func (self *EchoMskIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *EchoMskIE) Name() string {
	return `EchoMsk extractor`
}

func (self *EchoMskIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *EchoMskIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	audio_url := rnt.SearchRegex(self, `<a rel="mp3" href="([^"]+)">`, webpage, `audio URL`, rnt.NoDefault, true, 0, nil)
	title := rnt.HTMLSearchRegex(self, `<a href="/programs/[^"]+" target="_blank">([^<]+)</a>`, webpage, `title`, rnt.NoDefault, true, 0, nil)
	air_date := rnt.HTMLSearchRegex(self, `(?s)<div class="date">(.+?)</div>`, webpage, `date`, nil, false, 0, nil)
	if (air_date).IsSet() && (air_date.Get()) != "" {
		air_date = rnt.AsOptString(rnt.ReSub(rnt.Re, `(\s)\1+`, `\1`, air_date.Get(), 0, 0))
		if (air_date).IsSet() && (air_date.Get()) != "" {
			title = rnt.AsOptString(rnt.StrFormat(`%s - %s`, title, air_date))
		}
	}
	return map[string]interface{}{`id`: video_id,
		`url`:   audio_url,
		`title`: title}
}
func (self *EchoMskIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`EchoMsk`, NewEchoMskIE)
}
