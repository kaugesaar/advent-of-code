# Advent of Code

Learning Go with advent of code.


### Results

Running on a MacBook Pro (13-inch, M1, 16GB, 2020) 

| Day | ⭐️      | ⭐️⭐️      |
|-----|---------|----------|
| 1   | 0.109ms | 0.108ms |
| 2   | 0.101ms | 0.113ms |
| 3   | 0.027ms | 0.023ms |
| 4   | 0.730ms | 0.726ms |
| 5   | 0.591ms | 0.593ms |
| 6   | 0.064ms | 0.572ms |
| 7   | 0.051ms | 0.051ms |
| 8   | 0.404ms | 4.706ms |
| 9   | 0.863ms | 2.182ms |
| 10  | 0.017ms | 0.023ms |
| 11  | 0.064ms | 13.87ms |
| 12  | 0.132ms | 0.106ms |
| 13  | 1.015ms | 1.717ms |
| 14  | 1.811ms | 72.58ms |
| 15  | 317.6ms | 215.8ms |


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
