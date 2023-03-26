
optocode = []
with open('input.txt', 'r') as file:
    arr = file.readline().strip().split(',')
    for char in arr:
        optocode.append(int(char))

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
            print("Fatal error!")
            break
    return code

optocode = process_code(optocode)
print("Answer: {}".format(optocode[0]))


assert process_code([1,0,0,0,99]) == [2,0,0,0,99]
assert process_code([2,3,0,3,99]) == [2,3,0,6,99]
assert process_code([2,4,4,5,99,0]) == [2,4,4,5,99,9801]
assert process_code([1,1,1,4,99,5,6,0,99]) == [30,1,1,4,2,5,6,0,99]
