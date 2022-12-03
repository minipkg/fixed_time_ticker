# fixed_time_ticker

## Description

Simple ticker at a fixed time.

### Installation

Run the following command to install the package:

```
go get github.com/minipkg/fixed_time_ticker
```

## Basic Usage

```go
import (
	"context"
	
	"github.com/minipkg/fixed_time_ticker"
)
// usual in config
const (
    JobInterval   = 24 * time.Hour
    JobTimeHour   = 10 //	in UTC
    JobTimeMinute = 0
    JobTimeSecond = 0
)

ctx, cansel := context.WithCancel(context.Background())
t   := fixed_time_ticker.NewTicker(JobInterval, JobTimeHour, JobTimeMinute, JobTimeSecond)

go func() {
    for {
        select {
        case <-ctx.Done():
            c.t.Stop()
            return
        case <-c.t.C:
            // useful work here
            c.t.SetNextTime2Tick()
        }
    }
}()

```