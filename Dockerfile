FROM registry.access.redhat.com/ubi8/ubi:latest
LABEL Component="testapp" \
    Name="testapp" \
    Version="1.0" \
    Release="1"

LABEL io.k8s.description="Housekeeping for released Persistent Volumes" \
    io.openshift.tags="testapp"

WORKDIR /
COPY /bin/testapp /app/testapp
COPY certs/* /etc/pki/ca-trust/source/anchors/
RUN yum update -y && \
    chgrp -R 0 /app && \
    chmod -R g=u /app &&  \
    update-ca-trust
CMD /app/testapp
