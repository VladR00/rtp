# Description:

Generate sequence and based on it generate multipliers such that the formula RTP:=sum(transformed)/count(sequence) corresponds to the entered RTP.
If multiplier > sequence: transformed = sequence; otherwise 0

# How to start

```bash
go run ./cmd/main.go -rtp=1
```

# How to generate 10000 sequence and multipliers:

```bash
curl http://localhost:8080/get/10000
```
