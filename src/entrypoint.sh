#!/bin/bash
gunicorn --chdir /home/asn/app main:app -w 2 --threads 2 -b 0.0.0.0:5000