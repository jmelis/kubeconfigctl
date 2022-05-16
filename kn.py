#!/usr/bin/env python3

import subprocess

import yaml
from pyfzf.pyfzf import FzfPrompt


def trimleft(s, prefix):
    if s.startswith(prefix):
        return s[len(prefix):]
    return s

def trimright(s, suffix):
    if s.endswith(suffix):
        return s[0:len(s)-len(suffix)]
    return s


def ncluster(c):
    c = trimleft(c, 'api-')
    c = trimright(c, '-p1-openshiftapps-com:6443')
    return c
    
fzf = FzfPrompt()

kube_config = yaml.safe_load(open('/Users/jmelis/.kube/config'))

clusters = list(set([ncluster(c['context']['cluster']) for c in kube_config['contexts']]))
cluster = fzf.prompt(clusters)[0]

namespaces = {}
for c in kube_config['contexts']:
    if cluster in c['context']['cluster']:
        namespaces[c['context']['namespace']] = c['name']

if len(namespaces) > 1:
    ns = fzf.prompt(namespaces.keys())[0]
else:
    ns = list(namespaces.keys())[0]

context = namespaces[ns]

subprocess.run(['oc','config','use-context', context])
