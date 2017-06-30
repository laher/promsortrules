# promrulesort

Sort and reformat prometheus rules files alphabetically


## Installation and usage

```
go get github.com/laher/promrulesort
promrulesort alerts.rules | tee tmp.rules && mv tmp.rules alerts.rules
```

## Formatting

I use this tool to easily compare different prometheus rules files. 

 * Note: promrulesort reformats your rules in ways you might not love. It's a bit squashed, but it's consistent, which I'm happy about.


## Thanks

This is completely derived from promtool (it's a based on `promtool check-rules`) from the prometheus project. I'll fork and PR to that project, but for now I just want something I can use.

## Vendoring

I'm holding off for `dep`, thanks.
