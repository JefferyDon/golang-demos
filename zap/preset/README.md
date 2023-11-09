# Preset Demos

In this package, I wrote 3 demos about how to use `zap.NewExample`, `zap.NewDevelopment` and 
`zap.NewProduction` to create `zap.Logger`, and also gave some examples about how to use
`zap.Logger` and `zap.SugaredLogger` to log messages.

# Note

- In `runZapDevelopmentLoggerDemo()` you can see log format is not in JSON format. This is because `zap.NewDevelopment` is using `console` as its log encoding format. For more detail, see code `go.uber.org/zap/config.go:NewDevelopmentConfig()`