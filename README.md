# time-resolution-measurement

According to the Go programming language document, the time returned by `time.Now()` returns an instant in time with nanosecond precision.

However, the **time resolution** that you can utilize may be lower, meaning that you may not be able to measure a very short period of time. Why does this happen? My understanding is that although the instant returned by `time.Now()` has a precision of nanosecond, the minimum (nanosecond) interval that you can observe may depend on the operating system. For example, if you call `time.Now()` once and get the value `20000`, it could happen that in the following `10000` nanoseconds, even if you call `time.Now()` again, you will still get `20000`. After `10000` nanoseconds, you may get `30000` by calling `time.Now()`. In this case, `30000 - 20000 = 10000` is the **time resolution** that you can utilize.

When you run this Go program, it tries to obtain an upper bound of the **time resolution** in your current environment.

The **measurement resolution** is the minimal **time resolution** that can be "measured" by this Go program. Since each "measure" needs to run a piece of Go code, it also takes some time. This time is the **measurement resolution**.

You may wonder why I can measure the **measurement resolution**, since sometimes it is smaller than the **time resolution** on the operating system. Well, my solution is to repeat the measurement computation `1000000` times. If it uses `T` nanoseconds to complete this, the **measurement resolution** will be `T / 1000000`.

You can read the code to understand some details.
