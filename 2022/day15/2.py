import re
data = [[int(num) for num in re.findall(r"[-\d]+", line.strip())] for line in open('input.txt', 'r')]

def manhattan_distance(a, b):
    return abs(a[0]-b[0]) + abs(a[1]-b[1])

max = 4000000

def is_unspotted(point, sensors):
    for sensor in sensors.keys():
        dist_from_point = manhattan_distance(point, sensor)
        dist_from_beacon = sensors[sensor]
        if dist_from_point <= dist_from_beacon:
            return False
    return True

sensors = {}
for line in data:
    sensor = (line[1], line[0])
    nearest_beacon = (line[3], line[2])
    sensors[sensor] = manhattan_distance(sensor, nearest_beacon)

for line in data:
    sensor = (line[1], line[0])
    nearest_beacon = (line[3], line[2])
    dist = manhattan_distance(sensor, nearest_beacon) +1
    for j in range(sensor[0]-dist, sensor[0]+dist+1):
        inverse_dist = dist-abs(sensor[0] - j)
        gap = 2*inverse_dist
        if gap == 0:
            gap = 1
        for i in range(sensor[1]-inverse_dist, sensor[1]+1+inverse_dist, gap):
            point = (j, i)
            if 0 <= j <= max and 0 <= i <= max and is_unspotted(point, sensors):
                print(i * 4000000 + j)
                exit()
