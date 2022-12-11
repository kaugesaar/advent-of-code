# Advent of Code

Learning Go with advent of code.


### Results

Running on a MacBook Pro (13-inch, M1, 16GB, 2020) 

| Day | ⭐️      | ⭐️⭐️     |
|-----|---------|---------|
| 1   | 0.100ms | 0.099ms |
| 2   | 0.094ms | 0.104ms |
| 3   | 0.019ms | 0.015ms |
| 4   | 0.713ms | 0.710ms |
| 5   | 0.578ms | 0.570ms |
| 6   | 0.057ms | 0.564ms |
| 7   | 0.043ms | 0.043ms |
| 8   | 0.396ms | 4.753ms |
| 9   | 0.858ms | 2.184ms |
| 10  | 0.017ms | 0.023ms |


## Run
From the root directory you can.

Run all
```
go run main.go
```

Run specific day
```
go run main.go {day}
```

Run specific day and part
```
go run main.go {day} {part}
```

## Benchmark
Run benchmarks that outputs copy paste friendly markdown formated table
```
go run ./benchmark/benchmark.go
```

## Scaffold
Set your AOC_SESSION_COOKIE environment variable for your machine. Then you can scaffold each day with. 

```
./scaffold {day}
```
