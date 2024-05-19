import collections
import math

def parse_input(filename):

    with open(filename, 'r') as f:
        data = f.read()

    process_map = {}

    for line in data.splitlines():
        module, outputs = line.split(' -> ')
        outputs = outputs.split(', ')

        if module == 'broadcaster':
            type = None
        else:
            type = module[0]
            module = module[1:]

        process_map[module] = (type, outputs)

    return process_map

def part1(filename):
    data = parse_input(filename)

    num_low = 0
    num_high = 0
    memory = {}

    input_map = collections.defaultdict(list)

    for node, (_, dests) in data.items():
        for d in dests:
            input_map[d].append(node)

    for node, (t, _) in data.items():
        if t is None:
            continue
        if t == '%':
            memory[node] = False
        if t == '&':
            memory[node] = {d:False
                            for d in input_map[node]}

    print(memory)

    for _ in range(1000):
        todo = [(None, 'broadcaster', False)]

        while todo:
            new_todo = []

            for src, node, is_high_pulse in todo:
                if is_high_pulse:
                    num_high += 1
                else:
                    num_low += 1

                info = data.get(node)
                if info is None:
                    continue

                t, dests = info
                if t == '%':
                    if is_high_pulse:
                        continue
                    state = memory[node]
                    memory[node] = not state
                    for d in dests:
                        new_todo.append((node, d, not state))
                    continue
                if t == '&':
                    state = memory[node]
                    state[src] = is_high_pulse

                    if sum(state.values()) == len(state):
                        to_send = False
                    else:
                        to_send = True

                    for d in dests:
                        new_todo.append((node, d, to_send))
                    continue
                if t is None:
                    for d in dests:
                        new_todo.append((node, d, is_high_pulse))
                    continue
                assert(False)

            todo = new_todo

    answer = num_low * num_high

    print(num_low, num_high)
    print(answer)

def part2(filename):
    data = parse_input(filename)

    num_low = 0
    num_high = 0
    memory = {}

    input_map = collections.defaultdict(list)

    for node, (_, dests) in data.items():
        for d in dests:
            input_map[d].append(node)

    for node, (t, _) in data.items():
        if t is None:
            continue
        if t == '%':
            memory[node] = False
        if t == '&':
            memory[node] = {d:False
                            for d in input_map[node]}
            
    print(input_map['rx'][0])

    print(input_map['kc'])
    sources = input_map[input_map['rx'][0]]

    print(sources)
    cycle_map = {}
    index = 0
    while len(cycle_map) < len(sources):
        index += 1

        todo = [(None, 'broadcaster', False)]

        while todo:
            new_todo = []

            for src, node, is_high_pulse in todo:
                if node in sources:
                    if not is_high_pulse:
                        if node not in cycle_map:
                            cycle_map[node] = index

                if is_high_pulse:
                    num_high += 1
                else:
                    num_low += 1

                info = data.get(node)
                if info is None:
                    continue

                t, dests = info
                if t == '%':
                    if is_high_pulse:
                        continue
                    state = memory[node]
                    memory[node] = not state
                    for d in dests:
                        new_todo.append((node, d, not state))
                    continue
                if t == '&':
                    state = memory[node]
                    state[src] = is_high_pulse
                    if sum(state.values()) == len(state):
                        to_send = False
                    else:
                        to_send = True

                    for d in dests:
                        new_todo.append((node, d, to_send))
                    continue
                if t is None:
                    for d in dests:
                        new_todo.append((node, d, is_high_pulse))
                    continue
                assert(False)

            todo = new_todo

    answer = num_low * num_high

    print(num_low, num_high)
    print(answer)
    print(multiple_lcm(*cycle_map.values()))

def multiple_lcm(*args):
    """Return the least common multiple of multiple integers."""
    if len(args) == 0:
        return None
    result = args[0]
    for i in range(1, len(args)):
        result = lcm(result, args[i])
    return result

def lcm(a, b):
    """Return the least common multiple of two integers."""
    return abs(a * b) // math.gcd(a, b)

if __name__ == '__main__':
    # part1('full.txt')
    part2('full.txt')

