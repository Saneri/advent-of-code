import math

with open('input.txt', 'r') as file:
    
    total_fuel_needed = 0
    
    for line in file:
        mass = int(line.strip())
        total_fuel_needed += math.floor(mass / 3) - 2

    print(total_fuel_needed)
