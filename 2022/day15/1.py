import re
data = [[int(num) for num in re.findall(r"[-\d]+", line.strip())] for line in open('input.txt', 'r')]

def manhattan_distance(a, b):
    return abs(a[0]-b[0]) + abs(a[1]-b[1])

coverage = set()
y = 2000000

for line in data:
    sensor = (line[1], line[0])
    nearest_beacon = (line[3], line[2])
    dist = manhattan_distance(sensor, nearest_beacon)
    for j in range(sensor[0]-dist, sensor[0]+dist+1):
        inverse_dist = dist-abs(sensor[0] - j)
        if j == y:
            for i in range(sensor[1]-inverse_dist, sensor[1]+1+inverse_dist):
                if (j, i) != nearest_beacon:
                    coverage.add((j, i))

print(len(coverage))