from common import bench_and_trace_file_names, working_directory
from utils import parse_result, to_csv, generate_plot

results = []

for bench, trace_file_name in bench_and_trace_file_names:
    for max_cycles in [10000, 100000, 1000000, 10000000]:
        results.append(
            parse_result(working_directory(bench, 64, 'OddEven', 'BufferLevel', max_cycles), bench=bench)
        )

    to_csv('results/general.csv', results, [
        ('Bench', lambda r: r.props['bench']),
        ('Max Cycles', lambda r: r.configs['MaxCycles']),
        ('Simulation Time (Seconds)', lambda r: r.stats['SimulationTimeInSeconds']),
        ('Throughput', lambda r: r.stats['Throughput']),
        ('Average Packet Delay', lambda r: r.stats['AveragePacketDelay']),
    ])

    generate_plot('results/general.csv',
                  'results/max_cycles_vs_simulation_time', 'Bench', 'Simulation Time (Seconds)',
                  'Max Cycles', 'Simulation Time (Seconds)')

    generate_plot('results/general.csv',
                  'results/max_cycles_vs_throughput', 'Bench', 'Throughput',
                  'Max Cycles', 'Throughput')

    generate_plot('results/general.csv',
                  'results/max_cycles_vs_average_packet_delay', 'Bench', 'Average Packet Delay',
                  'Max Cycles', 'Average Packet Delay')
