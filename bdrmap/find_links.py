#!/usr/bin/env python

import json
import re
import csv
from collections import defaultdict

TRACES='./bdrmap.json'
BDR_OUT='./bdrmap.out'
OUTFILE='./links.csv'

routers = {}
links = defaultdict(list)

with open(BDR_OUT) as in_file:
    for line in in_file:
        if not line.startswith(' '):
            continue
        if 'silent' in line:
            continue
        meta = re.search('(\d+\.\d+\.\d+\.\d+)\*\n$', line)
        ip = meta.group(1)
        meta = re.search('^\s(\d+)', line)
        asn = meta.group(1)
        routers[ip] = asn

with open(TRACES) as traces:
    for trace in traces:
        obj = json.loads(trace)
        if obj['type'] != 'trace':
            continue
        if not 'hops' in obj:
            continue
        for hop in obj['hops']:
            for router in routers.keys():
                if hop['addr'] == router:
                    links[router].append((obj['dst'], int(hop['probe_ttl'])))

with open(OUTFILE, 'w', newline='') as link_file:
    writer = csv.writer(link_file, delimiter=';')
    for router, targets in links.items():
        writer.writerow([targets[0][0], targets[0][1]-1, targets[0][1], router, routers[router]])
