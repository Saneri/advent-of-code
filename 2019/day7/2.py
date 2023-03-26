from intcode import process_code

from itertools import permutations

def amplifier_row_output(amp_seq, code):
    process_memory = [code] * len(amp_seq)
    pointer_memory = [0] * len(amp_seq)
    amp_input = 0
    end_command = False
    last_output = 0
    while not end_command:
        i = 0
        for amp in amp_seq:
            print("amp {}".format(amp))
            amp_input, process_memory[i], pointer_memory[i] = process_code(process_memory[i].copy(), [amp, amp_input], True, pointer_memory[i])
            i += 1
            print(amp_input)
            print("pointer: {}".format(i))
            if i == len(amp_seq) -1 and amp_input != 99:
                last_output = amp_input
            if amp_input == 99:
                end_command = True
            print(process_memory)
            print(pointer_memory)
    print(last_output)
    return last_output

def get_biggest_output(code):
    amp_seq = [5,6,7,8,9]
    max_output = 0
    for seq in permutations(amp_seq, len(amp_seq)):
        output = amplifier_row_output(seq, code)
        if output > max_output:
            max_output = output
    return max_output
    


code = []
with open('input.txt', 'r') as file:
    firstline = file.readline().split(',')
    for code_str in firstline:
        code.append(int(code_str))


#print(get_biggest_output(code))
#print(process_code([3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26, 27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5]))
assert amplifier_row_output([9,8,7,6,5], [3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26, 27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5]) == 139629729


#assert get_biggest_output([3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26, 27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5]) == 139629729
#assert get_biggest_output([3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54, -5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4, 53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10]) == 18216
