FROM python:3.9.0-alpine3.12

RUN apk add --no-cache git && \
    pip install pre-commit && \
    mkdir -p /pre-commit

#ENV PRE_COMMIT_COLOR=never

WORKDIR /pre-commit
COPY ./ /pre-commit

RUN pre-commit install-hooks

CMD ["pre-commit"]
