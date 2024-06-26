# FROM golang:1.22-alpine AS flagserv-builder
FROM golang@sha256:6522f0ca555a7b14c46a2c9f50b86604a234cdc72452bf6a268cae6461d9000b AS flagserv-builder

WORKDIR /root/flagserv

COPY flagserv/go.mod flagserv/go.sum ./
RUN go mod download && go mod verify

COPY flagserv .
RUN go build -v .

# FROM ubuntu:22.04
FROM ubuntu@sha256:19478ce7fc2ffbce89df29fea5725a8d12e57de52eb9ea570890dc5852aac1ac

WORKDIR /root

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y slapd ldap-utils git openssh-server && \
    rm -rf /var/lib/apt/lists/*

#===============
#  LDAP server
#===============

# Configure LDAP server
RUN tr -dc A-Za-z0-9 </dev/urandom | head -c 16 > admin_password.txt && \
    echo "slapd slapd/password1 password $(cat admin_password.txt)" | debconf-set-selections && \
    echo "slapd slapd/password2 password $(cat admin_password.txt)" | debconf-set-selections && \
    echo "slapd slapd/domain string wani.example.com" | debconf-set-selections && \
    echo "slapd shared/organization string WaniWani" | debconf-set-selections && \
    DEBIAN_FRONTEND=noninteractive dpkg-reconfigure slapd

# Add user "alice" to the LDAP directory
COPY setup.ldif setup.ldif
RUN slapd && \
    tr -dc A-Za-z0-9 </dev/urandom | head -c 16 > alice_password.txt && \
    sed -i "s#HASHED_ALICE_PASSWORD#$(slappasswd -T alice_password.txt)#" setup.ldif && \
    ldapadd -x -D cn=admin,dc=wani,dc=example,dc=com -y admin_password.txt -f setup.ldif

RUN rm admin_password.txt alice_password.txt setup.ldif

#===============
#  Git server
#===============

# Create git user
RUN useradd --user-group git && \
    echo 'git:G1Ta34WdBLDd!w%Z' | chpasswd && \
    chsh -s "$(which git-shell)" git

# Configure SSH server
RUN echo 'PermitRootLogin no' >> /etc/ssh/sshd_config && \
    echo 'AllowUsers git' >> /etc/ssh/sshd_config && \
    mkdir /run/sshd

# Create bare git repository
RUN mkdir /devserv.git && \
    cd /devserv.git && \
    git init --bare

# Make initial commit
COPY . /root/devserv
RUN cd /root/devserv && \
    echo 'FAKE{flag_is_redacted}' > flagserv/flag.txt && \
    git init && \
    git add . && \
    git config user.email "wani@example.com" && \
    git config user.name "Wani Hackase" && \
    git commit -m 'Initial commit' && \
    git remote add origin /devserv.git && \
    git push origin master && \
    rm -rf /root/devserv

RUN chown -R git:git /devserv.git && \
    chmod -R a-w /devserv.git

#===============
#  Flag server
#===============

COPY --from=flagserv-builder /root/flagserv/flagserv /root/flagserv

ENTRYPOINT slapd && /usr/sbin/sshd && /root/flagserv
