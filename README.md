Go Video Downloader
===================

[![GoDoc](https://godoc.org/github.com/tenta-browser/go-video-downloader?status.svg)](https://godoc.org/github.com/tenta-browser/go-video-downloader)

Extracts video info and download URL from various sites using extractors transpiled from [ytdl-org/youtube-dl](https://github.com/ytdl-org/youtube-dl).

Contact: developer@tenta.io

Installation
============

1. `go get github.com/tenta-browser/go-video-downloader`

Usage
=====

```go
matcher.ReEngine = matcherpcre.NewEngine() // select matcher engine

// The downloader has to be initialized once, before usage.
// This is a slow, computationally intensive step.
if downloader.Init() != nil {
    // failed to initialize downloader
    return
}

// Check is a quick method to verify if an URL holds a downloadable video.
found, err := downloader.Check("https://example.com/video")
if err == nil {
    // internal error occurred while matching the specified URL
    return
}

connector := &downloader.Connector{
    Client:    &http.Client{},
    UserAgent: "Mozilla/5.0 (X11; Linux x86_64)..",
    Cookie:    "SESSIONID=12345",
}

videoData, err := downloader.Extract("https://example.com/video", connector)
if err != nil {
    // failed to extract video data
}

// success, videoData holds download URL, suggested file name, video title, etc.
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