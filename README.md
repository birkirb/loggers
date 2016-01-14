# loggers : Golang Abstract Loggers
loggers define an abstract and common logging interface in three flavors.

## Inspiration

If you have been using Go for a while, you're probably asking yourself: "Why, oh why, wasn't the standard library logger made an interface!?"
Often was I faced with having to decide about what kind of logger I needed before I was ready to, wondering:

  * "Where's Golang's answer to log4j". 
  * Is there a log4go or log4golang?

Well this package should help. Install and call `log.Info("Log all my stuff")` and you're off and you can easily switch out loggers later with only a single line of code.

## Design

All loggers are interfaces and should be declared and used as such. The actual implementations can vary and can be easily switched out.

### Standard
Standard interface same as used by the Go standard library `log` package.

### Advanced
A common pattern for level discrete loggers using `debug`, `info`, `warn` and `error` levels, along with those defined by the standard interface.

### Contextual
A superset of Advanced, adds contextual logging, such that lines can have a number of additional parameters set to provide clearer seperation of message and context.

## Installation

    go get github.com/birkirb/loggers

## Usage

You can choose between declaring one of the three interfaces above that best suits your needs, or use the built in Contextual interface directly with a default standard library Logger implementation to do your bidding. You can switch out all the loggers later as long as they satisfy the right interface.

### Direct

You can use the loggers interface as a drop in replacement for the standard library logger.
Just change your import statement from "log" to "github.com/birkirb/loggers/log".
It should work just the same and you can make use of advanced and contextual methods only if you so decide.
You can then easily switch out the log package implementation later with your own logger as long as it implements the Contextual interface.

```Go
    log.Infof("Logger is started") // Defaults to stdout.
    log.Logger = stdlib.NewLogger(fileWriter, "", log.LstdFlags)
    log.Infof("Now logging to fileWriter") // writes to fileWriter
```

### Embedded

Declare your own project logging interface.

```Go
    var Logger loggers.Standard

    func init() {
        Logger = log.New(writer, "myapp", log.LstdFlags)
        Logger.Println("Logger is started")
    }
```

### Mappers

A few loggers have been mapped to the above interfaces and could thus be used with any of them.
Instead of the using the Standard logger as above, we could use the standard logger much like a leveled logger.

```Go
    var Logger loggers.Advanced

    func init() {
        Logger = stdlib.NewDefaultLogger()
        Logger.Info("Logger is started")
    }
```

A level mapper exist to easy with implementing plugins/mappers for other loggers that don't naturally implement any of the interfaces of loggers.
This can be found in the mappers package. 

## Existing mappers

* [Revel](https://github.com/revel/revel/) mapper ()
* [Logrus](https://github.com/Sirupsen/logrus) mapper ()

# Contributing

Any new mappers for different Go logging solutions would be most welcome.
