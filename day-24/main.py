import z3
from collections import namedtuple

HailData = namedtuple('HailData', ['initial_position', 'velocity'])

def read_data(file_path):
    with open(file_path) as file:
        data = [
            HailData(
                initial_position=[int(coord) for coord in initial_position.split(',')],
                velocity=[int(vel) for vel in velocity.split(',')]
            )
            for line in file
            for initial_position, velocity in (line.split('@'),)
        ]
    return data

def main(filename: str):
    hail_data = read_data(filename)

    rock_initial_position = z3.RealVector('rock_initial_position', 3)
    rock_velocity = z3.RealVector('rock_velocity', 3)
    time = z3.RealVector('t', 3)

    solver = z3.Solver()
    equations = [
        rock_initial_position[d] + rock_velocity[d] * t == hail.initial_position[d] + hail.velocity[d] * t
        for t, hail in zip(time, hail_data) for d in range(3)
    ]
    print(equations)
    solver.add(*equations)
    if solver.check() == z3.sat:
        print(solver.model().eval(sum(rock_initial_position)))
    else:
        print("No solution found")
                 
if __name__ == "__main__":
    main("example.txt")