FROM python:alpine
RUN apk add --update \
    py-pip
COPY . /app
WORKDIR /app
RUN pip install flask
EXPOSE 3000
ENTRYPOINT ["python"]
CMD ["app.py"]