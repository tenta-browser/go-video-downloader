package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type SlutloadIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewSlutloadIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &SlutloadIE{}
	ret.Context = ctx
	ret._VALID_URL = `^https?://(?:\w+\.)?slutload\.com/video/[^/]+/(?P<id>[^/]+)/?$`
	return ret
}

func (self *SlutloadIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *SlutloadIE) Key() string {
	return "Slutload"
}

func (self *SlutloadIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *SlutloadIE) Name() string {
	return `Slutload extractor`
}

func (self *SlutloadIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *SlutloadIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_title := rnt.StrStrip(rnt.HTMLSearchRegex(self, `<h1><strong>([^<]+)</strong>`, webpage, `title`, rnt.NoDefault, true, 0, nil).Get(), ``)
	video_url := rnt.HTMLSearchRegex(self, `(?s)<div id="vidPlayer"\s+data-url="([^"]+)"`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	thumbnail := rnt.HTMLSearchRegex(self, `(?s)<div id="vidPlayer"\s+.*?previewer-file="([^"]+)"`, webpage, `thumbnail`, rnt.NoDefault, false, 0, nil)
	return map[string]interface{}{`id`: video_id,
		`url`:       video_url,
		`title`:     video_title,
		`thumbnail`: thumbnail,
		`age_limit`: 18}
}
func (self *SlutloadIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`Slutload`, NewSlutloadIE)
}
