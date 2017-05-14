bench_and_trace_file_name_range = [
    ('simple_pthread', 'traces/simple_pthread.trace.21454.0'),
]

max_cycles_range = [10000, 100000, 1000000, 10000000, 100000000]


def working_directory(bench, num_nodes, routing, selection, max_cycles):
    return 'results/' + str(num_nodes) + '/' + routing + '/' + selection + '/' + bench + '/' + str(max_cycles)