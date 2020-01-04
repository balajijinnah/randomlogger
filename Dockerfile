FROM frolvlad/alpine-glibc

COPY randomlogger /bin/randomlogger

CMD ["bin/randomlogger"]