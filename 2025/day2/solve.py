import re

def is_valid(num: str) -> bool:
    regex = r'^(\d+)\1+$'
    return re.search(regex, num) is None

with open('input.txt', 'r') as file:
    input = file.readline().strip().split(',')

sum = 0
for line in input:
    split = line.split('-')
    first = int(split[0])
    second = int(split[1])
    for i in range(first, second + 1):
        if not is_valid(str(i)):
            sum += i
print(sum)
    



