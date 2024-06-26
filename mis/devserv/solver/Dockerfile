FROM ubuntu:22.04

WORKDIR /root

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y slapd ldap-utils && \
    rm -rf /var/lib/apt/lists/*

# Configure LDAP server
RUN echo "slapd slapd/password1 password admin_pass" | debconf-set-selections && \
    echo "slapd slapd/password2 password admin_pass" | debconf-set-selections && \
    echo "slapd slapd/domain string wani.example.com" | debconf-set-selections && \
    echo "slapd shared/organization string WaniWani" | debconf-set-selections && \
    DEBIAN_FRONTEND=noninteractive dpkg-reconfigure slapd

# Add user "alice" to the LDAP directory
COPY setup.ldif setup.ldif
RUN slapd && \
    ldapadd -x -D cn=admin,dc=wani,dc=example,dc=com -w admin_pass -f setup.ldif

ENTRYPOINT slapd -d None
