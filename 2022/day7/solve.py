import re

memory = {}
data = [line.strip() for line in open('input.txt', 'r')]

path = []
for line in data:
    if '$ ls' in line:
        pass
    elif line == '$ cd ..':
        path.pop()
    elif line == '$ cd /':
        path.append('/')
        memory['/'] = 0
    elif '$ cd' in line:
        location = line.split(' ')[2]
        current_path = path[-1] + '/' + location
        path.append(current_path)
        memory[current_path] = 0
    elif 'dir' in line:
        pass
    elif re.search("^\d+(?= )", line):
        size = int(line.split(' ')[0])
        for location in path:
            memory[location] += size
    else:
        raise Exception('illegal line: ' + line)

unused_space = 70000000 - memory['/']
print(sum([size for size in memory.values() if size < 100_000]))
print(min([size for size in memory.values() if unused_space + size >= 30000000]))