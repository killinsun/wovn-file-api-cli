# WOVN File translation API Tools

## Description

This tools is a wrapper tool for the WOVN File Translation API.

This can these features like,

- Make an archive file which contains translated HTML files with keeping directory structure by URL list file.

  - This is a kind of download translated html file.

- Send a report WOVN using Upload API by URL list file. (not avairable now.)
  - This is a kind of upload original html file.

## Supported api versions

This tool uses **WOVN File translation API v1**.

## How to use

### Build

```bash
$ git clone https://github.com/killinsun/wovn-file-api-cli.git
$ go build
```

`wovn-file-api-cli` binary will be put on your directory.

### Make required files.

You need to prepare `url.txt` and `settings.json` in advance.

`settings.json` has required parameters.

- `projectToken` ... You can check it at the WOVN dashboard.
- `apiKey` ... You can generate it at the WOVN dashboard.
- `targetLangCode` ... Languages which you want to translate. but you have to add source language code.
- `srcLangCode` ... Language wich your website's original one.

and other value is a fixed option. so you cannot customize for now.

```bash
$ cp settings_template.json settings.json
```

`url.txt` is a simple line-breaked URL list file like below.

```url.txt
https://example.com
https://example.com/page1.html
https://example.com/page2.html
```

We recommend you should put these files on same directory with `wovn-file-api-cli` binary.

### Run

You can download HTML files by `-d` option.

```bash
$ ./wovn-file-api-cli -d
```

then, you can see `./WovnTranslatedHtml` in your current directory.
Now, you can upload HTML files on your server.
