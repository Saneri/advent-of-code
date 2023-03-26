
optocode = []
with open('input.txt', 'r') as file:
    arr = file.readline().strip().split(',')
    for char in arr:
        optocode.append(int(char))

def choose_mode(position, mode, code):
    if mode == 0:
        return code[position]
    elif mode == 1:
        return position
    raise("invalid mode!")

def process_code(code):
    i = 0
    while True:
        instruction = "{:05d}".format(code[i])
        command = int(instruction[3:5])
        mode1 = int(instruction[2])
        mode2 = int(instruction[1])
        mode3 = int(instruction[0])
        
        i += 1
        if command == 1 or command == 2:
            argument1 = choose_mode(i, mode1, code)
            i += 1
            argument2 = choose_mode(i, mode2, code)
            i += 1
            destination = choose_mode(i, mode3, code)
            i += 1

            if command == 1:
                code[destination] = code[argument1] + code[argument2]
            elif command == 2:
                code[destination] = code[argument1] * code[argument2]

        elif command == 3:
            destination = choose_mode(i, mode1, code)
            i += 1
            value = int(input())
            code[destination] = value
        
        elif command == 4:
            position = choose_mode(i, mode1, code)
            i += 1
            print(code[position])
        
        elif command == 99:
            break
        else:
            print("Fatal error!")
            break
    return code

optocode = process_code(optocode)


assert process_code([1,0,0,0,99]) == [2,0,0,0,99]
assert process_code([2,3,0,3,99]) == [2,3,0,6,99]
assert process_code([2,4,4,5,99,0]) == [2,4,4,5,99,9801]
assert process_code([1,1,1,4,99,5,6,0,99]) == [30,1,1,4,2,5,6,0,99]

assert process_code([1002,4,3,4,33]) == [1002,4,3,4,99]
assert process_code([1101,33,66,4,0]) == [1101,33,66,4,99]
