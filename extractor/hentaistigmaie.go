package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HentaiStigmaIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewHentaiStigmaIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &HentaiStigmaIE{}
	ret.Context = ctx
	ret._VALID_URL = `^https?://hentai\.animestigma\.com/(?P<id>[^/]+)`
	return ret
}

func (self *HentaiStigmaIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *HentaiStigmaIE) Key() string {
	return "HentaiStigma"
}

func (self *HentaiStigmaIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *HentaiStigmaIE) Name() string {
	return `HentaiStigma extractor`
}

func (self *HentaiStigmaIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *HentaiStigmaIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	webpage := rnt.DownloadWebpage(self, url, video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	title := rnt.HTMLSearchRegex(self, `<h2[^>]+class="posttitle"[^>]*><a[^>]*>([^<]+)</a>`, webpage, `title`, rnt.NoDefault, true, 0, nil)
	wrap_url := rnt.HTMLSearchRegex(self, `<iframe[^>]+src="([^"]+mp4)"`, webpage, `wrapper url`, rnt.NoDefault, true, 0, nil)
	wrap_webpage := rnt.DownloadWebpage(self, wrap_url.Get(), video_id, rnt.OptString{}, rnt.OptString{}, true, 1, 5, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	video_url := rnt.HTMLSearchRegex(self, `file\s*:\s*"([^"]+)"`, wrap_webpage, `video url`, rnt.NoDefault, true, 0, nil)
	return map[string]interface{}{`id`: video_id,
		`url`:       video_url,
		`title`:     title,
		`age_limit`: 18}
}
func (self *HentaiStigmaIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`HentaiStigma`, NewHentaiStigmaIE)
}
