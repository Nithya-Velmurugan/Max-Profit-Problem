# Max Profit Problem

This repository contains the solution to the "Max Profit Problem" coding assignment, implemented in Go.

## Problem Statement

Mr. X owns a strip of land on Mars and can develop it into either Theatres, Pubs, or Commercial Parks. Each property has different build times and earning potentials per unit of time operational. The objective is to determine the optimal mix of properties to develop in order to maximize absolute earnings within a given `n` units of time.

- **Theatre**: 5 units of time to build, earns $1500/unit time.
- **Pub**: 4 units of time to build, earns $1000/unit time.
- **Commercial Park**: 10 units of time to build, earns $2000/unit time.

*(Properties cannot be built in parallel. They must be developed sequentially).*

## Approach & Algorithm

The core of the problem requires maximizing absolute profit considering the **time value** of each property. As properties continue to earn money for every unit of time they are operational *after* being built, the traditional flat-profit Knapsack Dynamic Programming approach is invalid here. 

Instead, our approach leverages a highly-efficient **greedy-optimized combinatorial algorithm**:
1. **Prioritization by Profit Velocity**: We mathematically deduce that the earning-to-build-time ratio (profit velocity) strictly commands an optimal sequential sequence: Theatres (`$300/u`) > Pubs (`$250/u`) > Commercial Parks (`$200/u`). 
2. **Combinatorial State Evaluation**: Using a bounded search space `O((n/5) * (n/4) * (n/10))`, the algorithm natively evaluates all valid building multisets and correctly calculates the rolling operational time-elapsed profits. 
3. **Optimality and Speed**: The implementation evaluates viable constraint-abiding permutations in near $O(n^3)$ operations. This establishes an exact formulation for maximum absolute profit, naturally accounts for tying multiple solutions dynamically, and establishes an instantaneous sub-millisecond calculation for realistic time bounds.

## How to Run

```bash
go run main.go
```
When prompted, enter a Time Unit integer (e.g., `7`, `8`, `13`) to see the maximum earnings and optimal developmental mixtures!
