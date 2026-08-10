FROM mcr.microsoft.com/playwright:v1.62.1-noble AS dev

RUN apt-get update \
    && apt-get install -y --no-install-recommends x11vnc \
    && rm -rf /var/lib/apt/lists/*

ENV DISPLAY=:99

COPY ./deploy/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENV DISPLAY=:99
