with open('input.txt', 'r') as file:
    lines = [l.strip('\n') for l in file.readlines()]

height = len(lines)
width = len(lines[0])
operators = lines[-1].split()
area = len(operators)-1

sum = 0
calc = []
for i in reversed(range(width)):
    num = []
    for j in range(height-1):
        if not lines[j][i].isspace():
            num += [lines[j][i]]
    if num:
        calc.append(''.join(num))
    if not num or i == 0:
        sum += eval(operators[area].join(calc))
        area -= 1
        calc = []
print(sum)