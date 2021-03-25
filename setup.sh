#!/bin/bash

py='python'
if [[ "$(python3 -V)" =~ "Python 3" ]]; then
  py='python3'
fi

"${py}" -m venv env
source env/bin/activate

pip install -r requirements.txt
