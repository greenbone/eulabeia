# Logging

In eulabeia we introduced a logging framework which needs be initiated once in the main cmd:

```
import (
	_ "github.com/greenbone/eulabeia/logging/configuration"
	"github.com/rs/zerolog/log"
)
```

and can be used afterwards by simply importing zerolog:


```
import (
	"github.com/rs/zerolog/log"
)
```

It does support multiple LogLevels which are interpreted as follows.

## Panic

Panic is used to indicate a non resolvable error and is followed by a panic.

Usually it is a unforeseeably issue and therefore should contain useful information for a developer and operator.

```
log.Panic().Err(err).Msgf("A msg why it panics %s", aVariable)

```

## Fatal

Fatal is used to indicate a non resolvable error and is followed by an exit 1.

Usually it is a infrastructural issue and therefore should be dealt within the underlying platform.

Therefore it should contain useful information for an operator.

```
log.Fatal().Err(err).Msg("A msg for the operator of the platform")

```

## Error 

Error is used to indicate a non resolvable issue that is preventing some but not all functionality of a service.

```
log.Error().Err(err).Msg("A msg to explain what functionality is not working and hopefully why")

```

## Warn

Warn is used to indicate a self-resolvable issue that maybe prevented some functionality to work for a period of time.

```
log.Warn().Err(err).Msg("A msg to explain what functionality is not working and hopefully why")

```

## Info

Info is used to log non repetitive information about a running service.

```
log.Info().Msgf("Starting %s on %s", "director", host)

```

## Debug

Debug is used to log information which will be used to gather information about a internal state. 

Usually it is turned on when a Error, Fatal, Panic or Warn did occur to get more information for e.g. a bug ticket.

Debug may be repetitive but should not be overwhelming.

```
log.Debug().Msgf("Started doing %s", "feed update")
```

## Trace

Developer information.

This should not be turned on in production systems and is only used by a developer.

```
log.Trace().Msgf("Called function %s with %v", "functionName", input)
```
