FROM dir.esilva.org/esilva/esilva-alpine:v1.0.0


# Install the required packages
RUN apk --no-cache add bash zsh git go \
  && rm -rf /tmp/*


RUN mkdir /usr/share/webapps/
COPY ./build/expensesweb/ /usr/share/webapps/expensesweb
RUN chown -R esilva:esilva /usr/share/webapps/expensesweb


EXPOSE 8081


# Add the entrypoint
USER $UNAME
COPY ./start.sh /start.sh
STOPSIGNAL SIGTERM
ENTRYPOINT [ "/start.sh" ]
