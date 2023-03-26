orbit_map = []
with open("input.txt") as file:
    for line in file:
        orbit_map.append(line.strip())


def orbit_count(orbit_map):
    orbits = {}
    for orbit in orbit_map:
        components = orbit.split(")")
        middle = components[0]
        orbiter = components[1]
        if orbiter in orbits:
            orbits[orbiter].append(middle)
        else:
            orbits[orbiter] = [middle]
    
    def iter_orbit(orbiters):
        count = 0
        for orbiter in orbiters:
            if orbiter in orbits:
                count += 1
                count += iter_orbit(orbits[orbiter])
        return count

    return iter_orbit(orbits.keys())
    
print(orbit_count(orbit_map))

assert orbit_count(["COM)B", "B)C", "C)D" ,"D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"]) == 42


