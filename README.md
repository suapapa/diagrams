# diagrams-back

## Usage

1. Create container
2. Checkout example


## Test diagrams container

Build and push docker image:

    $ docker build -t diagrams:dev .

Example Run:

    $ cat sample/diagram_err.py | docker run -i --rm diagrams:dev

Explore container:

    $ docker run -it --rm -v $(pwd)/sample:/sample --entrypoint /bin/bash diagrams:dev
