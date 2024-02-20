# https://docs.percona.com/percona-operator-for-mongodb/users.html#system-users
# Values are base64 encoded https://www.base64encode.org/
kubectl patch secret/prod-secrets -p '{"data":{"PMM_SERVER_USER": "YWRtaW4=","PMM_SERVER_PASSWORD":"PT5BaSBjUHEkVi1iNXJfNg=="}}'
kubectl patch secret/booking-secrets -p '{"data":{"PMM_SERVER_USER": "YWRtaW4=","PMM_SERVER_PASSWORD":"PT5BaSBjUHEkVi1iNXJfNg=="}}'

kubectl patch secret/miller-secrets -p '{"data":{"PMM_SERVER_USER": "YWRtaW4=","PMM_SERVER_PASSWORD":"PT5BaSBjUHEkVi1iNXJfNg=="}}'

