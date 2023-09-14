## Overview

Introduction:

This proof of work miner is based on Argon2ID algorithm which is both GPU and ASIC resistant.
It allows all participants to mine blocks fairly. Your mining speed is directly proportional to
the number of miners you are running (you can run many on a single computer). The difficulty of
mining is auto adjusted based on the verifier node algorithm which aproximately targets production
speed of 1 block per second.

## Installation

### CLI

Install all the required modules by executing the command below. Make sure you have at least python3 and pip3 installed in order to proceed.

`pip install -U -r requirements.txt`

To start your miner just execute this command. Note you should adjust account at the top of the file to be your ethereum address if you want to claim your blocks and superblocks later

```sh
python3 miner.py
```

### Docker(Recommand)

Install docker first, refer to [here](https://docs.docker.com/engine/install/)

```sh
# use docker run a single container
docker run -it --rm \
    -e ACCOUNT=0xF120007d00480034fAf40000e1727C7809734b20 \
    -e STAT_CYCLE=100000 \
    -e DIFFICULTY=1 \
    -e CORE=1 \
    cnsumi/xen-miner:latest \
    -d
```

### Docker compose(Recommand)

Install docker compose first, refer to [here](https://docs.docker.com/compose/install/)

```sh
# use docker compose run scaleable container
# know your machine cpu cores num first
nproc --all # 8

docker compose up --scale miner=8 -d
# this will run 8 container to mine
```

# donate

EVM: `0xF120007d00480034fAf40000e1727C7809734b20`
