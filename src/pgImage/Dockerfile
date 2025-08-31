FROM dir.esilva.org/esilva/esilva-alpine:v1.0.0


# Required ENV variables for pg
ENV LANG=en_GB.utf8


# Downloads the required packages for the webserver
RUN apk --no-cache add su-exec zsh \
  postgresql16 \
  &&  rm -rf /tmp/*


# Directory for the pg socket
RUN mkdir -p /run/postgresql && \
    chown -R $UNAME:$UNAME /run/postgresql
RUN mkdir -p /var/lib/postgresql/data && \
    chown -R $UNAME:$UNAME /var/lib/postgresql


COPY ./schema.sql /schema.sql


# Expose the pgsql port
EXPOSE 5431


# Copies the start script
COPY ./start.sh /start.sh
USER $UNAME
STOPSIGNAL SIGTERM
ENTRYPOINT [ "/start.sh" ]
