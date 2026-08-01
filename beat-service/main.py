from celery import Celery
from celery.signals import beat_init
import os
from celery.schedules import crontab
from datetime import timedelta
from rabbitmq.celeryconfig import queues, routes
from metrics.metrics import RAM_USAGE
from prometheus_client import start_http_server
import psutil
import threading
import time


cron_hour = int(os.getenv("CRON_HOUR", "2"))
cron_minute = int(os.getenv("CRON_MINUTE", "30"))

app = Celery('marketdemo')
app.conf.update(
    broker_url=os.environ.get("BROKER"),
    task_queues=queues.task_queues,
    task_routes=routes.task_routes,
    beat_schedule={
        'run-spider': {
            'task': 'main.run_modelo_spider',
            'schedule': crontab(hour=cron_hour, minute=cron_minute),
            'options': {'queue': os.getenv("CMD_SCRAPY_START_QUEUE", "cmd_scrapy_start_queue")},
        },
        'run-recommendations': {
            'task': 'main.calculate_recommendations_task',
            'schedule': crontab(hour=cron_hour, minute=cron_minute),
            'options': {'queue': os.getenv("CMD_RECOMMENDER_CALCULATE_QUEUE", "cmd_recommender_calculate_queue")},
        },
    },
    timezone='UTC',
)

def start_metrics_server():
    global metrics_started
    if metrics_started:
        return
    metrics_started = True
    port = int(os.getenv("PORT", 8000))
    start_http_server(port)
    while True:
        RAM_USAGE.set(psutil.virtual_memory().percent)
        time.sleep(5)

@beat_init.connect
def on_worker_ready(**kwargs):
    threading.Thread(
        target=start_metrics_server,
        daemon=True
    ).start()