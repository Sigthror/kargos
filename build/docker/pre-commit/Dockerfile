FROM python:3.9.0-alpine3.12

RUN apk add --no-cache git && \
    pip install pre-commit && \
    mkdir -p /pre-commit

WORKDIR /pre-commit

CMD ["pre-commit", "run"]
