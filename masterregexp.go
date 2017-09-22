package downloader

func init() {
	masterRegexp = `(?<extr_Bild>https?://(?:www\.)?bild\.de/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?<extr_Criterion>https?://(?:www\.)?criterion\.com/films/(?:[0-9]+)-.+)|(?<extr_EchoMsk>https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?<extr_FiveTV>(?x)
                    http://
                        (?:www\.)?5-tv\.ru/
                        (?:
                            (?:[^/]+/)+(?:\d+)|
                            (?:[^/?#]+)(?:[/?#])?
                        )
                    )|(?<extr_Hark>https?://(?:www\.)?hark\.com/clips/(?:.+?)-.+)|(?<extr_HentaiStigma>^https?://hentai\.animestigma\.com/(?:[^/]+))|(?<extr_Keek>https?://(?:www\.)?keek\.com/keek/(?:\w+))|(?<extr_Morningstar>https?://(?:www\.)?morningstar\.com/[cC]over/video[cC]enter\.aspx\?id=(?:[0-9]+))|(?<extr_Slutload>^https?://(?:\w+\.)?slutload\.com/video/[^/]+/(?:[^/]+)/?$)|(?<extr_WDRMobile>(?x)
        https?://mobile-ondemand\.wdr\.de/
        .*?/fsk(?:[0-9]+)
        /[0-9]+/[0-9]+/
        (?:[0-9]+)_(?:[0-9]+))|(?<extr_XNXX>https?://(?:video|www)\.xnxx\.com/video-?(?:[0-9a-z]+)/)`
}
