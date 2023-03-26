def choose_mode(position, mode, code):
    if mode == 0:
        return code[position]
    elif mode == 1:
        return position
    raise("invalid mode!")

def process_code(code, inputs=None, ret_first_output=False, i=0):
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
            if inputs:
                value = inputs.pop(0)
            else:
                value = int(input())
            code[destination] = value
        
        elif command == 4:
            position = choose_mode(i, mode1, code)
            i += 1
            if ret_first_output:
                return code[position], code, i
            print(code[position])
        
        elif command == 5 or command == 6:
            argument1 = choose_mode(i, mode1, code)
            i += 1
            argument2 = choose_mode(i, mode2, code)
            i += 1
            if command == 5:
                boolean = code[argument1] != 0
            elif command == 6:
                boolean = code[argument1] == 0
            if boolean:
                i = code[argument2]

        elif command == 7 or command == 8:
            argument1 = choose_mode(i, mode1, code)
            i += 1
            argument2 = choose_mode(i, mode2, code)
            i += 1
            destination = choose_mode(i, mode3, code)
            i += 1
            if command == 7:
                boolean = code[argument1] < code[argument2]
            elif command == 8:
                boolean = code[argument1] == code[argument2]
            if boolean:
                code[destination] = 1
            else:
                code[destination] = 0

        elif command == 99:
            if ret_first_output:
                return 99, code, i
            break
        else:
            print("Fatal error!")
            break
    return code
