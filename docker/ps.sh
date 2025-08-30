#docker compose -p myapp ps
docker ps -a --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | awk 'NR==1 || /^myapp/'