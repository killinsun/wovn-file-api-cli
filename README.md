# WOVN File translation API Tool

## Description

This tools is a wrapper tool for WOVN File Translation API.

This tool has following features.

- Make an archive file which contains translated HTML files keeping directory structure by URL list file.

  - This is a kind of download translated html file.

- Send a report WOVN using Upload API by URL list file. (not avairable now.)
  - This is a kind of upload original html file.

## Supported api versions

This supports **WOVN File translation API v1**.

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
- `apiKey` ... You can generate it at the WOVN dashboard. In order to generate API token, you need to get API feature from WOVN marketplace.
- `targetLangCode` ... The languages that you want to translate. It's also needed to inlucde source language.
- `srcLangCode` ... The language that your website's original one.

Other values are dummy parameters, so you can't use them for now.

```bash
$ cp settings_template.json settings.json
```

`url.txt` is a simple line-breaked URL list file like below.

```url.txt
https://example.com
https://example.com/page1.html
https://example.com/page2.html
```

### Run

You can download translated HTML files by `-d` option.

```bash
$ ./wovn-file-api-cli -d
```

then, you can see `WovnTranslatedHtml` directory in your current directory.
Now, you can upload HTML files on your own server.
