package log

import "time"

func init() {
	now = func() time.Time {
		return time.Date(2020, time.April, 17, 10, 16, 0, 0, time.UTC)
	}
}

func ExampleNew_Simple_1() {
	log := New(FormatterSimple)
	log.Info("Hello World!")
	log.Info("Hello \"World\" 2!", "foo", 1, "bar", true, "baz", now())
	// Output:
	// 2020-04-17T10:16:00Z [INF] Hello World!
	// 2020-04-17T10:16:00Z [INF] Hello "World" 2! FOO: 1 BAR: true BAZ: 2020-04-17T10:16:00Z
}

func ExampleNew_Simple_2() {
	opts := []Option{
		OptionWithTimeFormat(time.Stamp),
		OptionWithStaticKV("hostname", "localhost"),
		OptionWithStaticKV("ip", "127.0.0.1"),
		OptionDisableLevel(true),
	}
	log := New(FormatterSimple, opts...)
	log.Info("Hello World!")
	log.Info("Hello \"World\" 2!", "foo", 1, "bar", true, "baz", now())
	// Output:
	// Apr 17 10:16:00 Hello World! IP: 127.0.0.1 HOSTNAME: localhost
	// Apr 17 10:16:00 Hello "World" 2! IP: 127.0.0.1 HOSTNAME: localhost FOO: 1 BAR: true BAZ: Apr 17 10:16:00
}

func ExampleNew_KV() {
	opts := []Option{
		OptionWithLevelKey("level"),
		OptionWithTimeKey("time"),
		OptionWithMessageKey("message"),
	}
	log := New(FormatterKV, opts...)
	log.Info("Hello \"World\"!", "foo", 1, "bar", true, "baz", now())
	// Output:
	// time: "2020-04-17T10:16:00Z" level: "Info" message: "Hello \"World\"!" foo: "1" bar: "true" baz: "2020-04-17T10:16:00Z"
}

func ExampleNew_JSON() {
	log := New(FormatterJSON)
	log.Info("Hello World!", "foo", 1, "bar", true, "baz", now())
	// Output:
	// {"LEVEL":"Info","MESSAGE":"Hello World!","TIME":"2020-04-17T10:16:00Z","bar":true,"baz":"2020-04-17T10:16:00Z","foo":1}
}

func ExampleNew_Debug_1() {
	log := New(FormatterSimple)
	log.Debug("This is a debug message")
	// Output:
}

func ExampleNew_Debug_2() {
	log := New(FormatterSimple, OptionEnableDebug(true))
	log.Debug("This is a debug message")
	// Output:
	// 2020-04-17T10:16:00Z [DEB] This is a debug message
}
