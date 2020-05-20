FROM openjdk:11.0.7-jre-slim


ENV ANTLR_VERSION 4.8

ADD https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar /
RUN chmod +r /antlr-${ANTLR_VERSION}-complete.jar

WORKDIR /workspace

ENTRYPOINT java -jar ../antlr-${ANTLR_VERSION}-complete.jar -Dlanguage=Go -o parser query.g4 
