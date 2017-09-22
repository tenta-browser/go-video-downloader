package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type XNXXIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewXNXXIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &XNXXIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:video|www)\.xnxx\.com/video-?(?P<id>[0-9a-z]+)/`
	return ret
}

func (self *XNXXIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *XNXXIE) Key() string {
	return "XNXX"
}

func (self *XNXXIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *XNXXIE) Name() string {
	return `XNXX extractor`
}

func (self *XNXXIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *XNXXIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := rnt.SearchRegex(self, `flv_url=(.*?)&amp;`, webpage, `video URL`, rnt.NoDefault, true, 0, nil)
	video_url = rnt.AsOptString(rnt.ParseUnquote(video_url.Get()))
	video_title := rnt.HTMLSearchRegex(self, `<title>(.*?)\s+-\s+XNXX.COM`, webpage, `title`, rnt.NoDefault, true, 0, nil)
	video_thumbnail := rnt.SearchRegex(self, `url_bigthumb=(.*?)&amp;`, webpage, `thumbnail`, rnt.NoDefault, false, 0, nil)
	return map[string]interface{}{`id`: video_id,
		`url`:       video_url,
		`title`:     video_title,
		`ext`:       `flv`,
		`thumbnail`: video_thumbnail,
		`age_limit`: 18}
}
func (self *XNXXIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`XNXX`, NewXNXXIE)
}
