# Datelp 
> A go library for parsing and extracting dates from text

[![Build Status](https://travis-ci.org/jonmorehouse/datelp.svg?branch=master)](https://travis-ci.org/jonmorehouse/sns-ingest)

**DISCLAIMER:** this library is currently a WIP. I'm using this particular version locally in [Datebook]() (a small cli tool for taking notes). Its worth mentioning that an overhaul and official first release are coming in the next few weeks once https://github.com/jonmorehouse/datelp/pull/3 has been finished.


## Usage

```golang
package main

import (
  "datelp"
  "log"
)

func main() {
  date, err := datelp.Parse("next tuesday")
  if err != nil {
    log.Info("No dates available")
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

