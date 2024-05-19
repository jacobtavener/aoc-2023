from typing import List
from collections import defaultdict

class Brick:

    def __init__(self, end1: List[int], end2: List[int], supported_by: List["Brick"]=[]):
        self.end1 = end1
        self.end2 = end2

        self.x1 = end1[0]
        self.y1 = end1[1]
        self.z1 = end1[2]

        self.x2 = end2[0]
        self.y2 = end2[1]
        self.z2 = end2[2]

        self.supporting = set()
        self.supported_by = supported_by
        self.occupied = set()

    @property
    def min_z(self):
        return min(self.z1, self.z2)
    
    @property
    def max_z(self):
        return max(self.z1, self.z2)
    
    @property
    def min_x(self):
        return min(self.x1, self.x2)
    
    @property
    def max_x(self):
        return max(self.x1, self.x2)
    
    @property
    def min_y(self):
        return min(self.y1, self.y2)
    
    @property
    def max_y(self):
        return max(self.y1, self.y2)
    
    
    def __repr__(self):
        return f"Brick({self.end1}, {self.end2})"
    
    def __str__(self):
        return f"Brick({self.end1}, {self.end2})"


def parse_input(filename):

    with open(filename, 'r') as f:
        data = f.read()

    brickList = []

    for line in data.splitlines():
        end1, end2 = line.split('~')
        end1 = [int(x) for x in end1.split(',')]
        end2 = [int(x) for x in end2.split(',')]
        brick = Brick(end1, end2)
        brickList.append(brick)
    
    brickList.sort(key=lambda x: x.min_z)
        
        

    return brickList


def drop_brick(max_map, brick, tower):
    max_z = max(max_map[(x,y)][0] for x in range(brick.min_x, brick.max_x + 1) for y in range(brick.min_y, brick.max_y + 1))
    supporting_bricks = set([max_map[(x,y)][1] for x in range(brick.min_x, brick.max_x + 1) for y in range(brick.min_y, brick.max_y + 1) if max_map[(x,y)][0] == max_z and max_map[(x,y)][1] is not None])
    z_offset = max(brick.min_z - max_z - 1, 0)
    dropped_brick = Brick(
        end1=(brick.x1, brick.y1, brick.z1 - z_offset),
        end2=(brick.x2, brick.y2, brick.z2 - z_offset),
    )
    for sb in supporting_bricks:
        sb.supporting.add(dropped_brick)
    dropped_brick.supported_by = supporting_bricks
    return dropped_brick

def find_disintegratable_map(brickList):
    tower = []
    max_map = defaultdict(lambda: (0, None))
    for brick in brickList:
        dropped_brick = drop_brick(max_map, brick, tower)
        tower.append(dropped_brick)
        for x in range(dropped_brick.min_x, dropped_brick.max_x + 1):
            for y in range(dropped_brick.min_y, dropped_brick.max_y + 1):
                max_map[(x,y)] = (dropped_brick.max_z, dropped_brick)
    
    disintegratable = {brick: True for brick in tower}
    for brick in tower:
        if len(brick.supported_by) == 1:
            disintegratable[list(brick.supported_by)[0]] = False
    return disintegratable, tower

def part1(filename):
    brickList = parse_input(filename)
    disintegratable, _= find_disintegratable_map(brickList)
    return sum(disintegratable.values())

def find_all_supporting_bricks(disintegratable, brick, visited=None):
    if visited is None:
        visited = set()

    visited.add(brick)

    for supporting_brick in brick.supporting:
            if supporting_brick not in visited:
                if supporting_brick.supported_by <= visited:
                    find_all_supporting_bricks(disintegratable, supporting_brick, visited)

    return visited


def part_2(filename):
    brickList = parse_input(filename)
    disintegratable, tower = find_disintegratable_map(brickList)

    total = 0
    for brick in tower:
        if not disintegratable[brick]:
            supported_bricks = find_all_supporting_bricks(disintegratable, brick)
            total += len(supported_bricks) - 1 # -1 to exclude the brick itself

    return total

            





if __name__ == '__main__':
    # print(part1("full.txt"))
    print(part_2("full.txt"))