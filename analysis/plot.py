#!/usr/bin/env python

import glob
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

def get_colormap(ips, column):
    colors = np.linspace(0, 1, len(ips))
    colormap = dict(zip(ips, colors))
    return [colormap[ip] for ip in column]

for filename in glob.iglob('../tslp/res-*.csv'):
    basename = filename[8:-4]
    data = pd.read_csv(filename, sep=';',
                names=['timestamp', 'ip1', 'ip2', 'time1', 'time2', 'diff'],
                index_col='timestamp')
    data.index = pd.to_datetime(data.index, unit='s')
    data = data.loc[data.index > '2018-12-05T04:00:00']
    ip2s = data['ip2'].unique()

    fig = plt.figure()
    plt.scatter(data.index, data['time1'],
                s=1,
                color='green')
    plt.scatter(data.index, data['time2'],
                s=1,
                c=get_colormap(ip2s, data['ip2']),
                cmap='autumn')
    axes = plt.gca()
    axes.set_ylim([0, 80])
    axes.set_xlim(['2018-12-05', '2018-12-17'])
    fig.autofmt_xdate()
    plt.xlabel('time')
    plt.ylabel('latency (ms)')
    plt.savefig(basename + '.pdf')
    plt.close(fig)
