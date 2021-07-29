FROM httpd:latest

WORKDIR /app

RUN apt update && \
    apt install vim -y

COPY resumo_docker.md /usr/share/nginx

CMD ["bash"]

