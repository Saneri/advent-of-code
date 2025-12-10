with open('input.txt', 'r') as file:
    lines = [line.strip() for line in file.readlines()]

def find_max_subsequence(s, length):
    n = len(s)
    if length > n:
        return None
    
    result = []
    start = 0
    while len(result) < length:
        end = n - length + len(result) + 1
        nums = s[start:end]
        max_val = str(max(int(x) for x in nums))
        max_idx = nums.index(max_val)
        start = start + max_idx + 1
        result.append(max_val)
    return int(''.join(result))

count = 0
for line in lines:
    max_subseq = find_max_subsequence(line, 12)
    if max_subseq:
        count += max_subseq
print(count)