# xkcd

Exercise:

```
The popular web comic xkcd has a JSON interface. For example, a request to
https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many
favorites. Download each URL (once!) and build an offline index. Write a tool xkcd that, using this
index, prints the URL and transcript of each comic that matches a search term provided on the
command line.
```

To run:

```bash
# to build the index
go run ./cmd/xkcd_fetch

# or

# to find all comics with the search term "banana"
go run ./cmd/xkcd "banana"
```
