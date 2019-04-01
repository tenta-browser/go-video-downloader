Go Video Downloader
===================

[![GoDoc](https://godoc.org/github.com/tenta-browser/go-video-downloader?status.svg)](https://godoc.org/github.com/tenta-browser/go-video-downloader)

Extracts video info and download URL from various sites using extractors transpiled from [ytdl-org/youtube-dl](https://github.com/ytdl-org/youtube-dl).

Contact: developer@tenta.io

Installation
============

```
go get github.com/tenta-browser/go-video-downloader
```

The extractor leans heavily on a regular expression engine supporting Perl Compatible Regular Expressions (PCRE). This library explicitly doesn't reference one, but uses the interface provided by [tenta-browser/go-pcre-matcher](https://github.com/tenta-browser/go-pcre-matcher). An implementation is provided there, to use that:
 - install libpcre, on Debian based platforms: `sudo apt-get install libpcre++-dev`
 - make sure the `github.com/tenta-browser/go-pcre-matcher/matcherpcre` package is fetched
 - reference it in `matcher.ReEngine` (see the example below) 

Usage
=====

```go
package main

import (
    "net/http"

    matcher "github.com/tenta-browser/go-pcre-matcher"
    "github.com/tenta-browser/go-pcre-matcher/matcherpcre"
    downloader "github.com/tenta-browser/go-video-downloader"
)

func main() {
    // Select PCRE matcher engine.
    matcher.ReEngine = matcherpcre.NewEngine()

    // The downloader has to be initialized once, before usage.
    if err := downloader.Init(); err != nil {
        panic(err.Error())
    }

    // Quick check to see if the specified URL contains a video
    // (for which extraction is supported).
    hasVideo, err := downloader.Check("https://tenta.com/how-to-download-videos")
    if err != nil {
        panic(err.Error())
    }
    println("Has video:", hasVideo)

    // Pluggable HTTP client and common headers to be used at each page request
    // done by the extractor.
    connector := &downloader.Connector{
        Client:    &http.Client{},
        UserAgent: "Mozilla/5.0 (X11; Linux x86_64) ..",
        Cookie:    "SESSIONID=12345",
    }

    // Extracts various properties of the video contained at the specified URL,
    // most importantly it's direct URL.
    // This is a slow call, the extractor will possibly do multiple HTTP requests.
    videoData, err := downloader.Extract("https://tenta.com/how-to-download-videos", connector)
    if err != nil {
        panic(err.Error())
    }
    println("Video title:", videoData.Title)
    println("Raw video URL:", videoData.URL)
    println("Suggested filename:", videoData.Filename)
}
```

Supported sites
===============

[List of supported sites](SUPPORTEDSITES.md)

Notices
=======

We rely on excellent open source libraries. 
For a complete list of our dependencies and required notification, please take a look at [NOTICES](NOTICES.md).

License
=======

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

For any questions, please contact developer@tenta.io

Contributing
============

We welcome contributions, feedback and plain old complaining. Feel free to open
an issue or shoot us a message to developer@tenta.io. If you'd like to contribute,
please open a pull request and send us an email to sign a contributor agreement.

About Tenta
===========

This library is brought to you by Team Tenta. Tenta is your [private, encrypted browser](https://tenta.com) that protects your data instead of selling it. We're building a next-generation browser that combines all the privacy tools you need, including built-in OpenVPN. Everything is encrypted by default. That means your bookmarks, saved tabs, web history, web traffic, downloaded files, IP address and DNS. A truly incognito browser that's fast and easy.