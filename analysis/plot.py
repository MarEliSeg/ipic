#!/usr/bin/env python

import glob
import pandas as pd
import matplotlib.pyplot as plt

for filename in glob.iglob('../tslp/res-*.csv'):
    basename = filename[8:-4]
    data = pd.read_csv(filename, sep=';',
                names=['timestamp', 'ip1', 'ip2', 'time1', 'time2', 'diff'],
                index_col='timestamp')
    data.index = pd.to_datetime(data.index, unit='s')
    data = data.loc[data.index > '2018-12-05T04:00:00']

    plt.figure()
    ax = data.plot(linestyle='', marker='.', markersize=2,
                y=['time1', 'time2'],
                ylim=(0, 80),
                label=['near end', 'far end'],
                color=['green', 'red'])
    ax.set_xlabel('time')
    ax.set_ylabel('latency (ms)')
    plt.savefig(basename + '.pdf')
