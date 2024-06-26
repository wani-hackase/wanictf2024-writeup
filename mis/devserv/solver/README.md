# Solver

```sh
docker build -t devserv-solver-ldap .
docker run -it --rm -p 1234:389 devserv-solver-ldap
```

Open another terminal and run the following command:

```sh
ssh -NR '5678:localhost:1234' git@chal-lz56g6.wanictf.org
```

Then open the browser and sign in with user id `alice`, password `alice_pass`, and ldap_url `ldap://localhost:5678`.
