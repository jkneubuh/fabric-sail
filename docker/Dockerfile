FROM registry.access.redhat.com/ubi8/ubi

RUN yum -y install jq

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# fabric binaries
RUN (pushd /tmp && \
      curl -sSL https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/bootstrap.sh \
      | bash -s -- -s -d)

RUN cp -r /tmp/bin/* /usr/local/bin