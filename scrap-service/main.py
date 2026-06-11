from prometheus_client import start_http_server
import threading
import psutil
import time
from marketdemo.metrics import RAM_USAGE
from celery import Celery, chain
import os
import subprocess
import psycopg2
from rabbitmq.celeryconfig import queues, routes

SCRAPY_QUEUE = os.getenv("CMD_SCRAPY_START_QUEUE", "cmd_scrapy_start_queue")

app = Celery('marketdemo')
app.conf.update(
    broker_url=os.environ.get("BROKER"),
    task_queues=queues.task_queues_scrap,
    task_routes={
        'main.run_modelo_spider': {'queue': SCRAPY_QUEUE}
    }
)

metrics_thread_started = False

def start_metrics_server():
    port = int(os.getenv("PORT", 8000))
    start_http_server(port)
    while True:
        RAM_USAGE.set(psutil.virtual_memory().percent)
        time.sleep(5)

@app.task
def run_dummy(id_job):
    print(f"Inicia run_dummy")
    subprocess.run(["scrapy", "crawl", "dummy", "-a", f"id_job={id_job}"])

@app.task
def ejecutar_funcion_items_postgres(_=None):
    print(f"Inicia ejecutar_funcion_postgres")
    url = os.getenv("DATABASE_URL")
    conn = psycopg2.connect(url)
    try:
        with conn.cursor() as cur:
            cur.execute("SELECT marketplacedemo.normalize_marketplace();")
            conn.commit()
    finally:
        conn.close()

@app.task
def run_modelo_spider():
    global metrics_thread_started
    if not metrics_thread_started:
        threading.Thread(target=start_metrics_server, daemon=True).start()
        metrics_thread_started = True
    id_job = int(time.time())  
    chain(
        run_dummy.s(id_job).set(queue=SCRAPY_QUEUE),
        ejecutar_funcion_items_postgres.s().set(queue=SCRAPY_QUEUE)
    ).apply_async()
