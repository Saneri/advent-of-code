def process_code(code):
    i = 0
    while True:
        command = code[i]
        i += 1
        if command == 1 or command == 2:
            position1 = code[i]
            i += 1
            position2 = code[i]
            i += 1
            destination = code[i]
            i += 1

            if command == 1:
                code[destination] = code[position1] + code[position2]
            elif command == 2:
                code[destination] = code[position1] * code[position2]

        elif command == 99:
            break
        else:
            break
    return code

optocode = []
with open('input.txt', 'r') as file:
    arr = file.readline().strip().split(',')
    for char in arr:
        optocode.append(int(char))

for i in range(0,100):
    for j in range(0,100):
        code = optocode.copy()
        code[1] = i
        code[2] = j
        if process_code(code)[0] == 19690720:
            print("i:{} - j:{}".format(i, j))
            print(100 * i + j)
