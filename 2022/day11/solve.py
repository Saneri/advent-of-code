import math

data = open('input.txt', 'r').read().split('\n\n')

class Monkey:
    def __init__(self, input):
        lines = [line.strip() for line in input.split('\n')]

        self.items = [int(x) for x in lines[1].split(': ')[1].split(', ')]
        self.operation = lines[2].split(': ')[1].split('d ')[1]
        self.divider = int(lines[3].split('y ')[1])
        self.trueMonkey = int(lines[4].split('y ')[1])
        self.falseMonkey = int(lines[5].split('y ')[1])
        self.inspection_count = 0
    
    def calculate_stress(self, old):
        return eval(str(old) + self.operation)

    def test(self, item):
        if item % self.divider == 0:
            return self.trueMonkey
        return self.falseMonkey

    def add_item(self, item):
        self.items.append(item)

    def inspect_all_items(self, monkey_list, isRelieved, common_divider):
        for _ in range(len(self.items)):
            item = self.items.pop(0)
            worry_level = self.calculate_stress(item)
            if isRelieved:
                worry_level = worry_level // 3
            else: 
                worry_level = worry_level % common_divider
            new_monkey = self.test(worry_level)
            monkey_list[new_monkey].add_item(worry_level)
            self.inspection_count += 1

    def __str__(self):
        return 'items: ' + str(self.items) + ' | operation: ' + self.operation + ' | divider: ' + str(self.divider) + ' | true: ' + str(self.trueMonkey) + ' | false: ' + str(self.falseMonkey)

monkeys = [Monkey(definition) for definition in data]
monkeys2 = [Monkey(definition) for definition in data]

def find_max_values_after_rounds(monkeys, rounds, relief):
    common_divider = math.prod([monkey.divider for monkey in monkeys])
    for _ in range (rounds):
        for monkey in monkeys:
            monkey.inspect_all_items(monkeys, relief, common_divider)
    counts = sorted([monkey.inspection_count for monkey in monkeys])
    return counts[-1] * counts[-2]

print(find_max_values_after_rounds(monkeys, 20, True))
print(find_max_values_after_rounds(monkeys2, 10000, False))