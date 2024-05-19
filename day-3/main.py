import re

from typing import List

PART_PATTERN = re.compile('\d+')
SYMBOL_PATTERN = re.compile('[^.0-9]')

def parse_file(filename: str = 'full_file.txt') -> List[str]:
    with open(filename, 'r') as f:
        file_lines = f.readlines()
    return file_lines

def part_1():
    file_lines = parse_file()

    offset = len(file_lines)
    symbol_positions = []  
    for i, line in enumerate(file_lines):
        line = line.strip()
        symbol_positions += [m.start(0) + offset*i for m in SYMBOL_PATTERN.finditer(line)]
        
    symbol_positions = set(symbol_positions)


    total = 0
    for i, line in enumerate(file_lines):
        line = line.strip()
        for m in PART_PATTERN.finditer(line):
            num = m.group(0)
            for j in [-offset, 0, offset]:
                index_offset = i*offset+j
                lower_bound = m.start(0)+j+(i*offset) - 1
                upper_bound = m.end(0)+j+(i*offset) + 1
                possible_positions = set(range(max(index_offset, lower_bound), min(index_offset+offset, upper_bound)))
                if symbol_positions.intersection(possible_positions):
                    total += int(num)
                    break

    print(f"Part 1: {total}")


def part_2():
    file_lines = parse_file()

    offset = len(file_lines)
    number_map = {}
    for i, line in enumerate(file_lines):
        line = line.strip()
        for m in PART_PATTERN.finditer(line):
            number_map[range(m.start(0) + i*offset, m.end(0)+ i*offset)] = int(m.group(0))

    total = 0
    for i, line in enumerate(file_lines):
        line = line.strip()
        for m in SYMBOL_PATTERN.finditer(line):
            number_matches = []
            for j in [-offset, 0, offset]:
                index_offset = i*offset+j
                lower_bound = m.start(0)+j+(i*offset) - 1
                upper_bound = m.end(0)+j+(i*offset) + 1
                possible_positions = set(range(max(index_offset, lower_bound), min(index_offset+offset, upper_bound)))
                for key, value in number_map.items():
                    if possible_positions.intersection(set(key)):
                        number_matches.append(value)
            if len(number_matches) == 2:
                total += number_matches[0]*number_matches[1]

    print(f"Part 2: {total}")
if __name__ == '__main__':
    part_1()
    part_2()