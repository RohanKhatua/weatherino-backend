# Weatherino Backend

![Go](https://img.shields.io/badge/Go-1.20-blue)
![Gin Gonic](https://img.shields.io/badge/Gin%20Gonic-v1.9-lightblue)
![InfluxDB](https://img.shields.io/badge/InfluxDB-v2.0-orange)
![Docker Compose](https://img.shields.io/badge/Docker%20Compose-1.29.2-blue)
![License](https://img.shields.io/badge/license-MIT-green)

A RESTful API server for an Arduino-based weather station. Weatherino collects temperature and humidity data from an Arduino using the DHT11 sensor and stores it in a time-series database for efficient querying and analysis. The system ensures historical data integrity by calculating daily averages and removing granular data to optimize database size.

---

## Features

- Collects temperature and humidity data from an Arduino weather station.
- Stores data in **InfluxDB**, a time-series database optimized for high-performance queries.
- Calculates daily averages using a scheduled **cron job**, keeping only essential historical data.
- RESTful API built with **Gin Gonic** for fast and lightweight backend operations.
- Containerized with **Docker Compose** for easy setup and deployment.

---

## Tech Stack

- **Language:** [Go](https://golang.org/)
- **Web Framework:** [Gin Gonic](https://gin-gonic.com/)
- **Database:** [InfluxDB](https://www.influxdata.com/)
- **Task Scheduler:** [Cron](https://en.wikipedia.org/wiki/Cron)
- **Containerization:** [Docker Compose](https://docs.docker.com/compose/)

---

## Database Optimization

- **Minute-Level Data:** Removed daily via the cron job to prevent database bloating.
- **Daily Averages:** Calculated and stored for long-term historical analysis.

---

## Cron Job

The cron job is scheduled to:
1. Calculate daily average temperature and humidity at 00:00 UTC.
2. Delete minute-level data for the previous day.

Cron job configuration is handled within the Docker container for portability.

---





