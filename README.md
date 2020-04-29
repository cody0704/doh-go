# DoH Go

* Version: 1.14

## DoH Browser Setting

### Desktop

* Firefox
    * 75.0

* Chrome
    * 81.0.4044.122

### Mobile
* Firefox
    * 68.7

* Chrome
    * 81.0.4044.111

### Ref

1. [Mozilla Wiki - Trusted Recursive Resolver](https://wiki.mozilla.org/Trusted_Recursive_Resolver#network.trr.mode)

## Test

### GET

```bash
curl -H 'accept: application/dns-message' -v 'https://{DoH IP}/dns-query?dns=q80BAAABAAAAAAAAA3d3dwdleGFtcGxlA2NvbQAAAQAB' | hexdump

0000000 ab cd 81 80 00 01 00 01 00 00 00 00 03 77 77 77
0000010 07 65 78 61 6d 70 6c 65 03 63 6f 6d 00 00 01 00
0000020 01 03 77 77 77 07 65 78 61 6d 70 6c 65 03 63 6f
0000030 6d 00 00 01 00 01 00 00 50 78 00 04 5d b8 d8 22
0000040
```

### POST

```bash
go run tools/doh_client_post.go

Method:POST. OK!
```