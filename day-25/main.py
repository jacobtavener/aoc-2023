import networkx as nx

def parse_file(filename) -> nx.Graph:
    G = nx.Graph()
    with open(filename, 'r') as f:
        for line in f:
            component, connections = line.strip().split(': ')
            connections = connections.split(' ')
            G.add_edges_from((component, connection) for connection in connections)
    return G

def part_1():
    G = parse_file("full.txt")
    # the initial graph is 3 edge connected since we know we can cut 3 edges and split the graph into 2 components
    # so we want to find the 4 edge connected components
    total = 1
    for edge_components in nx.k_edge_components(G, k=4):
        total *= len(edge_components)
    print("Total:", total)


if __name__ == '__main__':
    part_1()