FROM python:3.10-slim as base

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

FROM base

LABEL maintainer="go.cnsumi@gmail.com"

COPY miner.py .

ENTRYPOINT [ "python3", "/app/miner.py" ]
