#!/usr/bin/env python3

import os
import multiprocessing as mp

from common import bench_and_trace_file_names


def run(bench, trace_file_name, num_nodes, routing, selection):
    dir = 'results/' + str(num_nodes) + '/' + routing + '/' + selection + '/' + bench

    os.system('rm -fr ' + dir)
    os.system('mkdir -p ' + dir)

    cmd_run = '~/GoProjects/bin/heo -d=' + dir + ' -b=' + bench + ' -f=' + trace_file_name \
              + ' -n=' + str(num_nodes) + ' -r=' + routing + ' -s=' + selection \
              + ' -c=' + str(100000000)
    print(cmd_run)
    os.system(cmd_run)


def run_experiment(args):
    bench, trace_file_name, num_nodes, routing, selection = args
    run(bench, trace_file_name, num_nodes, routing, selection)


experiments = []


def run_experiments():
    num_processes = mp.cpu_count()
    pool = mp.Pool(num_processes)
    pool.map(run_experiment, experiments)

    pool.close()
    pool.join()


def add_experiment(bench, trace_file_name, num_nodes, routing, selection):
    args = bench, trace_file_name, num_nodes, routing, selection
    experiments.append(args)


def add_experiments(bench, trace_file_name):
    add_experiment(bench, trace_file_name, 64, 'OddEven', 'BufferLevel')

for (bench, trace_file_name) in bench_and_trace_file_names:
    add_experiments(bench, trace_file_name)

run_experiments()
