#!/usr/bin/env python

import json
from subprocess import call

FILE='./ixs_201802.jsonl'
TMPFILE='./ixp.txt.tmp'
OUTFILE='./ixp.txt'

with open(FILE) as in_file:
    with open(TMPFILE, 'w') as out_file:
        for line in in_file:
            if line.startswith('#'):
                continue
            obj = json.loads(line)
            for prefix in obj['prefixes']['ipv4']:
                out_file.write(prefix + '\n')

call("cat " + TMPFILE + " | sort | uniq > " + OUTFILE, shell=True)
call("rm " + TMPFILE, shell=True)
