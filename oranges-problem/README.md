# Monte Carlo Simulation of Oranges

## Problem Statement

There are two boxes of oranges, each with N=30 oranges initially. We take one
orange from either box at random (50% probability). Once one of the boxes is
empty, what is the expected number of oranges left in the other box?

# DP/Recursive

```shell-session
capacity: 1     | expected oranges: 1.000000
capacity: 2     | expected oranges: 1.500000
capacity: 3     | expected oranges: 1.875000
capacity: 30    | expected oranges: 6.154690
capacity: 100   | expected oranges: 11.269696
capacity: 1000  | expected oranges: 35.678022
capacity: 10000 | expected oranges: 112.836506
```

## Monte Carlo

```shell-session
capacity: 1     | iterations: 9000     | estimate: 1.000000
capacity: 2     | iterations: 170000   | estimate: 1.499985
capacity: 3     | iterations: 253000   | estimate: 1.872293
capacity: 30    | iterations: 1085000  | estimate: 6.156124
capacity: 100   | iterations: 2542000  | estimate: 11.268692
capacity: 1000  | iterations: 6859000  | estimate: 35.653048
capacity: 10000 | iterations: 16491000 | estimate: 112.616423
```

