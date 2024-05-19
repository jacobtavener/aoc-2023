from functools import lru_cache

@lru_cache(maxsize=None)
def findValidPermuations(condition, groups, result=0):
    if len(groups) == 0:
        all_groups_assigned = "#" not in condition
        return all_groups_assigned

    total_springs = len(condition)
    total_working = sum(groups)
    cur_group, groups = groups[0], groups[1:]
    max_index_next_group = total_springs - total_working - len(groups) + 1

    for i in range(max_index_next_group):
        upper_bound = i + cur_group
        if upper_bound > total_springs or "#" in condition[:i]:
            break
        
        solidGroup = "." not in condition[i : upper_bound]
        continuedGroup = condition[upper_bound : upper_bound + 1] == "#"
        validGroupFound = solidGroup and not continuedGroup
        if validGroupFound:
            result += findValidPermuations(condition[upper_bound + 1:], groups)

    return result

    

with open("full.txt", "r") as file:
    reports = [x.split() for x in file.read().splitlines()]
    total = 0
    for report in reports:
        condition, groups = report
        condition = "?".join([condition] * 5)
        groups = tuple(map(int, groups.split(",")))*5
        total += findValidPermuations(condition, groups)
    print(total)