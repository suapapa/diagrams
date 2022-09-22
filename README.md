# diagrams-back

## sandbox

Build and push docker image:

    $ docker build -f Dockerfile -t diagrams_sandbox:dev .

Example Run:

    $ cat sample/diagram_err.py | docker run -i --rm diagrams_sandbox:dev

Explore container:

    $ docker run -it --rm -v $(pwd)/sample:/sample --entrypoint /bin/bash diagrams_sandbox:dev
