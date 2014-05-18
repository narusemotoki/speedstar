#!/usr/bin/env python
# -*- coding: utf-8 -*-
from bottle import post, request, default_app
import json
import bjoern


@post('/')
def index():
    return json.dumps({
        'result': {
            'num': reduce(lambda x, y: {
                'add': lambda x, y: x + y,
                'sub': lambda x, y: x - y,
                'multi': lambda x, y: x * y,
                'div': lambda x, y: x / y,
                'mod': lambda x, y: x % y,
            }[y['operator']](x, y['num']), request.json['parts'], 1)
        }
    })


if __name__ == '__main__':
    bjoern.run(default_app(), '0.0.0.0', 8080)
