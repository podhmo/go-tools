```console
$ go run main.go
## load once
52 packages 682.644912ms
52 packages 648.572636ms
52 packages 663.843659ms
## load twice (without cache)
52 packages 1.447409222s
52 packages 1.559340827s
52 packages 1.206504886s
## load twice (with cache)
52 packages 728.334784ms
52 packages 800.806509ms
52 packages 675.893973ms
```
