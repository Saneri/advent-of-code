with open('input.txt', 'r') as file:
    sections = file.read().split('\n\n')

shapes = [s.split('\n')[1:] for s in sections[:-1]]
regions = [s.split(':') for s in sections[-1].split('\n')]

heuristic = [3*4, 4*4, 3*5, 3*4, 3*6, 3*6] # approximate areas for doubles read from input
count = 0
for region in regions:
    width, height = list(map(int, region[0].strip().split('x')))
    quantities = list(map(int, region[1].strip().split(' ')))

    needed_area = 0
    for q, h in zip(quantities, heuristic):
        needed_area += q * h / 2
    if needed_area < width * height:
        count += 1

print(count)
