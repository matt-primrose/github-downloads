# github-downloads
A small app for getting the number of times release assets have been downloaded from github.  Works for any public repository on github.  Returns the most recent 30 releases.  The summarized total downloads per release asset are displayed on the command line output.  The raw JSON output from github is written to a results.json file.

## compile
```bash
go build -o github-downloads.exe .\cmd\main.go
```

## run
```bash
github-downloads.exe <OWNER> <REPO>
```