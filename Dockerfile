FROM scratch

ADD build/castcloud /
ADD ca-certificates.crt /etc/ssl/certs/

VOLUME ["/data"]

ENTRYPOINT ["/castcloud"]
CMD ["-p=8080", "--dir=/data"]