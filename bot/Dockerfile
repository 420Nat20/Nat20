FROM python:3 as build
COPY ./.env ./.env
COPY ./Nat20.py ./Nat20.py
COPY ./requirements.txt ./requirements.txt

RUN pip install -r ./requirements.txt
CMD python3 Nat20.py
