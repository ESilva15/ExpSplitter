FROM alpine:3.24.2


# Install the required packages
RUN apk --no-cache add bash zsh git go \
  && rm -rf /tmp/*


RUN mkdir /usr/share/webapps/
COPY ./build/expensesweb/ /usr/share/webapps/expensesweb


EXPOSE 8081


# Add the entrypoint
COPY ./start.sh /start.sh
STOPSIGNAL SIGTERM
ENTRYPOINT [ "/start.sh" ]
