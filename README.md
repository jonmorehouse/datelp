# Datelp 
> A go library for parsing and extracting dates from text

## Usage

```golang
package main

import (
  "datelp"
  "log"
)

func main() {
  input := []string{"next", "monday"}
  results, err := datelp.Parse(input)
  if err != nil {
    log.Info("No dates available")
  }
  
  for _, res := range results {
    log.Info(res.Position, res.Time)
  }

}


```

## First Version Supported Formats

~~~ text
June 1st
June 1st 2015
June 1 2015

Tuesday
next Tuesday
last Wednesday

Tomorrow
Yesterday
~~~


