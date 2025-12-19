import math
import heapq
from itertools import combinations

with open('input.txt', 'r') as file:
    lines = [list(map(int, l.strip().split(','))) for l in file.readlines()]

def dist(a: list[int], b: list[int]) -> float:
    return math.sqrt((a[0] - b[0])**2 + (a[1] - b[1])**2 + (a[2] - b[2])**2)

circuits = {tuple(line): i for i, line in enumerate(lines)}
pq = []
for combo in combinations(lines, 2):
    d = dist(combo[0], combo[1])
    heapq.heappush(pq, (d, tuple(combo[0]), tuple(combo[1])))

while pq:
    _, box1, box2 = heapq.heappop(pq)
    cluster_a = circuits[box1]
    cluster_b = circuits[box2]
    if cluster_a != cluster_b:
        for c in circuits:
            if circuits[c] == cluster_a:
                circuits[c] = cluster_b

    if len(set(circuits.values())) == 1:
        print(box1[0] * box2[0])
        break