import math

with open('input.txt', 'r') as file:
    
    total_fuel_needed = 0
    
    for line in file:
        mass = int(line.strip())
      
        while 0 < mass:
            fuel = math.floor(mass / 3) - 2
            if 0 < fuel:
                total_fuel_needed += fuel
            mass = fuel

    print(total_fuel_needed)
