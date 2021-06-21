FROM frolvlad/alpine-glibc:latest

WORKDIR /

COPY build/backend .

CMD ./backend

EXPOSE 8081
