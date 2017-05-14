bench_and_trace_file_names = [
    ('simple_pthread', 'traces/simple_pthread.trace.21454.0'),
]


def working_directory(bench, num_nodes, routing, selection, max_cycles):
    return 'results/' + str(num_nodes) + '/' + routing + '/' + selection + '/' + bench + '/' + str(max_cycles)