FROM kalilinux/kali

#Seta diretório de refencia do container
WORKDIR /app

#Origen[local] destino[container_path]
#COPY . /app

RUN apt update

RUN apt install python3 sqlite3 -y

RUN apt install python3-pip -y

#python3 -m pip freeze > requirements.txt
#RUN python3 -m pip install -r requirements.txt

#CMD ["bash"]
CMD ["tail", "-f", "/dev/null"]