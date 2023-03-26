orbit_map = []
with open("input.txt") as file:
    for line in file:
        orbit_map.append(line.strip())


def shortest_dist(orbit_map):
    start = "YOU"
    destination = "SAN"

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
        orbit_path = []
        for orbiter in orbiters:
            if orbiter in orbits:
                orbit_path.append(orbiter)
                orbit_path += iter_orbit(orbits[orbiter])
        return orbit_path


    you_orbits = iter_orbit(orbits["YOU"])
    san_orbit = orbits["SAN"]
    san_dist = 0
    while True:
        san_orbit = san_orbit[0]
        if san_orbit in you_orbits:
            return san_dist + you_orbits.index(san_orbit)
        if san_orbit not in orbits:
            break
        san_dist += 1
        san_orbit = orbits[san_orbit]
    return -1


print(shortest_dist(orbit_map))

assert shortest_dist(["COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"]) == 4


