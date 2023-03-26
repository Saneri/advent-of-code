data = [line.strip() for line in open('input.txt', 'r')]

wanted_indeces = set([20, 60, 100, 140, 180, 220])
crt_width = 40

x = 1
cycles = 0
add = None
it = iter(data)
sum = 0

def draw(position, sprite):
    if (position-1) % crt_width in range(sprite-1, sprite+2):
        print('#', end ="")
    else:
        print('.', end ="")
    if i % crt_width == 0:
        print()

for i in range (1, 241):
    if cycles == 0:
        instruction = next(it)
        match instruction.split(' ')[0]:  
            case 'addx':
                cycles = 2
                add = int(instruction.split(' ')[1])
            case 'noop':
                cycles = 1
                add = None
    draw(i, x)
    if i in wanted_indeces:
        sum += i * x
    cycles -= 1
    if add and cycles == 0:
        x += add

print(sum)