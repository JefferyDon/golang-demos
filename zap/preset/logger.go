package preset

import "go.uber.org/zap"

// runZapExampleLoggerDemo shows how to use zap.NewExample
// to create a logger and how this logger works.
//
// Output:
//
//	{"level":"info","msg":"This message is logged by sugaredLogger."}
//	{"level":"info","msg":"This message is logged by logger."}
//	{"level":"info","msg":"This is a new message logged by sugaredLogger","key1":"I am the value1","key2":"I am the value2"}
//	{"level":"info","msg":"This is a new message logged by logger","key1":"I am the value1","key2":1}
func runZapExampleLoggerDemo() error {
	logger := zap.NewExample()
	// we would like to flush logs in buffer when program exists.
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	// First we try the simplest way to log.
	sugaredLogger.Info("This message is logged by sugaredLogger.")
	logger.Info("This message is logged by logger.")

	// Second we try to make it a little fancy.
	sugaredLogger.Infow("This is a new message logged by sugaredLogger",
		"key1", "I am the value1",
		"key2", "I am the value2")
	logger.Info("This is a new message logged by logger",
		zap.String("key1", "I am the value1"),
		zap.Int8("key2", 1))

	return nil
}

// runZapDevelopmentLoggerDemo shows how to use zap.NewDevelopment
// to create a logger and how this logger works.
//
// Output:
//
//	 2023-11-09T10:42:28.732+0800	INFO	preset/logger.go:47	This message is logged by sugaredLogger.
//		2023-11-09T10:42:28.733+0800	INFO	preset/logger.go:48	This message is logged by logger.
//		2023-11-09T10:42:28.733+0800	INFO	preset/logger.go:51	This is a new message logged by sugaredLogger	{"key1": "I am the value1", "key2": "I am the value2"}
//		2023-11-09T10:42:28.733+0800	INFO	preset/logger.go:54	This is a new message logged by logger	{"key1": "I am the value1", "key2": 1}
func runZapDevelopmentLoggerDemo() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	// we would like to flush logs in buffer when program exists.
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	// First we try the simplest way to log.
	sugaredLogger.Info("This message is logged by sugaredLogger.")
	logger.Info("This message is logged by logger.")

	// Second we try to make it a little fancy.
	sugaredLogger.Infow("This is a new message logged by sugaredLogger",
		"key1", "I am the value1",
		"key2", "I am the value2")
	logger.Info("This is a new message logged by logger",
		zap.String("key1", "I am the value1"),
		zap.Int8("key2", 1))

	return nil
}

// runZapProductionLoggerDemo shows how to use zap.NewProduction
// to create a logger and how this logger works.
//
// Output:
//
//	{"level":"info","ts":1699497882.407916,"caller":"preset/logger.go:80","msg":"This message is logged by sugaredLogger."}
//	{"level":"info","ts":1699497882.408025,"caller":"preset/logger.go:81","msg":"This message is logged by logger."}
//	{"level":"info","ts":1699497882.4080331,"caller":"preset/logger.go:84","msg":"This is a new message logged by sugaredLogger","key1":"I am the value1","key2":"I am the value2"}
//	{"level":"info","ts":1699497882.408055,"caller":"preset/logger.go:87","msg":"This is a new message logged by logger","key1":"I am the value1","key2":1}
func runZapProductionLoggerDemo() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	// we would like to flush logs in buffer when program exists.
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	// First we try the simplest way to log.
	sugaredLogger.Info("This message is logged by sugaredLogger.")
	logger.Info("This message is logged by logger.")

	// Second we try to make it a little fancy.
	sugaredLogger.Infow("This is a new message logged by sugaredLogger",
		"key1", "I am the value1",
		"key2", "I am the value2")
	logger.Info("This is a new message logged by logger",
		zap.String("key1", "I am the value1"),
		zap.Int8("key2", 1))

	return nil
}
