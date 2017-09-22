package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type HarkIE struct {
	*rnt.Context
	_VALID_URL string
}

func NewHarkIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &HarkIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.)?hark\.com/clips/(?P<id>.+?)-.+`
	return ret
}

func (self *HarkIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *HarkIE) Key() string {
	return "Hark"
}

func (self *HarkIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *HarkIE) Name() string {
	return `Hark extractor`
}

func (self *HarkIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *HarkIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	data := rnt.DownloadJSON(self, rnt.StrFormat(`http://www.hark.com/clips/%s.json`, video_id), video_id, rnt.AsOptString(`Downloading JSON metadata`), rnt.AsOptString(`Unable to download JSON metadata`), 0, true, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	return map[string]interface{}{`id`: video_id,
		`url`:         (data)[`url`],
		`title`:       (data)[`name`],
		`description`: rnt.DictGet(data, `description`, nil),
		`thumbnail`:   rnt.DictGet(data, `image_original`, nil),
		`duration`:    rnt.DictGet(data, `duration`, nil)}
}
func (self *HarkIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`Hark`, NewHarkIE)
}
