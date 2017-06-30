# promsortrules

Sort and reformat prometheus rules files alphabetically


## Installation and usage

```
go get github.com/laher/promsortrules
promsortrules alerts.rules 
```

## Formatting

I use this tool to easily compare different prometheus rules files.

 * You could use it to format your files in a consistent way: `promsortrules alerts.rules | tee tmp.rules && mv tmp.rules alerts.rules`
 * YMMV - I haven't thoroughly tested that the output, but it uses prometheus' parser completely.
 * Note: promsortrules reformats your rules in ways you might not love. It's a bit squashed, but it's consistent, which makes me happy.


## Thanks

This work is completely derived from promtool (it's a based on `promtool check-rules`) from the prometheus project, and it depends on that project directly. I'll fork and PR to promtool, but for now I just want something I can use.
