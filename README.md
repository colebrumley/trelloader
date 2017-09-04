# `trelloader`: A quick and dirty utility to create a Trello board

This utility creates a Trello board pre-populated with lists, cards, and labels from a JSON template. Since this was initially created to quickly throw up a board for AWS Well-Architected Framework assessments, see `examples/waf.json`

To get an AppKey and token, go [here](https://trello.com/app-key).

## Install

To install via go (make sure `$GOPATH/bin` is in your `PATH`), run

```
go get github.com/colebrumley/trelloader
```

To install via this repo, run

```
git clone https://github.com/colebrumley/trelloader.git
cd trelloader/
go build
mv trelloader /usr/local/bin/
```
