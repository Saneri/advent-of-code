import ast
import functools

def compare_packets(left, right):
    if type(left) != list:
        left = [left]
    if type(right) != list:
        right = [right]

    for zip_left, zip_right in zip(left, right):
        if type(zip_left) == int and type(zip_right) == int:
            if zip_left - zip_right != 0:
                return zip_left - zip_right

        if type(zip_left) == list or type(zip_right) == list:
            ret = compare_packets(zip_left, zip_right)
            if ret != 0:
                return ret

    if len(left) != len(right):
        return len(left) - len(right)

    return 0


file = open('input.txt', 'r').read()
data = [[ast.literal_eval(p) for p in pair.split('\n')] for pair in file.split('\n\n')]

sum = 0
for i, pair in enumerate(data):
    if compare_packets(pair[0], pair[1]) < 0:
        sum += i+1
print(sum)

divider_packets = [[[2]],[[6]]]
packet_list = [ast.literal_eval(p) for p in file.replace('\n\n', '\n').split('\n')]
packet_list.extend(divider_packets)

prod = 1
for i, packet in enumerate(sorted(packet_list, key=functools.cmp_to_key(compare_packets))):
    if packet in divider_packets:
        prod *= i+1
print(prod)
