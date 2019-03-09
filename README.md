# ping_meta_tags

Collect metatag attribute values (property value, name value, content value) from given url.

## Installation

Use `go get` to install this package:

```bash
$ go get github.com/mkusaka/ping_meta_tags
$ export url=`url1,url2,...` ping_meta_tags
```

## useage
ping_meta_tag will make result.csv under `tmp`.

CSV format is below.

|url|property value|name value|content value|timestamp|
| --- | --- | --- | --- | --- |
|url1|prop1|name1|content1|unix timestamp1|
