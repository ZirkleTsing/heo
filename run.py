#!/usr/bin/env python3

import os

from common import bench_and_trace_file_name_range, working_directory, max_cycles_range
from utils import add_experiment, run_experiments


def run(bench, trace_file_name, num_nodes, routing, selection, max_cycles):
    dir = working_directory(bench, num_nodes, routing, selection, max_cycles)

    os.system('rm -fr ' + dir)
    os.system('mkdir -p ' + dir)

    cmd_run = '~/GoProjects/bin/heo -d=' + dir + ' -b=' + bench + ' -f=' + trace_file_name \
              + ' -n=' + str(num_nodes) + ' -r=' + routing + ' -s=' + selection \
              + ' -c=' + str(max_cycles)
    print(cmd_run)
    os.system(cmd_run)


def run_experiment(args):
    run(*args)


experiments = []


for (bench, trace_file_name) in bench_and_trace_file_name_range:
    for max_cycles in max_cycles_range:
        add_experiment(experiments, bench, trace_file_name, 64, 'OddEven', 'BufferLevel', max_cycles)

run_experiments(experiments, run_experiment)
