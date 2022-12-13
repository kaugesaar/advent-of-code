# Advent of Code

Learning Go with advent of code.


### Results

Running on a MacBook Pro (13-inch, M1, 16GB, 2020) 

| Day | ⭐️      | ⭐️⭐️      |
|-----|---------|----------|
| 1   | 0.108ms | 0.111ms  |
| 2   | 0.104ms | 0.123ms  |
| 3   | 0.027ms | 0.024ms  |
| 4   | 0.757ms | 0.720ms  |
| 5   | 0.586ms | 0.591ms  |
| 6   | 0.065ms | 0.569ms  |
| 7   | 0.051ms | 0.051ms  |
| 8   | 0.403ms | 4.690ms  |
| 9   | 0.864ms | 2.176ms  |
| 10  | 0.017ms | 0.023ms  |
| 11  | 0.063ms | 13.799ms |
| 12  | 0.132ms | 0.109ms  |


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
go run main.go -bench
```

## Scaffold
Set your AOC_SESSION_COOKIE environment variable for your machine. Then you can scaffold each day with. 

```
./scaffold.sh {day}
```
