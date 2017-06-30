# promsortrules

Sort prometheus rules files alphabetically by rule name. 
promsortrules reformats the rule definitions in a consistent manner.

## Installation and usage

```
go get github.com/laher/promsortrules
promsortrules alerts.rules 
```

## Formatting

I use this tool to compare prometheus rules files more easily.

 * You _could_ use its output to format your files in a consistent way: `promsortrules alerts.rules | tee tmp.rules && mv tmp.rules alerts.rules`
 * YMMV - I haven't thoroughly tested that the output works, but it is simply using prometheus' parser, so it should be OK. Just test it before using it in production.
 * Note: promsortrules reformats your rules in ways you might not love. It's a bit squashed, but importantly it's consistent, which does make me happy.

## Thanks

This work is completely derived from promtool (it's a based on `promtool check-rules`) from the prometheus project, and it depends on that project directly. I'll fork and PR to promtool, but for now I just want something I can use.
