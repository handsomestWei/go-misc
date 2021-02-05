FROM centos:7

ADD   ./app   /work/app
ADD   localtime       /etc/localtime
ADD   ./conf/app.conf    /work/conf/app.conf
COPY  ./views  /work/views

WORKDIR /work/
USER root
CMD /bin/sh -c 'umask 000 && /work/app'
