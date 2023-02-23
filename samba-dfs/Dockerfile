FROM ubuntu:latest

WORKDIR /tmp

RUN apt update && apt upgrade -y && apt install samba winbind krb5-user -y && apt clean && apt autoremove -y 

CMD [ "sh","-c","echo $(systemctl status smbd)" ]