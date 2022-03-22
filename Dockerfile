# pull official base image
FROM python:3

LABEL version="1.0"
LABEL description="ASN Website"

# setup enviroment variables
ENV USER=asn
ENV GROUP=asn
ENV HOME=/home/asn
ENV ASN_HOME=/home/asn/app

# setup flask variables
ENV FLASK_ENV=production

# create the asn user
RUN addgroup --system $USER
RUN adduser --system --ingroup $GROUP --home $HOME  $USER

# copy project
COPY src/app $ASN_HOME
COPY src/entrypoint.sh $HOME

# change directory
WORKDIR $ASN_HOME

# install dependencies
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# chown all the files to the app user
RUN chown -R asn:asn $ASN_HOME

# change to the app user
USER asn

EXPOSE 5000

HEALTHCHECK CMD curl --fail http://localhost:5000/ || exit 1

# change directory
WORKDIR $HOME

# Entrypoint
ENTRYPOINT ["/home/asn/entrypoint.sh"]