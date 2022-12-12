# Advent of Code

Learning Go with advent of code.


### Results

Running on a MacBook Pro (13-inch, M1, 16GB, 2020) 

| Day | Part 1  | Part 2   |
|-----|---------|----------|
| 1   | 0.104ms | 0.110ms  |
| 2   | 0.027ms | 0.024ms  |
| 3   | 0.725ms | 0.722ms  |
| 4   | 0.586ms | 0.596ms  |
| 5   | 0.064ms | 0.572ms  |
| 6   | 0.051ms | 0.051ms  |
| 7   | 0.407ms | 4.757ms  |
| 8   | 0.868ms | 2.194ms  |
| 9   | 0.018ms | 0.023ms  |
| 10  | 0.064ms | 13.908ms |
| 11  | 0.134ms | 28.095ms |
| 12  | 0.133ms | 29.549ms |


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
