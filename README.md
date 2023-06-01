# md-to-pdf

## install pandoc and setup


install pandoc
```
sudo apt-get install pandoc
```

install texlive
```
sudo apt-get install texlive-xetex
```

get pandoc home `is needed for here: -pandoc-path`
```
sudo find / -type d -name "pandoc"
```

## Run

to run
```
go build -o converter *.go  && sudo ./converter -host 0.0.0.0:8080 -pandoc-path /usr/share/pandoc -home /home/aidan
```


## Requests

use need to convert a MD file to 
https://base64.guru/converter/encode/file


### to send as base64

- end point: `/convert`
- POST

- `input:` base64 of md file
- `write_to_home_dir:` if true will save the file to your home dir as well as /tmp

body
```json
{
  "input": "IyBSRUFETUUKCiMjIyBQcmUtcmVxdWlzaXRlcwoKLSBjaGVjayB3aGF0IHZlcnNpb24gb2YgZ28geW91IGhhdmUgaW5zdGFsbGVkIGdvIHZlcnNpb24KLSBpZiBpdHMgPCAxLjE4IHRoZSBkb3dubG9hZCBhbmQgaW5zdGFsbCBnbyAxLjE4IGZvciB5b3VyIE9TIGh0dHBzOi8vZ28uZGV2L2RsLwotIGluc3RhbGwvdXBncmFkZSB3YWlscyB5b3UgbmVlZCB0byByZWluc3RhbGwgaHR0cHM6Ly93YWlscy5pby9kb2NzL2dldHRpbmdzdGFydGVkL2luc3RhbGxhdGlvbgotIGluc3RhbGwgbm9kZSA+PSB2MTQKCiMjIyBJbnN0YWxsIG5wbSBwYWNrYWdlcwoKYGBgCmNkIGZyb250ZW5kCm5wbSBpbnN0YWxsCmBgYAoKIyMjIEhvdyB0byBzdGFydAoKCiMjIyMgSWYgaXQncyBmaXJzdCBydW4gKHRoaXMgaXMganVzdCBuZWVkZWQgYXQgdGhlIGZpcnN0IHRpbWUgZm9yIGdlbmVyYXRpbmcgZnJvbnRlbmQvZGlzdCBiZWZvcmUgc3RhcnRpbmcgYW55dGhpbmcpCmBgYAp3YWlscyBidWlsZApgYGAKCiMjIyMgRG93bmxvYWQgcHJpdmF0ZSByZXBvIGRlcGVuZGVuY2llcyBvbiBMaW51eApgYGAKZXhwb3J0IEdJVEhVQl9UT0tFTj08WU9VUl9HSVRIVUJfVE9LRU4+CmdpdCBjb25maWcgLS1nbG9iYWwgdXJsLiJodHRwczovLyRHSVRIVUJfVE9LRU46eC1vYXV0aC1iYXNpY0BnaXRodWIuY29tL051YmVJTy9ydWJpeC1hc3Npc3QiLmluc3RlYWRPZiAiaHR0cHM6Ly9naXRodWIuY29tL051YmVJTy9ydWJpeC1hc3Npc3QiCmdpdCBjb25maWcgLS1nbG9iYWwgdXJsLiJodHRwczovLyRHSVRIVUJfVE9LRU46eC1vYXV0aC1iYXNpY0BnaXRodWIuY29tL051YmVJTy9ydWJpeC1lZGdlIi5pbnN0ZWFkT2YgImh0dHBzOi8vZ2l0aHViLmNvbS9OdWJlSU8vcnViaXgtZWRnZSIKYGBgCgojIyMjIERvd25sb2FkIHByaXZhdGUgcmVwbyBkZXBlbmRlbmNpZXMgb24gbWFjCmBgYApleHBvcnQgR0lUSFVCX1RPS0VOPTxZT1VSX0dJVEhVQl9UT0tFTj4KZXhwb3J0IEdPUFJJVkFURT1naXRodWIuY29tL051YmVJTy9ydWJpeC1hc3Npc3QKZ28gZ2V0IC12CmV4cG9ydCBHT1BSSVZBVEU9Z2l0aHViLmNvbS9OdWJlSU8vcnViaXgtZWRnZQpnbyBnZXQgLXYKYGBgCgojIyMjIFN0YXJ0IGFwcAoKYGBgCndhaWxzIGRldgpgYGAKCiMjIyBGcm9udGVuZCBicm93c2VyIFVSTAoKYGBgCmh0dHA6Ly9sb2NhbGhvc3Q6MzQxMTUvCmBgYAoKCiMgYXBwIGljb24KCnBhc3RlIHRoZSBmaWxlIGBhcHBpY29uLnBuZ2AgaW50byBgZGlyYCBgYnVpbGRgCgpodHRwczovL3dhaWxzLmlvL2RvY3MvcmVmZXJlbmNlL29wdGlvbnMv",
  "write_to_home_dir":false
}
```

### to read a local file
- end point: `/convert/local`
- POST

body
```json
{
  "file": "/home/aidan/code/go/nube/md-to-pdf/README.md",
  "write_to_home_dir":true
}
```

found an example using templates

https://github.com/gogap/go-pandoc


## charts

```
github.com/vicanso/go-charts/v2
```
![test image](test.png)
