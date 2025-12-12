with open('input.txt', 'r') as file:
    lines = file.read().split('\n\n')

ranges = [(int(x[0]), int(x[1])) for x in [x.split('-') for x in lines[0].split('\n')]]

areas = [ranges[0]]
for range in ranges[1:]:
    overlapping = []
    for area in areas:
        # range outside area
        if range[1] < area[0] or range[0] > area[1]:
            continue
        overlapping.append(area)
        
    if overlapping:
        start = min([area[0] for area in overlapping] + [range[0]])
        end = max([area[1] for area in overlapping] + [range[1]])
        for area in overlapping:
            areas.remove(area)
        areas.append((start, end))
    if not overlapping:
        areas.append(range)

count = 0
for area in areas:
    count += area[1] - area[0] + 1
print(count)
