Go Video Downloader
===================

[![GoDoc](https://godoc.org/github.com/tenta-browser/go-video-downloader?status.svg)](https://godoc.org/github.com/tenta-browser/go-video-downloader)

Extracts video info and download URL from various sites using extractors transpiled from [rg3/youtube-dl](https://github.com/rg3/youtube-dl).

Contact: developer@tenta.io

Installation
============

1. `go get github.com/tenta-browser/go-video-downloader`

Usage
=====

```go
matcher.ReEngine = matcherpcre.NewEngine() // select matcher engine

checkResult := downloader.Check("http://somedomain.com/favourite-cat-video")
if checkResult == nil {
    // no suitable extractors found for the specified URL
    return
}

connector := &downloader.Connector{
    Client:    &http.Client{},
    UserAgent: "Mozilla/5.0 (X11; Linux x86_64)..",
}

videoData, err := downloader.Extract(checkResult, connector)
if err != nil {
    // failed to extract video data
}

// success
```

Supported sites
===============

[List of supported sites](SUPPORTEDSITES.md)

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

This regex library is brought to you by Team Tenta. Tenta is your [private, encrypted browser](https://tenta.com) that protects your data instead of selling it. We're building a next-generation browser that combines all the privacy tools you need, including built-in OpenVPN. Everything is encrypted by default. That means your bookmarks, saved tabs, web history, web traffic, downloaded files, IP address and DNS. A truly incognito browser that's fast and easy.